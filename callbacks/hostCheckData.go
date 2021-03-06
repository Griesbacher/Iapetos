package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/helper"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

const serviceDescription = "host_check"

//HostCheckData is a callback for neb.HostCheckData
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
			prom.HostName:           host.HostName,
			prom.ServiceDescription: serviceDescription,
			prom.CommandName:        host.CommandName,
		}

		setPerformanceData(host.PerfData, helper.CopyLabels(identifier))

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
