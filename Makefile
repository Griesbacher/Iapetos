#!/usr/bin/make -f

MAKE:=make
SHELL:=bash
GOVERSION:=$(shell go version | awk '{print $$3}' | sed 's/^go\([0-9]\.[0-9]\).*/\1/')
BUILD:=$(shell git rev-parse --short HEAD)
PROMLDFLAGS:= \
-X github.com/prometheus/common/version.Version=$(shell grep "var iapetosVersion" main.go | cut -d \" -f2) \
-X github.com/prometheus/common/version.Revision=$(BUILD) \
-X github.com/prometheus/common/version.Branch=$(shell git branch | cut -c 3-) \
-X github.com/prometheus/common/version.BuildUser=Iapetos \
-X github.com/prometheus/common/version.BuildDate=$(shell date -u '+%Y-%m-%d_%H:%M:%S%p') \

all: deps fmt cyclo misspell build_all

deps: versioncheck

updatedeps: versioncheck

build_naemon: fmt cyclo misspell
	go build -tags naemon -buildmode=c-shared -ldflags "-s -w -X main.Build=$(BUILD) $(PROMLDFLAGS)" -o iapetos_naemon

build_nagios3: fmt cyclo misspell
	go build -tags nagios3 -buildmode=c-shared -ldflags "-s -w -X main.Build=$(BUILD) $(PROMLDFLAGS)" -o iapetos_nagios3

build_nagios4: fmt cyclo misspell
	go build -tags nagios4 -buildmode=c-shared -ldflags "-s -w -X main.Build=$(BUILD) $(PROMLDFLAGS)" -o iapetos_nagios4

build_all: build_naemon build_nagios3 build_nagios4


debugbuild: deps fmt
	go build -buildmode=c-shared -race -ldflags "-X main.Build=$(shell git rev-parse --short HEAD)"

test: fmt 
	go test -v -tags naemon
	go test -v -tags nagios3
	go test -v -tags nagios4

citest: deps
	$(MAKE) build_all
	$(MAKE) test
	#
	# Checking gofmt errors
	#
	if [ $$(gofmt -s -l . | wc -l) -gt 0 ]; then \
		echo "found format errors in these files:"; \
		gofmt -s -l .; \
		exit 1; \
	fi
	#
	# Checking TODO items
	#
	if grep -r TODO: *.go; then exit 1; fi
	$(MAKE) lint
	$(MAKE) cyclo
	$(MAKE) misspell
	#
	# All CI tests successful
	#
	curl 'https://goreportcard.com/checks' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Accept-Encoding: gzip, deflate, br'  -H 'Connection: keep-alive' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'DNT: 1' -H 'Host: goreportcard.com' -H 'Referer: https://goreportcard.com/report/github.com/Griesbacher/Iapetos' --data 'repo=github.com%2FGriesbacher%2FIapetos'

benchmark: fmt
	go test -v -bench=B\* -run=^$$ . -benchmem

racetest: fmt
	go test -race -v

covertest: fmt
	go test -v -coverprofile=cover.out
	go tool cover -func=cover.out
	go tool cover -html=cover.out -o coverage.html

coverweb: fmt
	go test -v -coverprofile=cover.out
	go tool cover -html=cover.out

clean:
	# TODO: ...
	rm -f cover.out
	rm -f coverage.html

fmt:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .
	go tool vet -all -shadow -assign -atomic -bool -composites -copylocks -nilfunc -rangeloops -unsafeptr -unreachable .
	gofmt -w -s .

versioncheck:
	echo "**** Go Version: $$(go version)"
	@[ "$(GOVERSION)" = "devel" ] || [ $$(echo "$(GOVERSION)" | tr -d ".") -ge 15 ] || { \
		echo "**** ERROR:"; \
		echo "**** Iapetos requires at least golang version 1.5 or higher"; \
		echo "**** this is: $$(go version)"; \
		exit 1; \
	}

lint:
	#
	# Check if golint complains
	# see https://github.com/golang/lint/ for details.
	# Only works with Go 1.6 or up.
	#
	@( [ "$(GOVERSION)" != "devel" ] && [ $$(echo "$(GOVERSION)" | tr -d ".") -lt 16 ] ) || { \
		go get github.com/golang/lint/golint; \
		golint -set_exit_status ./...; \
	}

cyclo:
	go get github.com/fzipp/gocyclo
	#
	# Check if there are any too complicated functions
	# Any function with a score higher than 15 is bad.
	# See https://github.com/fzipp/gocyclo for details.
	#
	gocyclo -over 15 .

misspell:
	go get github.com/client9/misspell/cmd/misspell
	#
	# Check if there are common spell errors.
	# See https://github.com/client9/misspell
	#
	misspell -error .

version:
	OLDVERSION="$(shell grep "VERSION =" main.go | awk '{print $$3}' | tr -d '"')"; \
	NEWVERSION=$$(dialog --stdout --inputbox "New Version:" 0 0 "v$$OLDVERSION") && \
		NEWVERSION=$$(echo $$NEWVERSION | sed "s/^v//g"); \
		if [ "v$$OLDVERSION" = "v$$NEWVERSION" -o "x$$NEWVERSION" = "x" ]; then echo "no changes"; exit 1; fi; \
		sed -i -e 's/VERSION =.*/VERSION = "'$$NEWVERSION'"/g' main.go
