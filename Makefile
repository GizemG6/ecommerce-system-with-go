.PHONY: build up run test


build:
	docker compose build


up:
	docker compose up -d


run:
	go run ./cmd/api


test:
	go test ./... -v