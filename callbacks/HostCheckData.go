package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/davecgh/go-spew/spew"
	"github.com/griesbacher/Iapetos/helper"
	"github.com/griesbacher/Iapetos/prom"
)

func HostCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.HostCheckData {
		return neb.Error
	}
	host := structs.CastHostCheck(data)

	spew.Dump("---------")
	spew.Dump(host)
	if host.Type == neb.HostcheckInitiate {
		prom.HostChecksActive.Inc()
	}
	if host.Type == neb.HostcheckProcessed {
		//Increment global counter
		prom.HostResults.Inc()

		//Increment returncode counter
		prom.HostCheckReturnCode.
			With(map[string]string{"code": helper.ReturnCodeToString(host.ReturnCode)}).
			Inc()
	}
	return neb.Ok
}
