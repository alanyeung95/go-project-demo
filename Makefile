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

#.PHONY: mongo-up
#mongo-up: mongo
#	@echo ">>> compose-up Go Project Demo MongoDB"
#	docker-compose -f docker-compose.mongo.yml up -d

#.PHONY: mongo-down
#mongo-down:
#	@echo ">>> compose-down Go Project Demo MongoDB"
#	docker-compose -f docker-compose.mongo.yml down

### Development
#.PHONY: up
#up: #mongo-up
#	@echo ">>> Ensure compose up"
#	docker-compose up --build -d

.PHONY: run
run:
	@echo ">>> Starting API server"
	go run cmd/app/main.go start
	#docker-compose exec $(MAIN) sh -c  './scripts/start.sh'

.PHONY: test
test:
	@echo Running go test
	go test -coverprofile=coverage.out ./...


.PHONY: cover
cover: test
	@echo Running go coverage
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
