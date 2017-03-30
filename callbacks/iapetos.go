package callbacks

import (
	"time"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

var stop chan bool

func StartSelfObserver() {
	if stop != nil {
		return
	} else {
		stop = make(chan bool)
	}
	go func() {
		stats := &neb.Statistics{
			RegisteredCallbacksByType: make(chan map[int]int, 1000),
			OverallCallbackDuration:   make(chan map[int]time.Duration, 1000),
		}
		neb.Stats = stats
	Loop:
		for {
			select {
			case <-stop:
				break Loop
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
		logging.GetLogger().Info("Stopping SelfObserver")
		stop <- true
	}()
}

func StopSelfObserver() {
	if stop == nil {
		return
	}
	stop <- true
	<-stop
}
