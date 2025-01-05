#!/usr/bin/make -f
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
BINDIR ?= $(GOPATH)/bin

# for dockerized protobuf tools
DOCKER := $(shell which docker)
HTTPS_GIT := github.com/onsonr/hway.git
export RELEASE_DATE="$(date +%Y).$(date +%V).$(date +%u)"

all: install test

build: go.sum
	go build -o bin/hway .

install: go.sum
	go install -mod=readonly .

########################################
### Tools & dependencies
########################################
pkl-gen:
	@go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest
	@pkl-gen-go ./pkl/App.pkl

sqlc-gen:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest
	@cd internal && sqlc generate

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go install github.com/RobotsAndPencils/goviz@latest
	@goviz -i ./cmd/sonrd -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf .aider*
	rm -rf static
	rm -rf .out
	rm -rf hway.db
	rm -rf snapcraft-local.yaml build/
	rm -rf build

distclean: clean
	rm -rf vendor/


###############################################################################
###                                     help                                ###
###############################################################################

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Available targets:"
	@echo "  install             : Install the binary"
	@echo "  local-image         : Install the docker image"
	@echo "  proto-gen           : Generate code from proto files"
	@echo "  testnet             : Local devnet with IBC"
	@echo "  sh-testnet          : Shell local devnet"
	@echo "  ictest-basic        : Basic end-to-end test"
	@echo "  ictest-ibc          : IBC end-to-end test"

.PHONY: help
