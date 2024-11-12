include .env
export $(shell sed 's/=.*//' .env)

MIGRATE_CMD = migrate -path ./migrations -database $(DATABASE_URL)

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down
