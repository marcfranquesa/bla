CMD := ./cmd/bla
BINARY := bla.o

all: deploy

format:
	@go fmt ./...

test:
	@go test ./...

deploy:
	@./scripts/deploy.mysql.moderation.sh

mod:
	@docker exec -it bla-moderation /tui

data:
	./scripts/add-dummy-data.sh

.PHONY: all format test deploy mod data
