# Task runner for this project (the Rails equivalent is the Rakefile).
# Every target is a thin wrapper around go / npm so the daily commands stay short.
.DEFAULT_GOAL := help
SHELL := /bin/bash

AVAILABLE_TASKS = awk 'BEGIN {FS = ":.*## "}; /^[a-zA-Z_-]+:.*## / {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

help: ## Show the available tasks
	@grep -hE '^[a-zA-Z_-]+:.*## ' $(MAKEFILE_LIST) | sort | $(AVAILABLE_TASKS)

setup: ## Install Go + Node dependencies and build the dev assets
	go mod tidy
	npm install
	npm run build:dev

deps: ## Download and tidy Go modules (go.mod / go.sum)
	go mod tidy
	go mod download

update: ## Upgrade every Go dependency to its latest minor/patch release
	go get -u ./...
	go mod tidy

assets: ## Build JS + CSS for development (unminified, with sourcemaps)
	npm run build:dev

dev: ## Watch JS + CSS (Vite + Tailwind v4) while developing (alias: assets-watch)
	npm run dev

assets-watch: ## Watch JS + CSS (Vite + Tailwind v4) while developing
	npm run dev

assets-prod: ## Build minified JS + CSS for production (static/dist, embedded in the binary)
	npm run build:prod

build: ## Compile the Go binary into tmp/
	go build -o tmp/runner-build .

release: ## Build the production assets and the binary that embeds them
	npm run build:prod
	go build -o tmp/runner-build .

run: ## Run the application
	go run .

watch: ## Run the application with hot reload (fresh)
	fresh

migrate: ## Run the pending migrations (Rails: bin/rails db:migrate)
	go run ./db/migrate/cmd -direction=up

rollback: ## Roll back the last migration (Rails: bin/rails db:rollback)
	go run ./db/migrate/cmd -direction=down

migrate-status: ## Show the migration status (Rails: bin/rails db:migrate:status)
	go run ./db/migrate/cmd -direction=status

seed: ## Load the seed data (Rails: bin/rails db:seed)
	go run ./seeds

test: ## Run the test suite
	go test ./...

storybook: ## Run Storybook on port 6006
	npm run storybook

.PHONY: help setup deps update assets dev assets-watch assets-prod build release run watch migrate rollback migrate-status seed test storybook
