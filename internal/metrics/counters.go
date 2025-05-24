package metrics

import "github.com/prometheus/client_golang/prometheus"

var ApiRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "Total number of API requests",
	},
	[]string{"method", "route", "status"},
)
