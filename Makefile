.DEFAULT_GOAL := run

MAIN := app

.PHONY: build
build:
	@echo Building binaries
	sh ./scripts/build.sh

.PHONY: up
up:
	@echo ">>> Ensure compose up"
	docker-compose up --build -d

.PHONY: run
run: up
	@echo ">>> run"
	docker-compose exec $(MAIN) sh -c  './scripts/start.sh'