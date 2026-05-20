.PHONY: setup start docker-build docker-up docker-down docker-logs docker-clean

setup:
	test -f .env || cp env.example .env

start:
	air

docker-build:
	docker compose build

up: setup
	docker compose up -d

down:
	docker compose down

see-logs:
	docker compose logs -f

clean:
	docker compose down -v --remove-orphans

