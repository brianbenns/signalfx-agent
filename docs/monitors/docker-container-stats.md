<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# docker-container-stats

 This monitor reads container stats from a
Docker API server.  It is meant as a metric-compatible replacement of our
[docker-collectd](https://github.com/signalfx/docker-collectd-plugin)
plugin, which scales rather poorly against a large number of containers.

This currently does not support CPU share/quota metrics.

If you are running the agent directly on a host (outside of a container
itself) and you are using the default Docker UNIX socket URL, you will
probably need to add the `signalfx-agent` user to the `docker` group in
order to have permission to access the Docker API via the socket.

Requires Docker API version 1.22+.


Monitor Type: `docker-container-stats`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/docker)

**Accepts Endpoints**: No

**Multiple Instances Allowed**: Yes

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `dockerURL` | no | `string` | The URL of the docker server (**default:** `unix:///var/run/docker.sock`) |
| `timeoutSeconds` | no | `integer` | The maximum amount of time to wait for docker API requests (**default:** `5`) |
| `labelsToDimensions` | no | `map of string` | A mapping of container label names to dimension names. The corresponding label values will become the dimension value for the mapped name.  E.g. `io.kubernetes.container.name: container_spec_name` would result in a dimension called `container_spec_name` that has the value of the `io.kubernetes.container.name` container label. |
| `excludedImages` | no | `list of string` | A list of filters of images to exclude.  Supports literals, globs, and regex. |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are not marked as Custom are standard metrics and are monitored by default.

| Name | Type | Custom | Description |
| ---  | ---  | ---    | ---         |
| `blkio.io_service_bytes_recursive.async` | cumulative | X | Volume, in bytes, of asynchronous block I/O |
| `blkio.io_service_bytes_recursive.read` | cumulative |  | Volume, in bytes, of reads from block devices |
| `blkio.io_service_bytes_recursive.sync` | cumulative | X | Volume, in bytes, of synchronous block I/O |
| `blkio.io_service_bytes_recursive.total` | cumulative | X | Total volume, in bytes, of all block I/O |
| `blkio.io_service_bytes_recursive.write` | cumulative |  | Volume, in bytes, of writes to block devices |
| `blkio.io_serviced_recursive.async` | cumulative | X | Number of asynchronous block I/O requests |
| `blkio.io_serviced_recursive.read` | cumulative | X | Number of reads requests from block devices |
| `blkio.io_serviced_recursive.sync` | cumulative | X | Number of synchronous block I/O requests |
| `blkio.io_serviced_recursive.total` | cumulative | X | Total number of block I/O requests |
| `blkio.io_serviced_recursive.write` | cumulative | X | Number of write requests to block devices |
| `cpu.percent` | gauge | X | Percentage of host CPU resources used by the container |
| `cpu.percpu.usage` | cumulative | X | Jiffies of CPU time spent by the container, per CPU core |
| `cpu.percpu.usage` | cumulative | X | Jiffies of CPU time spent by the container, per CPU core |
| `cpu.throttling_data.periods` | cumulative | X | Number of periods |
| `cpu.throttling_data.throttled_periods` | cumulative | X | Number of periods throttled |
| `cpu.throttling_data.throttled_time` | cumulative | X | Throttling time in nano seconds |
| `cpu.usage.kernelmode` | cumulative | X | Jiffies of CPU time spent in kernel mode by the container |
| `cpu.usage.system` | cumulative |  | Jiffies of CPU time used by the system |
| `cpu.usage.total` | cumulative |  | Jiffies of CPU time used by the container |
| `cpu.usage.usermode` | cumulative | X | Jiffies of CPU time spent in user mode by the container |
| `memory.percent` | gauge | X | Percent of memory (0-100) used by the container relative to its limit (excludes page cache usage) |
| `memory.stats.swap` | gauge | X | Bytes of swap memory used by container |
| `memory.usage.limit` | gauge |  | Memory usage limit of the container, in bytes |
| `memory.usage.max` | gauge | X | Maximum measured memory usage of the container, in bytes |
| `memory.usage.total` | gauge |  | Bytes of memory used by the container |
| `network.usage.rx_bytes` | cumulative |  | Bytes received by the container via its network interface |
| `network.usage.rx_dropped` | cumulative | X | Number of inbound network packets dropped by the container |
| `network.usage.rx_errors` | cumulative | X | Errors receiving network packets |
| `network.usage.rx_packets` | cumulative | X | Network packets received by the container via its network interface |
| `network.usage.tx_bytes` | cumulative |  | Bytes sent by the container via its network interface |
| `network.usage.tx_dropped` | cumulative | X | Number of outbound network packets dropped by the container |
| `network.usage.tx_errors` | cumulative | X | Errors sending network packets |
| `network.usage.tx_packets` | cumulative | X | Network packets sent by the container via its network interface |

To specify custom metrics you want to monitor, add a negated `metricsToExclude` to the monitor configuration, as shown in the code snippet below. The snippet lists all available custom metrics. You can copy and paste the snippet into your configuration file, then delete any custom metrics that you do not want to monitor. 
Note that some of the custom metrics require you to set a flag as well as add them to the list. Check the monitor configuration file to see if a flag is required for gathering additional metrics.
```yaml 
metricsToExclude:
  - blkio.io_service_bytes_recursive.async
  - blkio.io_service_bytes_recursive.sync
  - blkio.io_service_bytes_recursive.total
  - blkio.io_serviced_recursive.async
  - blkio.io_serviced_recursive.read
  - blkio.io_serviced_recursive.sync
  - blkio.io_serviced_recursive.total
  - blkio.io_serviced_recursive.write
  - cpu.percent
  - cpu.percpu.usage
  - cpu.percpu.usage
  - cpu.throttling_data.periods
  - cpu.throttling_data.throttled_periods
  - cpu.throttling_data.throttled_time
  - cpu.usage.kernelmode
  - cpu.usage.usermode
  - memory.percent
  - memory.stats.swap
  - memory.usage.max
  - network.usage.rx_dropped
  - network.usage.rx_errors
  - network.usage.rx_packets
  - network.usage.tx_dropped
  - network.usage.tx_errors
  - network.usage.tx_packets
  negated: true
```





