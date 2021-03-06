<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# collectd/apache

 Monitors Apache webservice instances using
the information provided by `mod_status`.

See https://github.com/signalfx/integrations/tree/master/collectd-apache

Sample YAML configuration:

```
monitors:
 - type: collectd/apache
   host: localhost
   port: 80
```

If `mod_status` is exposed on an endpoint other than `/mod_status`, you can
use the `url` config option to specify the path:

```
monitors:
 - type: collectd/apache
   host: localhost
   port: 80
   url: "http://{{.Host}}:{{.Port}}/server-status?auto"
```


Monitor Type: `collectd/apache`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/apache)

**Accepts Endpoints**: **Yes**

**Multiple Instances Allowed**: Yes

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `host` | **yes** | `string` | The hostname of the Apache server |
| `port` | **yes** | `integer` | The port number of the Apache server |
| `name` | no | `string` | This will be sent as the `plugin_instance` dimension and can be any name you like. |
| `url` | no | `string` | The URL, either a final url or a Go template that will be populated with the host and port values. (**default:** `http://{{.Host}}:{{.Port}}/mod_status?auto`) |
| `username` | no | `string` |  |
| `password` | no | `string` |  |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are not marked as Custom are standard metrics and are monitored by default.

| Name | Type | Custom | Description |
| ---  | ---  | ---    | ---         |
| `apache_bytes` | cumulative |  | Bytes served by Apache |
| `apache_connections` | gauge |  | Connections served by Apache |
| `apache_idle_workers` | gauge |  | Apache workers that are idle |
| `apache_requests` | cumulative |  | Requests served by Apache |
| `apache_scoreboard.closing` | gauge | X | Number of workers in the process of closing connections |
| `apache_scoreboard.dnslookup` | gauge | X | Number of workers performing DNS lookup |
| `apache_scoreboard.finishing` | gauge | X | Number of workers that are finishing |
| `apache_scoreboard.idle_cleanup` | gauge | X | Number of idle threads ready for cleanup |
| `apache_scoreboard.keepalive` | gauge | X | Number of keep-alive connections |
| `apache_scoreboard.logging` | gauge | X | Number of workers writing to log file |
| `apache_scoreboard.open` | gauge |  | Number of worker thread slots that are open |
| `apache_scoreboard.reading` | gauge | X | Number of workers reading requests |
| `apache_scoreboard.sending` | gauge | X | Number of workers sending responses |
| `apache_scoreboard.starting` | gauge | X | Number of workers starting up |
| `apache_scoreboard.waiting` | gauge | X | Number of workers waiting for requests |

To specify custom metrics you want to monitor, add a negated `metricsToExclude` to the monitor configuration, as shown in the code snippet below. The snippet lists all available custom metrics. You can copy and paste the snippet into your configuration file, then delete any custom metrics that you do not want to monitor. 
Note that some of the custom metrics require you to set a flag as well as add them to the list. Check the monitor configuration file to see if a flag is required for gathering additional metrics.
```yaml 
metricsToExclude:
  - apache_scoreboard.closing
  - apache_scoreboard.dnslookup
  - apache_scoreboard.finishing
  - apache_scoreboard.idle_cleanup
  - apache_scoreboard.keepalive
  - apache_scoreboard.logging
  - apache_scoreboard.reading
  - apache_scoreboard.sending
  - apache_scoreboard.starting
  - apache_scoreboard.waiting
  negated: true
```



## Dimensions

The following dimensions may occur on metrics emitted by this monitor.  Some
dimensions may be specific to certain metrics.

| Name | Description |
| ---  | ---         |
| `plugin_instance` | Set to whatever you set in the `name` config option. |



