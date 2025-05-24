package middleware

import (
	"go-ecommerce/internal/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		route := c.FullPath()

		if route != "/metrics" {
			if route == "" {
				route = "unknown" // unmatched routes
			}

			metrics.ApiRequestsTotal.WithLabelValues(c.Request.Method, route, http.StatusText(status)).Inc()
		}
	}
}
