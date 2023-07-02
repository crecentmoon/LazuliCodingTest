DB_USER=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=3306
DB_NAME=lazuli

ps:
	docker-compose ps

build:
	docker-compose build --no-cache

up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose stop; docker-compose down

.PHONY: app
app:
	docker-compose exec app bash

.PHONY: db
db:
	docker-compose exec db bash

.PHONY: install-migrate
install-migrate:
	docker-compose exec app bash -c "go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest"

# .PHONY: migration-file
# migration-file:
# 	docker-compose exec app bash -c "migrate create -ext sql -dir migrations -seq hogehoge"

.PHONY: migrate-up
migrate-up:
	docker-compose exec app bash -c "migrate -path migrations -database 'mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true' -verbose up"