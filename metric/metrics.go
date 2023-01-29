package metric

import (
	"github.com/golang-queue/queue"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pschlump/fmcsa/status"
)

const namespace = "fmcsa_svr_"

// Metrics implements the prometheus.Metrics interface and
// exposes fmcsa_svr_ metrics for prometheus
type Metrics struct {
	TotalRequestsCount *prometheus.Desc
	FmcsaSuccess       *prometheus.Desc
	FmcsaError         *prometheus.Desc
	q                  *queue.Queue
}

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics(q *queue.Queue) Metrics {
	m := Metrics{
		TotalRequestsCount: prometheus.NewDesc(
			namespace+"total_push_count",
			"Number of push count",
			nil, nil,
		),
		FmcsaSuccess: prometheus.NewDesc(
			namespace+"ios_success",
			"Number of iOS success count",
			nil, nil,
		),
		FmcsaError: prometheus.NewDesc(
			namespace+"ios_error",
			"Number of iOS fail count",
			nil, nil,
		),
		q: q,
	}

	return m
}

// Describe returns all possible prometheus.Desc
func (c Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TotalRequestsCount
	ch <- c.FmcsaSuccess
	ch <- c.FmcsaError
}

// Collect returns the metrics with values
func (c Metrics) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		c.TotalRequestsCount,
		prometheus.CounterValue,
		float64(status.StatStorage.GetTotalCount()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.FmcsaSuccess,
		prometheus.CounterValue,
		float64(status.StatStorage.GetFmcsaSuccess()),
	)
	ch <- prometheus.MustNewConstMetric(
		c.FmcsaError,
		prometheus.CounterValue,
		float64(status.StatStorage.GetFmcsaError()),
	)
}
