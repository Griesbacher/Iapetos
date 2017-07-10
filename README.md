[![CircleCI](https://circleci.com/gh/Griesbacher/Iapetos.svg?style=shield)](https://circleci.com/gh/Griesbacher/Iapetos)
[![Go Report Card](https://goreportcard.com/badge/github.com/Griesbacher/Iapetos)](https://goreportcard.com/report/github.com/Griesbacher/Iapetos)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

# Iapetos
This is a Naemon / Nagios(3/4) / Icinga Prometheus exporter, which gathers information about Nagios and the checks it executes. It's build for the NEB Interface, if the core supports that, it should work.

## Supported Cores
- Naemon
- Nagios3 without daemon mode!
- Nagios4
- Icinga without daemon mode! (Uses the Nagios3 files)

# Issue
Nagios3 / Icinga freezes after a few seconds if run in daemon mode: [ConSol/go-neb-wrapper#1](https://github.com/ConSol/go-neb-wrapper/issues/1)
There is a fork happening within the core after the module has been loaded, which leads to the problem that Go will not start goroutines anymore.

# Installation
## Requirements if building from source
- CGo (tested with 1.7+, but could also work with older versions)
### Nagios3 / Nagios4
- Headerfiles are included
### Naemon
- Naemon dev package, for headerfiles, see www.naemon.org

## Building
- make build_naemon
- make build_nagios3
- make build_nagios4

If no 'make' is available have a look at the Makefile, it's just a shortcut for some go commands

## Pre-Build Binaries
If a CI test went well, there is also a binary on circleci, like this [one](https://circleci.com/gh/Griesbacher/Iapetos/18#artifacts/containers/0) but you have to change to the current build. 

The further releases will also contain pre-build x64 binaries.

# Configuration
## Iapetos
Nothing special here, just the log destination and the port.
```YAML
logging:
# supported targets: core (core logfile), stdout
  destination : "core"

prometheus:
  address: ":9245"
```
Save this as config.yaml

## Core
Pass the config as parameter, the rest is like every other NEB Module.
```
broker_module=/path/to/your/bin/iapetos_naemon config_file=/path/to/your/config/config.yaml
```

# Overview
![Overview](https://github.com/Griesbacher/Iapetos/blob/master/doc/Componentdiagram.bmp)

# Grafana Examples
## Overview
[JSON Dashboard](https://github.com/Griesbacher/Iapetos/blob/master/grafana_dashboards/Iapetos%20Stats.json)

![Overview1](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview1.PNG)
![Overview2](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview2.PNG)
![Overview3](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview3.PNG)

## Check
[JSON Dashboard](https://github.com/Griesbacher/Iapetos/blob/master/grafana_dashboards/Nagios%20Check%20Data.json)

![Check](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Check.PNG)
