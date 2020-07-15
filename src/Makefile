# https://www.cnblogs.com/notokoy/p/11746785.html
GO = GO111MODULE=off go
GO_FILES ?=  $(wildcard config/*.go server/*.go)
SH_FILES ?= $(shell find ./scripts -name *.sh)
GOVERSION ?= $(go version)
BUILDTIME ?= $(date +'%Y.%m.%d.%H%M%S')
GITCOMMIT ?= $(git log --pretty=oneline -n 1)
BranchInfo ?= $(git symbolic-ref HEAD)

.PHONY: all build-server clean

all: build-server

build-server:  ##Build pipe server
	@echo "build pipe server"
	$(GO) build -o ../bin/pipe-server server/log.go server/main.go 

clean:  ## Clean up intermediate build artifacts.
	@echo "cleaning" 
	rm -rf ../bin/*
