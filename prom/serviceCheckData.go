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

func initServiceCheckData() {
	prometheus.MustRegister(ServiceChecksActive)
	prometheus.MustRegister(ServiceChecksResults)
}
