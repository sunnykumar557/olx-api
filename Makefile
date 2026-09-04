.PHONY: build run

build:
	@go build -o bin/api ./cmd/api/main.go

run: build
	@./bin/api

migrate-up:
	@go run ./cmd/migrate/main.go up

migrate-down:
	@go run ./cmd/migrate/main.go down