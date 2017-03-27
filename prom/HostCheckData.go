package prom

import "github.com/prometheus/client_golang/prometheus"

//HostChecksActive is a Prometheus counter
var HostChecksActive = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostchecks_active",
		Help:      "Amount of active hostchecks executed",
	})

//HostResults is a Prometheus counter
var HostResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostcheck_results",
		Help:      "Amount of hostcheck results received",
	})

//HostCheckReturnCode is a Prometheus counter vector
var HostCheckReturnCode = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostcheck_result_returncode",
		Help:      "Amount of hostcheck which certain returncode",
	},
	[]string{"code"},
)

func initHostCheckData() {
	prometheus.MustRegister(HostChecksActive)
	prometheus.MustRegister(HostResults)
	prometheus.MustRegister(HostCheckReturnCode)
}
