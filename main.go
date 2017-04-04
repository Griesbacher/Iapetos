package main

import (
	"fmt"

	"net"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/griesbacher/Iapetos/callbacks"
	"github.com/griesbacher/Iapetos/config"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/griesbacher/nagflux/helper"
)

const configFileKey = "config_file"

// Build contains the current git commit id
// compile passing -ldflags "-X main.Build <build sha1>" to set the id.
var Build string

var prometheusListener net.Listener

//This is an example main file, which should demonstrate how to use the library.
func init() {

	//Start selfobserving
	callbacks.StartSelfObserver()

	// just some information about your plugin
	neb.Title = "Iapetos"
	neb.Name = neb.Title
	neb.Desc = "Offers a Prometheus interface for Nagios"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("0.1 - %s", Build)
	neb.Author = "Philip Griesbacher"

	//Set callbacks
	neb.AddCallback(neb.HostCheckData, callbacks.HostCheckData)
	neb.AddCallback(neb.ServiceCheckData, callbacks.ServiceCheckData)
	neb.AddCallback(neb.NotificationData, callbacks.NotificationCheckData)
	neb.AddCallback(neb.ContactNotificationData, callbacks.ContactNotificationCheckData)

	//Init Hook
	neb.NebModuleInitHook = func(flags int, args string) int {
		neb.CoreFLog("Init\n")
		neb.CoreFLog("Init flags: %d\n", flags)
		neb.CoreFLog("Init args: %s\n", args)
		neb.CoreFLog("CoreType %s\n", neb.CoreToString())
		argsMap := helper.StringToMap(args, ",", "=")
		if configFile, ok := argsMap[configFileKey]; ok {
			err := config.InitConfig(configFile)
			if err == nil {
				neb.CoreFLog("Loading Configfile: %s\n", args)
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
			neb.CoreFLog("Starting Prometheus at %s", config.GetConfig().Prometheus.Address)
			return neb.Ok
		}
		neb.CoreFLog("Could not starting Prometheus at %s. Error: %s", config.GetConfig().Prometheus.Address, err)
		return neb.Error
	}

	//Deinit Hook
	neb.NebModuleDeinitHook = func(flags, reason int) int {
		neb.CoreFLog("Deinit\n", neb.Title)
		neb.CoreFLog("Deinit flags: %d\n", neb.Title, flags)
		neb.CoreFLog("Deinit reason: %d\n", neb.Title, reason)

		callbacks.StopSelfObserver()
		prometheusListener.Close()

		return neb.Ok
	}
}

//DON'T USE MAIN, IT WILL NEVER BE CALLED! USE CALLBACKS.
func main() {}
