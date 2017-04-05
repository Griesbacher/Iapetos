package cyclic

import (
	"time"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
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
			meta := hosts.GenMetaHostAndServiceList()
			countTypes(meta, prom.StatsServicesCheckType)
			flapping, enabled := countMinorStats(meta)
			prom.StatsServicesFlapping.Set(flapping)
			prom.StatsServicesChecksEnabled.Set(enabled)
		}
	}
}
