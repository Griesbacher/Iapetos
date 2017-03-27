#!/usr/bin/make -f

MAKE:=make
SHELL:=bash
GOVERSION:=$(shell go version | awk '{print $$3}' | sed 's/^go\([0-9]\.[0-9]\).*/\1/')

all: deps fmt cyclo misspell build_all

deps: versioncheck

updatedeps: versioncheck

build_naemon: fmt cyclo misspell
	go build -tags naemon -buildmode=c-shared -ldflags "-s -w -X main.Build=$(shell git rev-parse --short HEAD)" -o iapetos_naemon

build_nagios3:
	go build -tags nagios3 -buildmode=c-shared -ldflags "-s -w -X main.Build=$(shell git rev-parse --short HEAD)" -o iapetos_nagios3

build_nagios4:
	go build -tags nagios4 -buildmode=c-shared -ldflags "-s -w -X main.Build=$(shell git rev-parse --short HEAD)" -o iapetos_nagios4

build_all: build_naemon build_nagios3 build_nagios4


debugbuild: deps fmt
	go build -buildmode=c-shared -race -ldflags "-X main.Build=$(shell git rev-parse --short HEAD)"

test: fmt 
	go test -short -v
	if grep -r TODO: *.go; then exit 1; fi

citest: deps
	#
	# Normal test cases
	#
	go test -v
	#
	# Benchmark tests
	#
	go test -v -bench=B\* -run=^$$ . -benchmem
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
	# go get -u golang.org/x/tools/cmd/goimports
	goimports -w .
	go tool vet -all -shadow -assign -atomic -bool -composites -copylocks -nilfunc -rangeloops -unsafeptr -unreachable .
	gofmt -w -s .

versioncheck:
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
