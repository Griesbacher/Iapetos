package cyclic

import (
	"time"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

//SelfObserver monitors the neb callbacks
type SelfObserver struct {
	Stoppable
}

func (s SelfObserver) run() {
	stats := &neb.Statistics{
		RegisteredCallbacksByType: make(chan map[int]int, 1000),
		OverallCallbackDuration:   make(chan map[int]time.Duration, 1000),
	}
	neb.Stats = stats
	for {
		select {
		case <-s.stop:
			logging.Flog("Stopping SelfObserver\n")
			s.stop <- true
		case registeredCallbacks := <-stats.RegisteredCallbacksByType:
			for callbackType, amount := range registeredCallbacks {
				prom.ModuleCallbacks.
					With(prometheus.Labels{"type": neb.CallbackTypeToString(callbackType)}).
					Set(float64(amount))
			}
		case callbackDuration := <-stats.OverallCallbackDuration:
			for callbackType, amount := range callbackDuration {
				prom.ModuleDuration.
					With(prometheus.Labels{"type": neb.CallbackTypeToString(callbackType)}).
					Observe(amount.Seconds())

			}
		}
	}
}

//StartSelfObserver creates an new SelfObserver and starts it
func StartSelfObserver() Stoppable {
	s := SelfObserver{Stoppable{stop: make(chan bool)}}
	go s.run()
	return s.Stoppable
}
