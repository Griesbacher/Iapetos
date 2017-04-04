package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemStatisticsHosts = "statistics_hosts"
)

var statsTypeLabelNames = []string{Type}

//StatsHostsOverall is a Prometheus gauge
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
		Help:      "Amount of Hosts with certain checkresults",
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
	prometheus.MustRegister(StatsHostsFlapping)
	prometheus.MustRegister(StatsHostsChecksEnabled)
}