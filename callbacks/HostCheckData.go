package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
)

func HostCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.HostCheckData {
		return neb.Error
	}
	host := structs.CastHostCheck(data)
	if host.Type == neb.HostcheckInitiate {
		prom.HostChecksActive.Add(1)
	}
	if host.Type == neb.HostcheckProcessed {
		prom.HostResults.Add(1)
	}
	return neb.Ok
}
