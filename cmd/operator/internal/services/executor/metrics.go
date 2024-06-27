package executor

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "executor"
)

type Metrics struct {
	nativeTokenSpentAmount       prometheus.Gauge
	nativeTokenCurrentBalance    prometheus.Gauge
	executedWorkflowsAmountTotal prometheus.Counter
	errorsTotal                  prometheus.Counter
}

func NewMetrics() *Metrics {
	return &Metrics{
		nativeTokenSpentAmount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "native_token_spent_amount",
			Help:      "Total amount of native token spent",
		}),
		nativeTokenCurrentBalance: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "native_token_current_balance",
			Help:      "Current balance of native token",
		}),
		executedWorkflowsAmountTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "executed_workflows_amount_total",
			Help:      "Total amount of executed workflows",
		}),
		errorsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "errors_total",
			Help:      "Total amount of operator internal errors",
		}),
	}
}

// Describe implements prometheus.Collector interface.
func (m *Metrics) Describe(descs chan<- *prometheus.Desc) {
	m.nativeTokenSpentAmount.Describe(descs)
	m.nativeTokenCurrentBalance.Describe(descs)
	m.executedWorkflowsAmountTotal.Describe(descs)
	m.errorsTotal.Describe(descs)
}

// Collect implements prometheus.Collector interface.
func (m *Metrics) Collect(metrics chan<- prometheus.Metric) {
	m.nativeTokenSpentAmount.Collect(metrics)
	m.nativeTokenCurrentBalance.Collect(metrics)
	m.executedWorkflowsAmountTotal.Collect(metrics)
	m.errorsTotal.Collect(metrics)
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m)
}
