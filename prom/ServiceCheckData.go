package prom

import "github.com/prometheus/client_golang/prometheus"

var ServiceChecksActive = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "servicechecks_active",
		Help:      "Amount of active servicechecks executed",
	})

var ServiceChecksResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "servicecheck_results",
		Help:      "Amount of servicecheck results received",
	})

//ServiceCheckReturnCode is a Prometheus counter vector
var ServiceCheckReturnCode = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "servicecheck_result_returncode",
		Help:      "Amount of servicecheck which certain returncode",
	},
	[]string{"code"},
)

func initServiceCheckData() {
	prometheus.MustRegister(ServiceChecksActive)
	prometheus.MustRegister(ServiceChecksResults)
	prometheus.MustRegister(ServiceCheckReturnCode)
}
