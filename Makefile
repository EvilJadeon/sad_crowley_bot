# Development commands
run:
	if [ -f .env ]; then \
		set -a; \
		source .env; \
		set +a; \
	fi; \
	go run ./cmd/bot

build:
	go build -o crowley ./internal/app/bot

# Docker commands
docker-build:
	docker-compose build --no-cache

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Testing and code quality
test:
	go test -v ./...

lint:
	golangci-lint run

# Cleanup
clean:
	rm -f crowley
	docker-compose down -v
	
.PHONY: run build docker-build docker-up docker-down docker-logs db-start db-stop test lint clean