package prom

import (
	"fmt"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/prometheus/client_golang/prometheus"
)

const subsystemNebModule = "neb_module"

//ModuleDuration is a Prometheus Histogram vector
var ModuleDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNebModule,
		Name:      "callback_durations_seconds",
		Help:      "Time in seconds for by each callback type",
		Buckets: []float64{
			.00001, .00005,
			.0001, .0005,
			.001, .005,
			.01, .05,
			.1, .5,
			1, 1.5,
			2},
	},
	[]string{"type"},
)

//ModuleCallbacks is a Prometheus Gauge vector
var ModuleCallbacks = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNebModule,
		Name:      "callbacks_total",
		Help:      "Amount of registered callbacks per type",
	},
	[]string{"type"},
)

//CoreType is a Prometheus Gauge
var CoreType = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNebModule,
		Name:      "core_type",
		Help:      fmt.Sprintf("The core this module is compiled for. Nagios3: %d, Nagios4: %d, Naemon: %d.", neb.CoreNagios3, neb.CoreNagios4, neb.CoreNaemon),
	},
)

func initIapetos() {
	prometheus.MustRegister(ModuleDuration)
	prometheus.MustRegister(ModuleCallbacks)
	prometheus.MustRegister(CoreType)
	CoreType.Set(float64(neb.GetCoreType()))
}
