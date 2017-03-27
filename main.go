package main

import (
	"fmt"

	"net"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/nlog"
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
	// just some information about your plugin
	neb.Title = "Iapetos"
	neb.Name = neb.Title
	neb.Desc = "Offers a Prometheus interface for Nagios"
	neb.License = "GPL v3"
	neb.Version = fmt.Sprintf("0.1 - %s", Build)
	neb.Author = "Philip Griesbacher"

	//There can be multiple of them
	neb.AddCallback(neb.HostCheckData, callbacks.HostCheckData)

	//Init Hook Example
	neb.NebModuleInitHook = func(flags int, args string) int {
		nlog.CoreLog(fmt.Sprintf("[%s] Init\n", neb.Title))
		nlog.CoreLog(fmt.Sprintf("[%s] Init flags: %d\n", neb.Title, flags))
		nlog.CoreLog(fmt.Sprintf("[%s] Init args: %s\n", neb.Title, args))
		argsMap := helper.StringToMap(args, ",", "=")
		if configFile, ok := argsMap[configFileKey]; ok {
			err := config.InitConfig(configFile)
			if err == nil {
				nlog.CoreLog(fmt.Sprintf("[%s] Loading Configfile: %s\n", neb.Title, args))
			} else {
				nlog.CoreLog(fmt.Sprintf("[%s] Could not loaded Configfile: %s, Error: %s\n", neb.Title, args, err.Error()))
				return neb.Error
			}
		} else {
			nlog.CoreLog(fmt.Sprintf("[%s] Could not file Configfile entry(%s) in the given config: %s\n", neb.Title, configFileKey, args))
			return neb.Error
		}
		var err error
		prometheusListener, err = prom.InitPrometheus(config.GetConfig().Prometheus.Address)
		if err == nil {
			nlog.CoreLog(fmt.Sprintf("[%s] Starting Prometheus at %s", neb.Title, config.GetConfig().Prometheus.Address))
			return neb.Ok
		}
		nlog.CoreLog(fmt.Sprintf("[%s] Could not starting Prometheus at %s. Error: %s", neb.Title, config.GetConfig().Prometheus.Address, err))
		return neb.Error
	}

	//Deinit Hook Example
	neb.NebModuleDeinitHook = func(flags, reason int) int {
		nlog.CoreLog(fmt.Sprintf("[%s] Deinit\n", neb.Title))
		nlog.CoreLog(fmt.Sprintf("[%s] Deinit flags: %d\n", neb.Title, flags))
		nlog.CoreLog(fmt.Sprintf("[%s] Deinit reason: %d\n", neb.Title, reason))
		prometheusListener.Close()
		return neb.Ok
	}

}

//DON'T USE MAIN, IT WILL NEVER BE CALLED! USE CALLBACKS.
func main() {}
