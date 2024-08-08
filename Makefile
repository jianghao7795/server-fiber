.PHONY: build dev start clean tool help

all: build

build:
	go build -v .

start:
	air

dev:
	go run cmd/main.go -c config.yaml

local:
	go run . -c config.yaml

tool:
	go tool vet . |& grep -v vendor; true
	gofmt -w .

clean:
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make dev: go run ."
	@echo "make clean: remove object files and cached files"
