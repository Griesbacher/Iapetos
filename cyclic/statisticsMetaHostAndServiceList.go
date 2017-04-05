package cyclic

import (
	"github.com/ConSol/go-neb-wrapper/neb/checkTypes"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func countStates(meta structs.MetaHostAndServiceList, target *prometheus.GaugeVec, counterMap map[string]float64, toString func(int) string) {
	for _, m := range meta {
		t := toString(m.CurrentState)
		if _, contained := counterMap[t]; !contained {
			counterMap[t] = 0
		}
		counterMap[t]++
	}
	for k, v := range counterMap {
		target.With(prometheus.Labels{
			prom.Type: k,
		}).Set(v)
	}
}

func countTypes(meta structs.MetaHostAndServiceList, target *prometheus.GaugeVec) {
	counterMap := map[string]float64{
		checkTypes.CheckTypeToString(checkTypes.Active):  0,
		checkTypes.CheckTypeToString(checkTypes.Passive): 0,
		checkTypes.CheckTypeToString(checkTypes.Parent):  0,
		checkTypes.CheckTypeToString(checkTypes.File):    0,
		checkTypes.CheckTypeToString(checkTypes.Other):   0,
	}
	for _, m := range meta {
		t := checkTypes.CheckTypeToString(m.CheckType)
		if _, contained := counterMap[t]; !contained {
			counterMap[t] = 0
		}
		counterMap[t]++
	}
	for k, v := range counterMap {
		target.With(prometheus.Labels{
			prom.Type: k,
		}).Set(v)
	}
}

func countMinorStats(meta structs.MetaHostAndServiceList, gFlapping, gEnabled, gFlexDowntime, gDowntime prometheus.Gauge) {
	flapping, enabled, flexDowntime, downtime := 0.0, 0.0, 0.0, 0.0
	for _, m := range meta {
		if m.IsFlapping > 0 {
			flapping++
		}
		if m.ChecksEnabled > 0 {
			enabled++
		}
		if m.PendingFlexDowntimeDepth > 0 {
			flexDowntime++
		}
		if m.ScheduledDowntimeDepth > 0 {
			downtime++
		}
	}
	gFlapping.Set(flapping)
	gEnabled.Set(enabled)
	gFlexDowntime.Set(flexDowntime)
	gDowntime.Set(downtime)
}
