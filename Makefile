#!/usr/bin/make -f

build:
ifeq ($(OS),Windows_NT)
	go build  -o build/lendlord-server.exe .
else
	go build  -o build/lendlord-server .
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

install:
	go build -o lendlord-server && mv lendlord-server $(GOPATH)/bin

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs goimports -w -local github.com/lendlord/lendlord-server

setup: build-linux
	@docker build -t lendlord-server .
	@rm -rf ./build
