package middlewares

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gofiber/fiber/v2"
	"time"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests to the application",
		},
		[]string{"handler", "method"},
	)

	requestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_request_duration_seconds",
			Help:    "Histogram of the request duration",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestsDuration)
}

func PrometheusMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	c.Next()

	duration := time.Since(start).Seconds()
	handler := c.Route().Path

	requestsTotal.WithLabelValues(handler, c.Method()).Inc()
	requestsDuration.WithLabelValues(handler).Observe(duration)

	return nil
}

func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
