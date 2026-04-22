.PHONY: setup build test lint

setup:
	go install -modfile=tools/go.mod tool
	lefthook install

build:
	go build ./...

test:
	go run ./scripts/test/main.go

lint:
	golangci-lint run
