package main

import (
	"fmt"

	"net"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/callbacks"
	"github.com/griesbacher/Iapetos/config"
	"github.com/griesbacher/Iapetos/cyclic"
	"github.com/griesbacher/Iapetos/logging"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/griesbacher/nagflux/helper"
	"github.com/prometheus/common/version"
)

const configFileKey = "config_file"

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

var prometheusListener net.Listener

var stoppables = []cyclic.Stoppable{}

var iapetosVersion = "1.0"

//This is an example main file, which should demonstrate how to use the library.
func init() {

	//Start selfobserving
	stoppables = append(stoppables, cyclic.StartSelfObserver())

	//Start Host and Service stats
	stoppables = append(stoppables, cyclic.StartHostStatistics())
	stoppables = append(stoppables, cyclic.StartServiceStatistics())

	// just some information about your plugin
	neb.Title = "Iapetos"
	neb.Name = neb.Title
	neb.Desc = "Offers a Prometheus interface for Nagios"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("%s - Build(%s)", iapetosVersion, Build)
	neb.Author = "Philip Griesbacher"

	//Set callbacks
	neb.AddCallback(neb.HostCheckData, callbacks.HostCheckData)
	neb.AddCallback(neb.ServiceCheckData, callbacks.ServiceCheckData)
	neb.AddCallback(neb.NotificationData, callbacks.NotificationCheckData)
	neb.AddCallback(neb.ContactNotificationData, callbacks.ContactNotificationCheckData)

	//Init Hook
	neb.NebModuleInitHook = func(flags int, args string) int {

		neb.CoreFLog("Init - %s - by %s\n", neb.Version, neb.Author)
		neb.CoreFLog("Init flags: %d\n", flags)
		neb.CoreFLog("Init args: %s\n", args)
		argsMap := helper.StringToMap(args, ",", "=")
		if configFile, ok := argsMap[configFileKey]; ok {
			err := config.InitConfig(configFile)
			if err == nil {
				neb.CoreFLog("Loading Configfile: %s\n", args)
				if logging.InitLogDestination() != nil {
					neb.CoreFLog(err.Error())
					return neb.Error
				}
				logging.Flog("Build Info %s\n", version.Info())
				logging.Flog("Build context %s\n", version.BuildContext())
			} else {
				neb.CoreFLog("Could not loaded Configfile: %s, Error: %s\n", args, err.Error())
				return neb.Error
			}
		} else {
			neb.CoreFLog("Could not file Configfile entry(%s) in the given config: %s\n", configFileKey, args)
			return neb.Error
		}
		var err error
		prometheusListener, err = prom.InitPrometheus(config.GetConfig().Prometheus.Address)
		if err == nil {
			logging.Flog("Starting Prometheus at %s\n", config.GetConfig().Prometheus.Address)
			return neb.Ok
		}
		logging.Flog("Could not starting Prometheus at %s. Error: %s\n", config.GetConfig().Prometheus.Address, err)
		return neb.Error
	}

	//Deinit Hook
	neb.NebModuleDeinitHook = func(flags, reason int) int {
		logging.Flog("Deinit\n")
		logging.Flog("Deinit flags: %d\n", flags)
		logging.Flog("Deinit reason: %d\n", reason)

		for _, s := range stoppables {
			s.Stop()
		}

		prometheusListener.Close()

		return neb.Ok
	}
}

//DON'T USE MAIN, IT WILL NEVER BE CALLED! USE CALLBACKS.
func main() {}
