package logging

import (
	"fmt"
	"strings"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/config"
)

const (
	core   = "core"
	stdout = "stdout"
)

var destination string

//InitLogDestination checks the config values and inits the logging
func InitLogDestination() error {
	dest := strings.ToLower(config.GetConfig().Logging.Destination)
	if dest != core && dest != stdout {
		return fmt.Errorf("This log destination is not supported. Supported are: %s %s", core, stdout)
	}

	neb.CoreFLog("Logging from now on to: %s", dest)
	destination = dest
	return nil
}

//Flog can be used to log ether to the core or stdout
func Flog(format string, a ...interface{}) {
	if destination == core {
		neb.CoreFLog(format, a)
	} else if destination == stdout {
		if a == nil {
			fmt.Print(format)
		} else {
			fmt.Printf(format, a)
		}
	}
}
