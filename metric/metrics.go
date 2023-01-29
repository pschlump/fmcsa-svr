package metric

// Copyright (c) Philip Schlump, 2023.
// This file is MIT licensed, see ../LICENSE.mit

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pschlump/fmcsa-svr/status"
)

// Metrics implements the prometheus.Metrics interface and
// exposes fmcsa_svr_ metrics for prometheus
type Metrics struct {
	TotalRequestsCount *prometheus.Desc
	FmcsaSuccess       *prometheus.Desc
	FmcsaError         *prometheus.Desc
	CacheSuccess       *prometheus.Desc
	CacheError         *prometheus.Desc
}

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics() Metrics {
	m := Metrics{
		TotalRequestsCount: prometheus.NewDesc(
			"fmcsa_svr_total_request_count",
			"Number of request count",
			nil, nil,
		),
		FmcsaSuccess: prometheus.NewDesc(
			"fmcsa_svr_fmcsa_success",
			"Number of FMCSA success count",
			nil, nil,
		),
		FmcsaError: prometheus.NewDesc(
			"fmcsa_svr_fmcsa_error",
			"Number of FMCSA fail count",
			nil, nil,
		),
		CacheSuccess: prometheus.NewDesc(
			"fmcsa_svr_cache_success",
			"Number of FMCSA requests from cache success count",
			nil, nil,
		),
		CacheError: prometheus.NewDesc(
			"fmcsa_svr_cache_error",
			"Number of FMCSA requests from cache fail count",
			nil, nil,
		),
	}

	return m
}

// Describe returns all possible prometheus.Desc
func (c Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TotalRequestsCount
	ch <- c.FmcsaSuccess
	ch <- c.FmcsaError
	ch <- c.CacheSuccess
	ch <- c.CacheError
}

// Collect returns the metrics with values
func (c Metrics) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.TotalRequestsCount, prometheus.CounterValue, float64(status.StatStorage.GetTotalCount()))
	ch <- prometheus.MustNewConstMetric(c.FmcsaSuccess, prometheus.CounterValue, float64(status.StatStorage.GetFmcsaSuccess()))
	ch <- prometheus.MustNewConstMetric(c.FmcsaError, prometheus.CounterValue, float64(status.StatStorage.GetFmcsaError()))
	ch <- prometheus.MustNewConstMetric(c.CacheSuccess, prometheus.CounterValue, float64(status.StatStorage.GetCacheSuccess()))
	ch <- prometheus.MustNewConstMetric(c.CacheError, prometheus.CounterValue, float64(status.StatStorage.GetCacheError()))
}
