#!/bin/bash
curl http://localhost:9091/metrics/job/demo_batch_job \
--data @<(cat <<EOF
# HELP demo_batch_job_last_run_timestamp_seconds Last Unix time sdwhen changing this group in the Pushgateway succeeded.
# TYPE demo_batch_job_last_run_timestamp_seconds gauge
demo_batch_job_last_run_timestamp_seconds 1
EOF
)