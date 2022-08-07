.PHONY: build up down logs ps test

DOCKER_TAG := latest

build:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

ps:
	docker-compose ps

test:
	go test -race -shuffle=on ./...
