COMPOSE_FILE ?= docker-compose.local.yml

compose-up:
	docker compose -f $(COMPOSE_FILE) up -d

compose-down:
	docker compose -f $(COMPOSE_FILE) down

compose-logs:
	docker compose -f $(COMPOSE_FILE) logs -f

test:
	go test -race ./...
lint:
	golangci-lint run
