APP_NAME = server
SQLC_QUERY? = sqlc generate
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:Vu123456@tcp(127.0.0.1:3300)/go_service"

swag:
	swag init -g cmd/server/main.go -o ./cmd/swag/docs
dev:
	go run cmd/${APP_NAME}/main.go
debug:
    go run -tags=jsoniter cmd/${APP_NAME}/main.go
docker_build:
	docker-compose up -d --build
docker_show:
	docker-compose ps
docker_stop:
	docker-compose down
docker_up:
	docker-compose up -d
create_migration:
	@Goose -dir=${GOOSE_MIGRATION_DIR} create $(name) sql
upMysql:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downMysql:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetMysql:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen: $(SQLC_QUERY)

.PHONY: dev upMysql downMysql resetMysql sqlgen docker_build docker_stop docker_up create_migration debug docker_show swag