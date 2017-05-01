package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemStatisticsHosts = "statistics_hosts"
)

var statsTypeLabelNames = []string{Type}

//StatsHostsAmount is a Prometheus gauge
var StatsHostsAmount = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "total",
		Help:      "Amount of Hosts total",
	},
)

//StatsHostsCheckType is a Prometheus gauge
var StatsHostsCheckType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "check_type_total",
		Help:      "Amount of Hosts with certain checktypes",
	},
	statsTypeLabelNames,
)

//StatsHostsStateType is a Prometheus gauge
var StatsHostsStateType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "state_type_total",
		Help:      "Amount of Hosts with certain state",
	},
	statsTypeLabelNames,
)

//StatsHostsFlapping is a Prometheus gauge
var StatsHostsFlapping = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "flapping_total",
		Help:      "Amount of Hosts currently flapping",
	},
)

//StatsHostsFlexDowntime is a Prometheus gauge
var StatsHostsFlexDowntime = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "downtime_flex_total",
		Help:      "Amount of Hosts currently with a flex downtime",
	},
)

//StatsHostsDowntime is a Prometheus gauge
var StatsHostsDowntime = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "downtime_total",
		Help:      "Amount of Hosts currently with a downtime",
	},
)

//StatsHostsChecksEnabled is a Prometheus gauge
var StatsHostsChecksEnabled = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatisticsHosts,
		Name:      "checks_enabled_total",
		Help:      "Amount of Hosts with enabled hockschecks",
	},
)

func initStatisticsHost() {
	prometheus.MustRegister(StatsHostsAmount)
	prometheus.MustRegister(StatsHostsCheckType)
	prometheus.MustRegister(StatsHostsStateType)
	prometheus.MustRegister(StatsHostsFlapping)
	prometheus.MustRegister(StatsHostsFlexDowntime)
	prometheus.MustRegister(StatsHostsDowntime)
	prometheus.MustRegister(StatsHostsChecksEnabled)
}
