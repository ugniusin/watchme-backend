.PHONY: default
default: help

.PHONY: help
help:
	@echo Tasks:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

FORCE:

.PHONY:
build: ## Build App
	@docker-compose up -d

.PHONY:
migrate-up: ## Migrate Up
	@docker exec -it backend sh -c "sql-migrate up"

.PHONY:
migrate-down: ## Migrate Down
	@docker exec -it backend sh -c "sql-migrate down"
