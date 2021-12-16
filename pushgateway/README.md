Test pushgateway

```
echo "some_metric 3.14" | curl --data-binary @- http://localhost:9091/metrics/job/some_job
./script.sh
```