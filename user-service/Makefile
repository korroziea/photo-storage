POSTGRES_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

run: ## run application `make run`
	go run cmd/server/main.go

migrate-create: ## create migration file `make migrate-create name=migration_name`
	@migrate create -dir migrations -seq -ext sql $(name)

migrate-up: ## up migration files `make migrate-up`
	@migrate -path=migrations -database ${POSTGRES_URL} -verbose up

migrate-down: ## down one migration file `make migrate-down`
	@migrate -path=migrations -database ${POSTGRES_URL} -verbose down 1

migrate-version: ## get current migration version `make migrate-version`
	@migrate -path=migrations -database ${POSTGRES_URL} version
