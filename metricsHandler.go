package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pschlump/fmcsa-svr/status"
)

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func appStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := status.App{}

		result.Version = GetVersion()
		result.TotalCount = status.StatStorage.GetTotalCount()

		c.JSON(http.StatusOK, result)
	}
}
