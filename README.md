<img src="img/logo.png">


Run docker solution
```
cd docker-solution
docker-compose build
docker-compose up
```


Sample metrics
```
demo_api_request_duration_seconds_count
demo_api_request_duration_seconds_count{status="200"}
demo_api_request_duration_seconds_count{method="GET",status="200"}
There are several other label matcher types that Prometheussupports besides equality matching:

● !=: Match only labels that have a different valuethan the one provided.
● =~: Match only labels whose value matches a providedregular expression.
● !~: Match only labels whose value does not match aprovided regular expression


demo_api_request_duration_seconds_count{path=~"/api/(foo|bar)"}
rate(demo_api_request_duration_seconds_count{job="demo"}[5m])

```

<b>rate()</b> calculates the per-second increase of a counter as averaged over a specified window of time. Functions only work for counter metrics.
```
rate(demo_api_request_duration_seconds_count{job="demo"}[5m])
```

Query the total increase over a given time
window. Functions only work for counter metrics.
```
increase(demo_api_request_duration_seconds_count{job="demo"}[1h])
```

<b>Counter metrics</b> treat any
decrease in value as a counter reset and can only output non-negative result

<b>Gauge metrics </b>track a value that can go up or down (like temperatures or memory usage or free disk space)

```
delta(demo_disk_usage_bytes{job="demo"}[15m])
```
<b>delta()</b> only considers the first and the last data point under the provided time window and
disregards overall trends in the intermediate data points

calculate by how much the disk usage is going up
or down per-second when looking at a 15-minute window
```
deriv(demo_disk_usage_bytes{job="demo"}[15m])
```

<b>predict_linear()</b> function is an extension to this that allows you to predict what the value of a
gauge will be in a given amount of time in the future
```
predict_linear(demo_disk_usage_bytes{job="demo"}[15m], 3600)
```

calculate total rates per instance and path, but not care about individual method or status results
```
sum without(method, status)
(rate(demo_api_request_duration_seconds_count{job="demo"}[5m]))
```

While <b>sum()</b> is the most commonly used aggregation operator, Prometheus supports many kinds of
aggregation operators, where some take extra arguments:
```
● sum(): Calculates the sum over dimensions.
● min(): Selects the minimum value across dimensions.
● max(): Selects the maximum value across dimensions.
● avg(): Calculates the average over dimensions.
● stddev(): Calculates the population standard deviation over dimensions.
● stdvar(): Calculates the population standard variance over dimensions.
● count(): Counts number of series in the vector.
● count_values(value_label, ...): Counts the number of series for each distinct sample
value.
● bottomk(k, ...): Selects the smallest k elements by sample value.
● topk(k, ...): Selects the largest k elements by sample value.
```

 calculate the average HTTP response time
from a histogram's _count series (total count of requests) and _sum series (total time spent in
requests), while preserving all dimensions in the output. The following query calculates the average
response time as averaged over the last five minutes, by dividing the increase in request counts by the
total time spent in requests during that time window
```
rate(demo_api_request_duration_seconds_count{job="demo"}[5m])
/
rate(demo_api_request_duration_seconds_sum{job="demo"}[5m])
```

The metric node_cpu_seconds_total tracks how many CPU seconds have been used since boot
time per core (cpu label) and per mode (mode label). To see how many cores are used in each core
and mode, calculate the rate over this counter
```
rate(node_cpu_seconds_total{job="node_exporter"}[1m])
```
Actual CPU usage, filter out the idle mode. To see the actual CPU usage over all cores
```
sum without(cpu) (rate(node_cpu_seconds_total{mode!="idle",job="node_exporter"}[1m]))
```

How much memory (in GiB) is available on the machine, add the free, buffers, and cached
memory amounts

```
(
node_memory_MemFree_bytes{job="node_exporter"}
+
node_memory_Buffers_bytes{job="node_exporter"}
+
node_memory_Cached_bytes{job="node_exporter"}
) / 1024^3
```

Sum of the incoming and outgoing network traffic on the machine, grouped by
network interface
```
rate(node_network_receive_bytes_total[1m])
+
rate(node_network_transmit_bytes_total[1m])
```

System's boot time as a Unix
timestamp
```
node_boot_time_seconds
```

Finds nodes that have rebooted more than 3 times in the last 30 minutes
```
changes(node_boot_time_seconds[30m]) > 3
```

Metric Types:
- Counters track values that can only ever go up, like HTTP request counts or CPU seconds used.
- Gauges track values that can go up or down, like temperatures or disk space.
- Summaries calculate client-side-calculated quantiles from a set of observations, like request latency percentiles. They also track the total count and total sum of observations (like the total count of requests and the total sum of request durations).
- Histograms track bucketed histograms of a set of observations like request latencies. They also track the total count and total sum of observations.

## Kubernetes solution 

```
kubectl create clusterrolebinding permissive-binding --clusterrole=cluster-admin --user=admin --user=kubelet --group=system:serviceaccounts
```

```
cd k8s-solution
kubectl apply -f . -nprometheus
```

Aggregate CPU usage for all containers,summed up by Kubernetes namespace:
```
sum by(namespace) (rate(container_cpu_usage_seconds_total[1m]))
```

```
- `prometheus.io/scrape`: Only scrape services that have a value of `true`
- `prometheus.io/scheme`: If the metrics endpoint is secured then you will need
  to set this to `https` & most likely set the `tls_config` of the scrape config.
- `prometheus.io/path`: If the metrics path is not `/metrics` override this.
- `prometheus.io/port`: If the metrics are exposed on a different port to the
  service then set this appropriately.
```
