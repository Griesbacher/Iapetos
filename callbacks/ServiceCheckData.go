package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/helper"
	"github.com/griesbacher/Iapetos/prom"
)

func ServiceCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.ServiceCheckData {
		return neb.Error
	}
	service := structs.CastServiceCheck(data)
	if service.Type == neb.ServicecheckInitiate {
		prom.ServiceChecksActive.Inc()
	}
	if service.Type == neb.ServicecheckProcessed {
		//Increment global counter
		prom.ServiceChecksResults.Inc()

		//Increment returncode counter
		prom.ServiceCheckReturnCode.
			With(map[string]string{"code": helper.ReturnCodeToString(service.ReturnCode)}).
			Inc()
	}
	return neb.Ok
}
