package writer

import (
	"context"
	"sync/atomic"

	"github.com/davecgh/go-spew/spew"
	"github.com/signalfx/golib/datapoint"
	"github.com/signalfx/golib/trace"
	"github.com/signalfx/signalfx-agent/internal/core/common/dpmeta"
	"github.com/signalfx/signalfx-agent/internal/core/writer/tracetracker"
	"github.com/signalfx/signalfx-agent/internal/utils"
	log "github.com/sirupsen/logrus"
)

func (sw *SignalFxWriter) listenForTraceSpans() {
	reqSema := make(chan struct{}, sw.conf.MaxRequests)

	// The only reason this is on the struct and not a local var is so we can
	// easily get diagnostic metrics from it
	sw.serviceTracker = sw.startGeneratingHostCorrelationMetrics()

	for {
		select {
		case <-sw.ctx.Done():
			return

		case span := <-sw.spanChan:
			buf := append(sw.spanBufferPool.Get().([]*trace.Span), span)
			buf = sw.drainSpanChan(buf)

			for i := range buf {
				sw.preprocessSpan(buf[i])
			}

			if *sw.conf.SendTraceHostCorrelationMetrics {
				sw.serviceTracker.AddSpans(sw.ctx, buf)
			}

			atomic.AddInt64(&sw.traceSpansInFlight, int64(len(buf)))

			go func() {
				defer sw.spanBufferPool.Put(buf[:0])

				// Wait if there are more than the max outstanding requests
				reqSema <- struct{}{}

				atomic.AddInt64(&sw.traceSpanRequestsActive, 1)
				// This sends synchonously
				err := sw.client.AddSpans(context.Background(), buf)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("Error shipping trace spans to SignalFx")
					// If there is an error sending spans then just forget about them.
					return
				}
				atomic.AddInt64(&sw.traceSpansSent, int64(len(buf)))
				log.Debugf("Sent %d trace spans to SignalFx", len(buf))

				<-reqSema

				atomic.AddInt64(&sw.traceSpanRequestsActive, -1)
				atomic.AddInt64(&sw.traceSpansInFlight, -int64(len(buf)))
			}()
		}
	}
}

func (sw *SignalFxWriter) drainSpanChan(buf []*trace.Span) []*trace.Span {
	for {
		select {
		case span := <-sw.spanChan:
			buf = append(buf, span)
			if len(buf) >= sw.conf.TraceSpanMaxBatchSize {
				return buf
			}
		default:
			return buf
		}
	}
}

func (sw *SignalFxWriter) preprocessSpan(span *trace.Span) {
	// Some spans aren't really specific to the host they are running
	// on and shouldn't have any host-specific tags.  This is indicated by a
	// special tag key (value is irrelevant).
	if _, ok := span.Tags[dpmeta.NotHostSpecificMeta]; !ok {
		span.Tags = sw.addhostIDFields(span.Tags)
	} else {
		// Get rid of the tag so it doesn't pass through to the backend
		delete(span.Tags, dpmeta.NotHostSpecificMeta)
	}

	if sw.conf.LogTraceSpans {
		log.Debugf("Sending trace span:\n%s", spew.Sdump(span))
	}
}

func (sw *SignalFxWriter) startGeneratingHostCorrelationMetrics() *tracetracker.ActiveServiceTracker {
	tracker := tracetracker.New(sw.conf.StaleServiceTimeout, func(dp *datapoint.Datapoint) {
		// Immediately send correlation datapoints when we first see a service
		sw.dpChan <- dp
	})

	// Send the correlation datapoints at a regular interval to keep the
	// service live on the backend.
	utils.RunOnInterval(sw.ctx, func() {
		for _, dp := range tracker.CorrelationDatapoints() {
			sw.dpChan <- dp
		}
	}, sw.conf.TraceHostCorrelationMetricsInterval)

	return tracker
}
