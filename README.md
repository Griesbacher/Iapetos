[![CircleCI](https://circleci.com/gh/Griesbacher/Iapetos.svg?style=shield)](https://circleci.com/gh/Griesbacher/Iapetos)
[![Go Report Card](https://goreportcard.com/badge/github.com/Griesbacher/Iapetos)](https://goreportcard.com/report/github.com/Griesbacher/Iapetos)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

# Iapetos
This is a Nagios / Naemon Prometheus exporter, which gathers information about Nagios and the checks it executes. It's build for the NEB Interface, if the core supports that, it should work.

# Issue
Currently it only works with Naemon. Nagios3 freezes after a few seconds: [ConSol/go-neb-wrapper#1](https://github.com/ConSol/go-neb-wrapper/issues/1)

# Installation
## Requirements
- CGo (tested with 1.7+, but could also work with older versions)
### Nagios3 / Nagios4
- Headerfiles are included
### Naemon
- Naemon dev package, for headerfiles, see http://www.naemon.org

## Building
- make build_naemon
- make build_nagios3
- make build_nagios4

If no 'make' is available have a look at the Makefile, it's just a shortcut for some go commands

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

## Core
Pass the config as parameter, the rest is like every other NEB Module.
```
broker_module=/path/to/your/bin/iapetos_naemon config_file=/path/to/your/config/config.yaml
```

# Overview
![Overview](https://github.com/Griesbacher/Iapetos/blob/master/doc/Componentdiagram.bmp)

# Grafana Examples
## Overview
![Overview1](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview1.PNG)
![Overview2](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview2.PNG)
![Overview3](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Overview3.PNG)

## Check
![Check](https://github.com/Griesbacher/Iapetos/blob/master/doc/Grafana%20-%20Check.PNG)
