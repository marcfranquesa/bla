CMD := ./cmd/bla
BINARY := bla.o

all: deploy

format:
	@go fmt ./...

test-sql:
	@sqlfluff lint db/init --dialect mysql

test-go:
	@if [ "$$(gofmt -l . | wc -l)" -gt 0 ]; then exit 1; fi
	@go test ./...

test: test-sql test-go

deploy:
	@./scripts/deploy.mysql.moderation.sh

clean:
	@docker-compose down
	@docker volume rm bla_data

mod:
	@docker exec -it bla-moderation /tui

data:
	@./scripts/add-dummy-data.sh

.PHONY: all format test-sql test-go test deploy clean mod data
