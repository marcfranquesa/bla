CMD := ./cmd/bla
BINARY := bla.o

all:

format:
	@go fmt ./...

test:
	@go test ./...

mod:
	@docker exec -it bla-moderation /tui

.PHONY: all format test mod
