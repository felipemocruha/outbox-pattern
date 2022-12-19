POSTGRES_URI="postgres://outbox:outbox@localhost:5432/outbox?sslmode=disable"

migrate: ## Apply migrations to database
	@migrate -database $(POSTGRES_URI) -path migrations up

undo_migration: ## Undo previous migration to database
	@migrate -database $(POSTGRES_URI) -path migrations down 1

force_migration: ## Force a specific version of migrations
# Usage: V=<version> make force_migration
	@migrate -database $(POSTGRES_URI) -path migrations force $(V)

reset_migrations: ## Reset all migrations
	@migrate -database $(POSTGRES_URI) -path migrations down

compile: ## Compile go binary
	@CGO_ENABLED=0 go build -o outbox-pattern -ldflags="-s -w"

docker:
	@docker buildx build -t felipemocruha/outbox-pattern:$(VERSION) -f Dockerfile .

help:  ## This menu
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: test build
.DEFAULT: help
