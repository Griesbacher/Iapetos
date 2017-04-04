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

//ServiceStatistics monitors the core host stats
type ServiceStatistics struct {
	Stoppable
}

//StartServiceStatistics creates an new ServiceStatistics and starts it
func StartServiceStatistics() Stoppable {
	s := ServiceStatistics{Stoppable{stop: make(chan bool)}}
	go s.run()
	return s.Stoppable
}

func (s ServiceStatistics) run() {
	for {
		select {
		case <-s.stop:
			logging.GetLogger().Info("Stopping ServiceStatistics")
			s.stop <- true
			return
		case <-time.After(time.Duration(10) * time.Second):
			hosts := neb.GetServices()
			if len(hosts) == 0 {
				continue
			}
			prom.StatsServicesAmount.Set(float64(len(hosts)))
			countServiceTypes(hosts)
			countMinorServiceStats(hosts)
		}
	}
}

func countServiceTypes(hosts structs.Servicelist) {
	counterMap := map[string]float64{}
	for _, h := range hosts {
		t := checkTypes.CheckTypeToString(h.CheckType)
		if _, contained := counterMap[t]; !contained {
			counterMap[t] = 0
		}
		counterMap[t]++
	}
	for k, v := range counterMap {
		prom.StatsServicesCheckType.With(prometheus.Labels{
			prom.Type: k,
		}).Set(v)
	}
}

func countMinorServiceStats(hosts structs.Servicelist) {
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

	prom.StatsServicesFlapping.Set(flapping)
	prom.StatsServicesChecksEnabled.Set(enabled)
}
