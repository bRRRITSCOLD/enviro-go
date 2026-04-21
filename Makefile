.PHONY: setup build test lint

setup:
	go install -modfile=tools/go.mod tool
	lefthook install

build:
	go build ./...

test:
	go test -v -race -count=1 ./...

lint:
	golangci-lint run
