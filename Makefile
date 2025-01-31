CMD := ./cmd
BINARY := bla.o

all: run_build

format:
	@go fmt ./...

test:
	@go test ./...

build:
	@go build -o=/tmp/bin/${BINARY} ${CMD}

run:
	@go run ${CMD}

run_build: build
	@/tmp/bin/${BINARY}

.PHONY: all format build run run_build test
