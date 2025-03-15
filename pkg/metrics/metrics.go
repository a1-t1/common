// Package metrics ...
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// DefaultLabels contains Global Shared Labels.
var DefaultLabels map[string]string

// NewCounter ...
func NewCounter(
	name, help string,
	labels []string,
) *prometheus.CounterVec {
	opts := prometheus.CounterOpts{
		Name:        name,
		Help:        help,
		ConstLabels: DefaultLabels,
	}
	return prometheus.NewCounterVec(opts, labels)
}

// NewGauge ...
func NewGauge(
	name, help string,
) prometheus.Gauge {
	opts := prometheus.GaugeOpts{
		Name:        name,
		Help:        help,
		ConstLabels: DefaultLabels,
	}
	return prometheus.NewGauge(opts)
}

// NewHistogram ...
func NewHistogram(
	name, help string,
	labels []string,
	Buckets []float64,
) *prometheus.HistogramVec {
	opts := prometheus.HistogramOpts{
		Name:        name,
		Help:        help,
		ConstLabels: DefaultLabels,
	}
	if len(Buckets) != 0 {
		opts.Buckets = Buckets
	}
	return prometheus.NewHistogramVec(opts, labels)
}
