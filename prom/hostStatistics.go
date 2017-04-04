package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemStatistics = "statistics"
)

var statsTypeLabelNames = []string{Type}

//StatsHostsOverall is a Prometheus gauge
var HostsStatsOverall = prometheus.NewGauge(
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

func initHostStatistics() {
	prometheus.MustRegister(HostsStatsOverall)
	prometheus.MustRegister(HostsStatsCheckType)
	prometheus.MustRegister(HostsStatsFlapping)
}
