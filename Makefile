.PHONY: help build build-local up down logs as test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t haku1217/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -v -race -shuffle=on ./...

cover: ## Execute tests and Build Coverage File
	go test -cover ./... -coverprofile=cover.out.tmp
	cat cover.out.tmp | grep -v "**_mock.go"  > cover.out
	rm cover.out.tmp
	go tool cover -html=cover.out -o cover.html
	open cover.html

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'