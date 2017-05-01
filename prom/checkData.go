package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemChecks = "checks"
	//Label is the string label
	Label = "label"
)

var checkLabelNames = []string{HostName, ServiceDescription, CommandName}
var performanceDataLabelNames = []string{HostName, ServiceDescription, CommandName, Type, Label}

//ChecksActive is a Prometheus counter
var ChecksActive = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "active_total",
		Help:      "Amount of active checks executed",
	})

//CheckResults is a Prometheus counter
var CheckResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "results_total",
		Help:      "Amount of check results received",
	})

//CheckReturnCode is a Prometheus gauge vector
var CheckReturnCode = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "returncode",
		Help:      "returncode of a check",
	},
	checkLabelNames,
)

//CheckExecutionTime is a Prometheus gauge vector
var CheckExecutionTime = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "execution_time_seconds",
		Help:      "ExecutionTime of a check in seconds",
	},
	checkLabelNames,
)

//CheckLatency is a Prometheus gauge vector
var CheckLatency = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "latency_seconds",
		Help:      "Latency of a check in seconds",
	},
	checkLabelNames,
)

//CheckCurrentAttempt is a Prometheus gauge vector
var CheckCurrentAttempt = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "current_attempt",
		Help:      "CurrentAttempt of a check",
	},
	checkLabelNames,
)

//CheckMaxAttempts is a Prometheus gauge vector
var CheckMaxAttempts = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "max_attempts",
		Help:      "MaxAttempts of a check",
	},
	checkLabelNames,
)

//CheckState is a Prometheus gauge vector
var CheckState = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "state",
		Help:      "State of a check",
	},
	checkLabelNames,
)

//CheckStateType is a Prometheus gauge vector
var CheckStateType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "state_type",
		Help:      "StateType of a check",
	},
	checkLabelNames,
)

//CheckPerfGauge is a Prometheus gauge vector
var CheckPerfGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "performance_data",
		Help:      "Performance data with unknown unit",
	},
	performanceDataLabelNames,
)

//CheckPerfPercent is a Prometheus gauge vector
var CheckPerfPercent = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "performance_data_percent",
		Help:      "Performance data with percent as unit",
	},
	performanceDataLabelNames,
)

//CheckPerfSeconds is a Prometheus gauge vector
var CheckPerfSeconds = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "performance_data_seconds",
		Help:      "Performance data with seconds as unit",
	},
	performanceDataLabelNames,
)

//CheckPerfBytes is a Prometheus gauge vector
var CheckPerfBytes = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "performance_data_bytes",
		Help:      "Performance data with bytes as unit",
	},
	performanceDataLabelNames,
)

func initCheckData() {
	prometheus.MustRegister(ChecksActive)
	prometheus.MustRegister(CheckResults)
	prometheus.MustRegister(CheckReturnCode)
	prometheus.MustRegister(CheckExecutionTime)
	prometheus.MustRegister(CheckLatency)
	prometheus.MustRegister(CheckCurrentAttempt)
	prometheus.MustRegister(CheckMaxAttempts)
	prometheus.MustRegister(CheckState)
	prometheus.MustRegister(CheckStateType)
	prometheus.MustRegister(CheckPerfGauge)
	prometheus.MustRegister(CheckPerfPercent)
	prometheus.MustRegister(CheckPerfSeconds)
	prometheus.MustRegister(CheckPerfBytes)
}
