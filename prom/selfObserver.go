package prom

import "github.com/prometheus/client_golang/prometheus"

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

func initIapetos() {
	prometheus.MustRegister(ModuleDuration)
	prometheus.MustRegister(ModuleCallbacks)
}
