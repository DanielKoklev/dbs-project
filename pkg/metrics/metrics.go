package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
    WriteLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name:    "write_latency_ms",
        Buckets: prometheus.ExponentialBuckets(0.25, 2, 12),
    })
)