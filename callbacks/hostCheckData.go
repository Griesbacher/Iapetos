package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func HostCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.HostCheckData {
		return neb.Error
	}

	host := structs.CastHostCheck(data)

	if host.Type == neb.HostcheckInitiate {
		prom.ChecksActive.Inc()
	} else if host.Type == neb.HostcheckProcessed {
		//Increment global counter
		prom.CheckResults.Inc()

		identifier := prometheus.Labels{
			"host_name":           host.HostName,
			"service_description": "host_check",
			"command_name":        host.CommandName,
		}

		prom.CheckReturnCode.With(identifier).Set(float64(host.ReturnCode))
		prom.CheckExecutionTime.With(identifier).Set(host.ExecutionTime)
		prom.CheckLatency.With(identifier).Set(host.Latency)
		prom.CheckCurrentAttempt.With(identifier).Set(float64(host.CurrentAttempt))
		prom.CheckMaxAttempts.With(identifier).Set(float64(host.MaxAttempts))
		prom.CheckState.With(identifier).Set(float64(host.State))
		prom.CheckStateType.With(identifier).Set(float64(host.StateType))
	}
	return neb.Ok
}
