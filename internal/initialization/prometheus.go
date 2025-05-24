package initialization

import (
	"go-ecommerce/internal/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func InitPrometheus() {
	prometheus.MustRegister(metrics.ApiRequestsTotal)
}
