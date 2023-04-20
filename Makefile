include .env
export $(shell sed 's/=.*//' .env)


init-local: pg

DOCKER_COMPOSE_BIN := docker-compose
COMPOSE := ${DOCKER_COMPOSE_BIN} -p tn-sample -f docker-compose.yaml

.PHONY: pg
pg:
	${COMPOSE} up -d pg

.PHONY: pg-down
pg-down:
	${COMPOSE} down

.PHONY: migrate
migrate:
	docker exec -it tn-samples-pg "pg/migrate.sh"

.PHONY: run
run:
	go run src/cmd/main.go
