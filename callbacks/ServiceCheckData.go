package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
)

func ServiceCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.ServiceCheckData {
		return neb.Error
	}
	host := structs.CastServiceCheck(data)
	if host.Type == neb.ServicecheckInitiate {
		prom.ServiceChecksActive.Add(1)
	}
	if host.Type == neb.ServicecheckProcessed {
		prom.ServiceChecksResults.Add(1)
	}
	return neb.Ok
}
