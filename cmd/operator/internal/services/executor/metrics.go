package executor

import (
	"context"
	"time"

	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/stats"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace               = "executor"
	backgroundCheckInterval = 5 * time.Second
)

type Metrics struct {
	nativeTokenSpentAmountTotal  prometheus.Counter
	nativeTokenCurrentBalance    prometheus.Gauge
	executedWorkflowsAmountTotal prometheus.Counter
	errorsTotal                  prometheus.Counter
	operatorCPUUsage             prometheus.Gauge
	// operatorMemoryUsage          prometheus.Gauge // prometheus collects this metric by default
}

func NewMetrics() *Metrics {
	return &Metrics{
		nativeTokenSpentAmountTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "native_token_spent_amount_total",
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
		operatorCPUUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "operator_cpu_usage",
			Help:      "CPU usage of the operator",
		}),
		// operatorMemoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
		// 	Namespace: namespace,
		// 	Name:      "operator_memory_usage",
		// 	Help:      "Memory usage of the operator",
		// }),
	}
}

func (m *Metrics) CountNativeTokenSpentAmountTotal(cnt uint64) {
	m.nativeTokenSpentAmountTotal.Add(float64(cnt))
}

func (m *Metrics) SetNativeTokenCurrentBalance(cnt int) {
	m.nativeTokenCurrentBalance.Set(float64(cnt))
}

func (m *Metrics) CountExecutedWorkflowsAmountTotal(cnt int) {
	m.executedWorkflowsAmountTotal.Add(float64(cnt))
}

func (m *Metrics) CountErrorsTotal(cnt int) {
	m.errorsTotal.Add(float64(cnt))
}

func (m *Metrics) SetCPUUsage(cnt float64) {
	m.operatorCPUUsage.Set(cnt)
}

// func (m *Metrics) SetOperatorMemoryUsage(cnt uint64) {
// 	m.operatorMemoryUsage.Set(float64(cnt))
// }

// Describe implements prometheus.Collector interface.
func (m *Metrics) Describe(descs chan<- *prometheus.Desc) {
	m.nativeTokenSpentAmountTotal.Describe(descs)
	m.nativeTokenCurrentBalance.Describe(descs)
	m.executedWorkflowsAmountTotal.Describe(descs)
	m.errorsTotal.Describe(descs)
	m.operatorCPUUsage.Describe(descs)
	// m.operatorMemoryUsage.Describe(descs)
}

// Collect implements prometheus.Collector interface.
func (m *Metrics) Collect(metrics chan<- prometheus.Metric) {
	m.nativeTokenSpentAmountTotal.Collect(metrics)
	m.nativeTokenCurrentBalance.Collect(metrics)
	m.executedWorkflowsAmountTotal.Collect(metrics)
	m.errorsTotal.Collect(metrics)
	m.operatorCPUUsage.Collect(metrics)
	// m.operatorMemoryUsage.Collect(metrics)
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m)
}

func (m *Metrics) CollectBackgroundMetrics(client ethereumClient) {
	for {
		// mem := stats.GetMemoryUsage()
		// m.SetOperatorMemoryUsage(mem)

		cpu, err := stats.GetCPUUsage()
		if err != nil {
			log.With(log.Err(err)).Error("get cpu usage error")
		} else {
			m.SetCPUUsage(cpu)
		}

		balance, err := client.GetBalance(context.Background())
		if err != nil {
			log.With(log.Err(err)).Error("get balance error")
		} else {
			m.nativeTokenCurrentBalance.Set(float64(balance.Int64()))
		}

		time.Sleep(backgroundCheckInterval)
	}
}
