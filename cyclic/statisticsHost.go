package cyclic

import (
	"time"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/hostStates"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
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
			meta := hosts.GenMetaHostAndServiceList()
			countTypes(meta, prom.StatsHostsCheckType)
			countStates(meta, prom.StatsHostsStateType, map[string]float64{
				hostStates.StateTypeToString(hostStates.Up):          0,
				hostStates.StateTypeToString(hostStates.Down):        0,
				hostStates.StateTypeToString(hostStates.Unreachable): 0,
			}, hostStates.StateTypeToString)
			countMinorStats(meta,
				prom.StatsHostsFlapping,
				prom.StatsHostsChecksEnabled,
				prom.StatsHostsFlexDowntime,
				prom.StatsHostsDowntime,
			)
		}
	}
}
