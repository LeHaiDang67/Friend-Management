.PHONY: db

APP_NAME := friend
APP_PATH := /$(APP_NAME)

COMPOSE = docker-compose -f docker-compose.yml

run: volumes db sleep migrate start

volumes:
	$(COMPOSE) up -d alpine
	docker cp ./data/migrations/. alpine-friend-local:/migrations
	docker cp $(shell pwd)/. alpine-friend-local:$(APP_PATH)

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