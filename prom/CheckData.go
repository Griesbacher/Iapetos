package prom

import "github.com/prometheus/client_golang/prometheus"

const subsystemChecks = "checks"

var checkLabelNames = []string{"host_name", "service_description", "command_name"}

//ChecksActive is a Prometheus counter
var ChecksActive = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "active",
		Help:      "Amount of active checks executed",
	})

//Results is a Prometheus counter
var CheckResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "results",
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

//CheckExecutionTime is a Prometheus counter vector
var CheckExecutionTime = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "execution_time",
		Help:      "ExecutionTime of a check in seconds",
	},
	checkLabelNames,
)

//CheckLatency is a Prometheus counter vector
var CheckLatency = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "latency",
		Help:      "Latency of a check in seconds",
	},
	checkLabelNames,
)

//CheckCurrentAttempt is a Prometheus counter vector
var CheckCurrentAttempt = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "current_attempt",
		Help:      "CurrentAttempt of a check",
	},
	checkLabelNames,
)

//CheckMaxAttempts is a Prometheus counter vector
var CheckMaxAttempts = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "max_attempts",
		Help:      "MaxAttempts of a check",
	},
	checkLabelNames,
)

//CheckState is a Prometheus counter vector
var CheckState = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "state",
		Help:      "State of a check",
	},
	checkLabelNames,
)

//CheckStateType is a Prometheus counter vector
var CheckStateType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemChecks,
		Name:      "state_type",
		Help:      "StateType of a check",
	},
	checkLabelNames,
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
}
