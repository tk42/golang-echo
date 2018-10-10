# Docker-compose command
DOCKERCMD=docker-compose up -d;docker-compose exec -T app make

# Go Parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v -cover
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=golang-echo
OUTPUT_DIR=bin

all: _clean _build _test _install

.PHONY: setenv
setenv :
	. ./setenv.sh

.PHONY: clean
clean :
	$(DOCKERCMD) _clean

.PHONY: _clean
_clean :
	$(GOCLEAN)
	@rm -rf $(OUTPUT_DIR)

format :
	go fmt ./$(BINARY_NAME)/...

.PHONY: build
build :
	$(DOCKERCMD) _build

.PHONY: _build
_build : format
	$(GOBUILD) -v ./$(BINARY_NAME)/...

.PHONY: test
test :
	$(DOCKERCMD) _test

.PHONY: _test
_test : format
	$(GOTEST) ./$(BINARY_NAME)/...

.PHONY: _install
_install :
	@. ./setenv.sh; $(GOINSTALL) ./$(BINARY_NAME)/...
	@echo "Install successful!"
