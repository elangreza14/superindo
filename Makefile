#!make
include .env
	
run-http:
	go run cmd/rest/main.go

FILENAME?=file-name

migrate:
	@read -p  "up or down or version? " MODE; \
	migrate -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}" -path ${MIGRATION_FOLDER} $$MODE
	
migrate-create:
	@read -p  "What is the name of migration? " NAME; \
	migrate create -ext sql -tz Asia/Jakarta -dir ${MIGRATION_FOLDER} -format "20060102150405" $$NAME

gen:
	go generate ./...

test:
	go test ./...

up:
	docker compose up -d

down:
	docker compose down

seed:
	cat ./db/seed/seed_1.sql | docker exec -i superindo-database psql -h localhost -U superindo -f-

.PHONY: migrate migrate-create run-http gen up down seed