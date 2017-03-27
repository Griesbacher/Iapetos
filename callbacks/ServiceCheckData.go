package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/davecgh/go-spew/spew"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

func ServiceCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.ServiceCheckData {
		return neb.Error
	}

	service := structs.CastServiceCheck(data)

	if service.Type == neb.ServicecheckInitiate {
		prom.ChecksActive.Inc()
	} else if service.Type == neb.ServicecheckProcessed {
		spew.Dump("---------")
		//spew.Dump(service)
		//Increment global counter
		prom.CheckResults.Inc()

		identifier := prometheus.Labels{
			"host_name":           service.HostName,
			"service_description": service.ServiceDescription,
			"command_name":        service.CommandName,
		}

		prom.CheckReturnCode.With(identifier).Set(float64(service.ReturnCode))
		prom.CheckExecutionTime.With(identifier).Set(service.ExecutionTime)
		prom.CheckLatency.With(identifier).Set(service.Latency)
		prom.CheckCurrentAttempt.With(identifier).Set(float64(service.CurrentAttempt))
		prom.CheckMaxAttempts.With(identifier).Set(float64(service.MaxAttempts))
		prom.CheckState.With(identifier).Set(float64(service.State))
		prom.CheckStateType.With(identifier).Set(float64(service.StateType))
	}
	return neb.Ok
}
