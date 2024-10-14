package executor

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "executor"
)

type Metrics struct {
	// Counters
	nativeTokenSpentAmountTotal  prometheus.Counter
	executedWorkflowsAmountTotal prometheus.Counter
	sentWorkflowsAmountTotal     prometheus.Counter
	errorsTotal                  prometheus.Counter

	// Gauges
	nativeTokenCurrentBalance prometheus.Gauge

	// Histograms
	miningLatencySeconds           prometheus.Histogram
	blockProcessingDurationSeconds prometheus.Histogram
}

func NewMetrics() *Metrics {
	return &Metrics{
		// Counters
		executedWorkflowsAmountTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "executed_workflows_total",
			Help:      "Total amount of executed workflows",
		}),
		sentWorkflowsAmountTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "sent_workflows_total",
			Help:      "Total amount of workflows sent to chain",
		}),
		errorsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "errors_total",
			Help:      "Total amount of operator internal errors",
		}),
		nativeTokenSpentAmountTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "native_token_spent_total",
			Help:      "Total amount of native token spent",
		}),

		// Gauges
		nativeTokenCurrentBalance: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "native_token_current_balance",
			Help:      "Current balance of native token",
		}),

		// Histograms
		miningLatencySeconds: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "mining_latency_seconds",
			Help:      "Latency in seconds between broadcast and inclusion into block",
		}),
		blockProcessingDurationSeconds: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "block_processing_duration_seconds",
			Help:      "Duration in seconds for processing a block",
		}),
	}
}

// Describe implements prometheus.Collector interface.
func (m *Metrics) Describe(descs chan<- *prometheus.Desc) {
	// Counters
	m.nativeTokenSpentAmountTotal.Describe(descs)
	m.executedWorkflowsAmountTotal.Describe(descs)
	m.sentWorkflowsAmountTotal.Describe(descs)
	m.errorsTotal.Describe(descs)

	// Gauges
	m.nativeTokenCurrentBalance.Describe(descs)

	// Histograms
	m.miningLatencySeconds.Describe(descs)
	m.blockProcessingDurationSeconds.Describe(descs)
}

// Collect implements prometheus.Collector interface.
func (m *Metrics) Collect(metrics chan<- prometheus.Metric) {
	// Counters
	m.nativeTokenSpentAmountTotal.Collect(metrics)
	m.executedWorkflowsAmountTotal.Collect(metrics)
	m.sentWorkflowsAmountTotal.Collect(metrics)
	m.errorsTotal.Collect(metrics)

	// Gauges
	m.nativeTokenCurrentBalance.Collect(metrics)

	// Histograms
	m.miningLatencySeconds.Collect(metrics)
	m.blockProcessingDurationSeconds.Collect(metrics)
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m)
}
