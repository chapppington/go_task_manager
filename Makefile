DC = docker compose
STORAGES_FILE = docker_compose/storages.yaml
STORAGES_CONTAINER = postgres
LOGS = docker logs
ENV = --env-file .env
EXEC = docker exec -it
APP_FILE = docker_compose/app.yaml
APP_CONTAINER = main-app
# Add Go bin directory to PATH (for go imports command)
PATH := $(HOME)/go/bin:$(PATH)
export PATH

.PHONY: all
all:
	${DC} -f ${STORAGES_FILE} -f ${APP_FILE} ${ENV} up --build -d

.PHONY: all-down
all-down:
	${DC} -f ${STORAGES_FILE} -f ${APP_FILE} ${ENV} down

.PHONY: app-logs
app-logs:
	${LOGS} ${APP_CONTAINER} -f

.PHONY: app-shell
app-shell:
	${EXEC} ${APP_CONTAINER} sh

.PHONY: app-down
app-down:
	${DC} -f ${APP_FILE} ${ENV} down

.PHONY: storages
storages:
	${DC} -f ${STORAGES_FILE} ${ENV} up --build -d

.PHONY: storages-down
storages-down:
	${DC} -f ${STORAGES_FILE} ${ENV} down

.PHONY: storages-logs
storages-logs:
	${LOGS} ${STORAGES_CONTAINER} -f

.PHONY: postgres
postgres:
	${EXEC} ${STORAGES_CONTAINER} psql -U postgres

.PHONY: build
build:
	go build -o bin/main ./cmd/main.go

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: test
test:
	go test ./tests/...

.PHONY: precommit
precommit:
	pre-commit run --all-files
