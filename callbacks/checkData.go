package callbacks

import (
	"github.com/griesbacher/Iapetos/helper"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	defaultFactor = 1
	msToSeconds   = 0.001
	minToSeconds  = 60
	hourToSeconds = minToSeconds * 60
)

func setPerformanceData(perfData string, labels prometheus.Labels) {
	for p := range helper.IteratePerformanceData(perfData) {
		labels["label"] = p.Label
		switch p.Unit {
		case "%", "percent", "pct":
			setGaugeValues(p.Data, defaultFactor, labels, prom.CheckPerfPercent)
		case "ms":
			setGaugeValues(p.Data, msToSeconds, labels, prom.CheckPerfSeconds)
		case "s":
			setGaugeValues(p.Data, defaultFactor, labels, prom.CheckPerfSeconds)
		case "m":
			setGaugeValues(p.Data, minToSeconds, labels, prom.CheckPerfSeconds)
		case "h":
			setGaugeValues(p.Data, hourToSeconds, labels, prom.CheckPerfSeconds)
		default:
			setGaugeValues(p.Data, defaultFactor, labels, prom.CheckPerfGauge)
		}
	}
}

func setGaugeValues(data map[string]float64, factor float64, labels prometheus.Labels, target *prometheus.GaugeVec) {
	for k, v := range data {
		labels["type"] = k
		target.With(labels).Set(v * factor)
	}
}
