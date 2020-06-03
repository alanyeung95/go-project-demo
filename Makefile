.DEFAULT_GOAL := run

MAIN := app

.PHONY: build
build:
	@echo Building binaries
	sh ./scripts/build.sh

# MongoDB 
MONGO_VOLUME := mongo
MONGO_VOLUME_EXIST = $(shell docker volume ls -f name=$(MONGO_VOLUME) -q | wc -l | tr -d ' \t\n\r\f')
.PHONY: mongo
mongo:
	@echo ">>> Ensure $(MONGO_VOLUME) volume"
ifneq ($(MONGO_VOLUME_EXIST), 1)
	docker volume create $(MONGO_VOLUME)
endif

.PHONY: mongo-up
mongo-up: mongo
	@echo ">>> compose-up Go Project Demo MongoDB"
	docker-compose -f docker-compose.mongo.yml up -d

.PHONY: mongo-down
mongo-down:
	@echo ">>> compose-down Go Project Demo MongoDB"
	docker-compose -f docker-compose.mongo.yml down

### Development
.PHONY: up
up:  mongo-up
	@echo ">>> Ensure compose up"
	docker-compose up --build -d

.PHONY: run
run: up
	@echo ">>> run"
	docker-compose exec $(MAIN) sh -c  './scripts/start.sh'