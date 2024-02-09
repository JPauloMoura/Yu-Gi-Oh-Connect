include env

# Run tests
test:
	@echo "==> running tests"
	@go test -v ./...

# Create Docker network if it doesn't exist
create-network:
	if ! docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 ; then \
		echo "creating network $(NETWORK_NAME)..."; \
		docker network create $(NETWORK_NAME); \
	else \
		echo "$(NETWORK_NAME) network already exists."; \
	fi

# Run infrastructure with Docker Compose
run: create-network
	@echo "==> running infrastructure with docker"
	@docker-compose up

# Run API
run-api:
	@echo "==> running api..."
	@go run ./cmd/main.go

# Create a new migration file
create-migration:
	@migrate create -ext sql -dir infrastructure/database/migrations -seq create_duelist_table

# Run migrations up
migrations-up:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose up

# Rollback migrations
migrations-down:
	@migrate -path infrastructure/database/migrations -database $(DB_CONNECTION_STRING) -verbose down

tidy:
	@go mod tidy
