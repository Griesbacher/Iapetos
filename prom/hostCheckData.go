package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace_core = "core"
	subsystem_events = "events"
)

var ActiveHostChecks = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "active_hostcheck",
		Help:      "Amount of active hostchecks executed",
	})

var HostResults = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: namespace_core,
		Subsystem: subsystem_events,
		Name:      "hostcheck_results",
		Help:      "Amount of hostcheck results recived",
	})

func initHostData() {
	prometheus.MustRegister(ActiveHostChecks)
	prometheus.MustRegister(HostResults)
}
