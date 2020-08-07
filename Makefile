# https://www.cnblogs.com/notokoy/p/11746785.html
GO = GO111MODULE=off go
SERVER_GO_FILES ?=  $(wildcard ./src/server/*.go)
SH_FILES ?= $(shell find ./scripts -name *.sh)
GOVERSION ?= $(shell go version)
BUILDTIME ?= $(shell date +'%Y.%m.%d.%H%M%S')
GITCOMMIT ?= $(shell git log --pretty=oneline -n 1)
BranchInfo ?= $(shell git rev-parse --abbrev-ref HEAD)

LDFlags=" \
    -X 'config.Commit=${GITCOMMIT}' \
    -X 'config.BuildBranch=${BranchInfo}' \
    -X 'config.Buildstamp=${BUILDTIME}' \
    -X 'config.goversion=${GOVERSION}' \
"

.PHONY: all build-server install clean

all: build-server install

build-server:  ##Build pipe server
	@echo "build pipe server"
	$(GO) build -ldflags $(LDFlags)  -o ./bin/pipe-server  $(SERVER_GO_FILES) 

install: ## Installing files to destination path
	@echo "Installing files to destination path"
	@chmod +x ./scripts/install.sh
	@./scripts/install.sh

clean:  ## Clean up intermediate build artifacts.
	@echo "cleaning" 
	@rm -rf ./bin/*
	@rm -rf /usr/local/pipe


