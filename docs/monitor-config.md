<!--- GENERATED BY gomplate from scripts/docs/monitor-main.md.tmpl --->

# Monitor Configuration

Monitors gather metrics from the host and from running applications.  They are
configured in a list called `monitors` in the [main agent config
file](./config-schema.md). These are all of the monitors included in the agent,
along with their possible configuration options:

- [cadvisor](./monitors/cadvisor.md)
- [collectd/activemq](./monitors/collectd-activemq.md)
- [collectd/apache](./monitors/collectd-apache.md)
- [collectd/cassandra](./monitors/collectd-cassandra.md)
- [collectd/chrony](./monitors/collectd-chrony.md)
- [collectd/consul](./monitors/collectd-consul.md)
- [collectd/couchbase](./monitors/collectd-couchbase.md)
- [collectd/cpu](./monitors/collectd-cpu.md)
- [collectd/cpufreq](./monitors/collectd-cpufreq.md)
- [collectd/custom](./monitors/collectd-custom.md)
- [collectd/df](./monitors/collectd-df.md)
- [collectd/disk](./monitors/collectd-disk.md)
- [collectd/docker](./monitors/collectd-docker.md)
- [collectd/elasticsearch](./monitors/collectd-elasticsearch.md)
- [collectd/etcd](./monitors/collectd-etcd.md)
- [collectd/genericjmx](./monitors/collectd-genericjmx.md)
- [collectd/hadoop](./monitors/collectd-hadoop.md)
- [collectd/hadoopjmx](./monitors/collectd-hadoopjmx.md)
- [collectd/haproxy](./monitors/collectd-haproxy.md)
- [collectd/health-checker](./monitors/collectd-health-checker.md)
- [collectd/interface](./monitors/collectd-interface.md)
- [collectd/jenkins](./monitors/collectd-jenkins.md)
- [collectd/kafka](./monitors/collectd-kafka.md)
- [collectd/kafka_consumer](./monitors/collectd-kafka_consumer.md)
- [collectd/kafka_producer](./monitors/collectd-kafka_producer.md)
- [collectd/kong](./monitors/collectd-kong.md)
- [collectd/load](./monitors/collectd-load.md)
- [collectd/marathon](./monitors/collectd-marathon.md)
- [collectd/memcached](./monitors/collectd-memcached.md)
- [collectd/memory](./monitors/collectd-memory.md)
- [collectd/mongodb](./monitors/collectd-mongodb.md)
- [collectd/mysql](./monitors/collectd-mysql.md)
- [collectd/nginx](./monitors/collectd-nginx.md)
- [collectd/openstack](./monitors/collectd-openstack.md)
- [collectd/postgresql](./monitors/collectd-postgresql.md)
- [collectd/processes](./monitors/collectd-processes.md)
- [collectd/protocols](./monitors/collectd-protocols.md)
- [collectd/python](./monitors/collectd-python.md)
- [collectd/rabbitmq](./monitors/collectd-rabbitmq.md)
- [collectd/redis](./monitors/collectd-redis.md)
- [collectd/signalfx-metadata](./monitors/collectd-signalfx-metadata.md)
- [collectd/spark](./monitors/collectd-spark.md)
- [collectd/statsd](./monitors/collectd-statsd.md)
- [collectd/uptime](./monitors/collectd-uptime.md)
- [collectd/vmem](./monitors/collectd-vmem.md)
- [collectd/zookeeper](./monitors/collectd-zookeeper.md)
- [docker-container-stats](./monitors/docker-container-stats.md)
- [host-metadata](./monitors/host-metadata.md)
- [internal-metrics](./monitors/internal-metrics.md)
- [kubelet-stats](./monitors/kubelet-stats.md)
- [kubernetes-cluster](./monitors/kubernetes-cluster.md)
- [kubernetes-events](./monitors/kubernetes-events.md)
- [kubernetes-volumes](./monitors/kubernetes-volumes.md)
- [prometheus-exporter](./monitors/prometheus-exporter.md)
- [trace-forwarder](./monitors/trace-forwarder.md)


## Common Configuration

The following config options are common to all monitors:

| Config option | Default | Required | Type | Description |
| --- | --- | --- | --- | --- |
| `type` |  | no | `string` | The type of the monitor |
| `discoveryRule` |  | no | `string` | The rule used to match up this configuration with a discovered endpoint. If blank, the configuration will be run immediately when the agent is started.  If multiple endpoints match this rule, multiple instances of the monitor type will be created with the same configuration (except different host/port). |
| `extraDimensions` |  | no | `map of string` | A set of extra dimensions (key:value pairs) to include on datapoints emitted by the monitor(s) created from this configuration. To specify metrics from this monitor should be high-resolution, add the dimension `sf_hires:1` |
| `intervalSeconds` | `0` | no | `integer` | The interval (in seconds) at which to emit datapoints from the monitor(s) created by this configuration.  If not set (or set to 0), the global agent intervalSeconds config option will be used instead. |
| `solo` | `false` | no | `bool` | If one or more configurations have this set to true, only those configurations will be considered -- useful for testing |
| `metricsToExclude` |  | no | `list of object (see below)` | A list of metric filters |
| `disableHostDimensions` | `false` | no | `bool` | Some monitors pull metrics from services not running on the same host and should not get the host-specific dimensions set on them (e.g. `host`, `AWSUniqueId`, etc).  Setting this to `true` causes those dimensions to be omitted.  You can disable this globally with the `disableHostDimensions` option on the top level of the config. |
| `disableEndpointDimensions` | `false` | no | `bool` | This can be set to true if you don't want to include the dimensions that are specific to the endpoint that was discovered by an observer.  This is useful when you have an endpoint whose identity is not particularly important since it acts largely as a proxy or adapter for other metrics. |

