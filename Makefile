BASE_DIR = $(shell pwd)
POSTGRES_USER ?= $(shell cat .env | grep POSTGRES_USER | cut -d '=' -f 2)
POSTGRES_PASSWORD ?= $(shell cat .env | grep POSTGRES_PASSWORD | cut -d '=' -f 2)
POSTGRES_DB ?= $(shell cat .env | grep POSTGRES_DB | cut -d '=' -f 2)
POSTGRES_PORT ?= $(shell cat .env | grep POSTGRES_PORT | cut -d '=' -f 2)
POSTGRES_HOST ?= $(shell cat .env | grep POSTGRES_HOST | cut -d '=' -f 2)

MIGRATE = docker run --rm -it -v "$(BASE_DIR)/db/migrations":"/migrations" --network host migrate/migrate -path=/migrations/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&timezone=UTC"

N_VERSION ?= $(N)

migrate:
	$(MIGRATE) up $(N_VERSION)

migrate-down:
	$(MIGRATE) down $(N_VERSION)

db-create:
	docker-compose exec \
		-e PGPASSWORD=$(POSTGRES_PASSWORD) \
		db psql -h $(POSTGRES_HOST) -U $(POSTGRES_USER) -c \
		"CREATE DATABASE $(POSTGRES_DB)"

db-close-all-connections:
	docker-compose exec \
		-e PGPASSWORD=$(POSTGRES_PASSWORD) \
		db psql -h $(POSTGRES_HOST) -U $(POSTGRES_USER) -c \
		"SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '#{database}' AND pid <> pg_backend_pid();"

db-drop: db-close-all-connections
	docker-compose exec \
		-e PGPASSWORD=$(POSTGRES_PASSWORD) \
		db psql -h $(POSTGRES_HOST) -U $(POSTGRES_USER) -c \
		"DROP DATABASE $(POSTGRES_DB)"

db-reset: db-drop db-create migrate

.PHONY: migrate migrate-down db-create db-close-all-connections db-drop db-close-all-connections db-drop