package prom

import "github.com/prometheus/client_golang/prometheus"

var HostChecksActive = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostchecks_active",
		Help:      "Amount of active hostchecks executed",
	})

var HostResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostcheck_results",
		Help:      "Amount of hostcheck results received",
	})

func initHostCheckData() {
	prometheus.MustRegister(HostChecksActive)
	prometheus.MustRegister(HostResults)
}
