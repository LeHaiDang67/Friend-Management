.PHONY: db

APP_NAME := friend
APP_PATH := /$(APP_NAME)

COMPOSE = docker-compose -f docker-compose.yml

run: volumes db sleep migrate start

volumes:
	$(COMPOSE) up -d alpine
	docker cp ./data/migrations/. alpine-$(APP_NAME)-$${CONTAINER_SUFFIX:-local}:/migrations
	docker cp $(shell pwd)/. alpine-$(APP_NAME)-$${CONTAINER_SUFFIX:-local}:$(APP_PATH)

db:
	$(COMPOSE) up -d db

migrate: MOUNT_VOLUME = -v $(shell pwd)/data/migrations:/migrations
migrate:
	$(COMPOSE) run --rm $(MOUNT_VOLUME) db-migrate \
	sh -c './migrate -path /migrations -database $$DATABASE_URL up'

api:
	$(COMPOSE) up friend-api

start:
	go run -mod=readonly cmd/main.go

sleep:
	sleep 5

# sqlboiler:
# 	sqlboiler --output models/orm --pkgname orm psql

# boil:
# 	sqlboiler --output internal/models2 --pkgname models2 postgres

# boilerplate generates DB and GraphQL boilerplate code
boilerplate: db sleep migrate generate-models

generate-models: SQLBOILER_GOMOD := $(shell grep sqlboiler go.mod)
generate-models: SQLBOILER_VER := $(word 2,$(strip $(SQLBOILER_GOMOD)))
generate-models:
	$(COMPOSE) run --rm runner \
		sh -c 'mkdir -p /sqlboiler && \
			cd /sqlboiler && \
			go mod init example.com/sqlboiler && \
			go get github.com/volatiletech/sqlboiler@$(SQLBOILER_VER) && \
			go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@$(SQLBOILER_VER) && \
			go get github.com/sqs/goreturns && \
			cd /app && \
			sqlboiler --no-tests --output internal/models --pkgname models psql && \
			goreturns -w internal/models/*.go'

generate-gqlgen:
	go run github.com/99designs/gqlgen -c cmd/graphql/gqlgen.yml