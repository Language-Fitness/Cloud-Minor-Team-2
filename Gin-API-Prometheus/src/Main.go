package main

import (
	"Gin-API-Prometeus/src/Custom-metrics"
	"Gin-API-Prometeus/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	Custom_metrics.RecordMetrics()
	router := gin.Default()
	routes.UserRoutes(router)
	router.GET("/metrics", prometheusHandler())
	router.Run(":8080")
}
