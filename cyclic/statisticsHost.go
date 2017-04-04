package cyclic

import (
	"time"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/checkTypes"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

//HostStatistics monitors the core host stats
type HostStatistics struct {
	Stoppable
}

//StartHostStatistics creates an new HostStatistics and starts it
func StartHostStatistics() Stoppable {
	s := HostStatistics{Stoppable{stop: make(chan bool)}}
	go s.run()
	return s.Stoppable
}

func (s HostStatistics) run() {
	for {
		select {
		case <-s.stop:
			logging.GetLogger().Info("Stopping HostStatistics")
			s.stop <- true
			return
		case <-time.After(time.Duration(10) * time.Second):
			hosts := neb.GetHosts()
			if len(hosts) == 0 {
				continue
			}
			prom.StatsHostsAmount.Set(float64(len(hosts)))
			countHostTypes(hosts)
			countMinorHostStats(hosts)
		}
	}
}

func countHostTypes(hosts structs.Hostlist) {
	counterMap := map[string]float64{}
	for _, h := range hosts {
		t := checkTypes.CheckTypeToString(h.CheckType)
		if _, contained := counterMap[t]; !contained {
			counterMap[t] = 0
		}
		counterMap[t]++
	}
	for k, v := range counterMap {
		prom.StatsHostsCheckType.With(prometheus.Labels{
			prom.Type: k,
		}).Set(v)
	}
}

func countMinorHostStats(hosts structs.Hostlist) {
	flapping := 0.0
	enabled := 0.0

	for _, h := range hosts {
		if h.IsFlapping > 0 {
			flapping++
		}
		if h.ChecksEnabled > 0 {
			enabled++
		}
	}

	prom.StatsHostsFlapping.Set(flapping)
	prom.StatsHostsChecksEnabled.Set(enabled)
}
