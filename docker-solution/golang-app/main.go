package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a test gauge with name and help string.
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_gauge",
		Help: "Test gauge. Its value is random.",
	})
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:        "golang_app_histogram",
		Help:        "Sample histogram",
		ConstLabels: map[string]string{"key": "value"},
		Buckets:     []float64{rand.Float64(), rand.Float64()},
	})
	// Register the gauge in the global metrics registry.
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	// Set new gauge
	gauge.Set(rand.Float64())
	//histogram.Write(prometheus.NewMetricVec(prometheus.Desc{
	//
	//}))
	// Expose all metrics in the default registry on /metrics.
	http.Handle("/metrics", promhttp.Handler())
	// Listen for HTTP requests on port 8005.
	log.Println("http server is live")
	log.Fatal(http.ListenAndServe(":8005", nil))
}
