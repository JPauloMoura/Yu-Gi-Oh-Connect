include .env
test:
	@echo "==> running tests"
	@go test -v ./...

create-network:
	if ! docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 ; then \
		echo "creating network $(NETWORK_NAME)..."; \
		docker network create $(NETWORK_NAME); \
	else \
		echo "$(NETWORK_NAME) network already exists."; \
	fi

run: create-network
	@echo "==> running infrastructure with docker"
	@docker-compose up

run-api:
	@echo "==> running api..."
	@go run ./cmd/main.go

run-api-with-air:
	@echo "==> running api..."
	@air

create-migration:
	@migrate create -ext sql -dir infrastructure/database/migrations -seq create_duelist_table

migrations-up:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose up

migrations-down:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose down