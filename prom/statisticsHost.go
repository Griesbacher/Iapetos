package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemStatistics = "statistics"
)

var statsTypeLabelNames = []string{Type}

//StatsHostsOverall is a Prometheus gauge
var HostsStatsAmount = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatistics,
		Name:      "hosts_total",
		Help:      "Amount of Hosts total",
	},
)

//StatsHostsCheckType is a Prometheus gauge
var HostsStatsCheckType = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatistics,
		Name:      "hosts_check_type_total",
		Help:      "Amount of Hosts with certain checkresults",
	},
	statsTypeLabelNames,
)

//StatsHostsFlapping is a Prometheus gauge
var HostsStatsFlapping = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatistics,
		Name:      "hosts_flapping_total",
		Help:      "Amount of Hosts currently flapping",
	},
)

//HostsStatsChecksEnabled is a Prometheus gauge
var HostsStatsChecksEnabled = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemStatistics,
		Name:      "hosts_checks_enabled_total",
		Help:      "Amount of Hosts with enabled hockschecks",
	},
)

func initHostStatistics() {
	prometheus.MustRegister(HostsStatsAmount)
	prometheus.MustRegister(HostsStatsCheckType)
	prometheus.MustRegister(HostsStatsFlapping)
	prometheus.MustRegister(HostsStatsChecksEnabled)
}
