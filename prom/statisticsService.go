package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemStatisticsServices = "statistics_services"
)

//StatsServicesOverall is a Prometheus gauge
var StatsServicesAmount = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsServices,
		Name:      "total",
		Help:      "Amount of Services total",
	},
)

//StatsServicesCheckType is a Prometheus gauge
var StatsServicesCheckType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsServices,
		Name:      "check_type_total",
		Help:      "Amount of Services with certain checkresults",
	},
	statsTypeLabelNames,
)

//StatsServicesFlapping is a Prometheus gauge
var StatsServicesFlapping = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsServices,
		Name:      "flapping_total",
		Help:      "Amount of Services currently flapping",
	},
)

//StatsServicesChecksEnabled is a Prometheus gauge
var StatsServicesChecksEnabled = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsServices,
		Name:      "checks_enabled_total",
		Help:      "Amount of Services with enabled hockschecks",
	},
)

func initStatisticsService() {
	prometheus.MustRegister(StatsServicesAmount)
	prometheus.MustRegister(StatsServicesCheckType)
	prometheus.MustRegister(StatsServicesFlapping)
	prometheus.MustRegister(StatsServicesChecksEnabled)
}
