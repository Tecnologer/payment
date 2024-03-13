.PHONY: *

# Variables
AUTH_PORT=8080
BANK_PORT=8081
LOCAL_DB=gatepay
LOCAL_DB_USER=postgres
LOCAL_DB_PASSWORD=S3cret*_2024
LOCAL_DB_HOST=localhost
LOCAL_DB_PORT=5432
DB_CONTAINER=gatepay-db

run-auth:
	./dist/auth --port $(AUTH_PORT)

build-all: build-auth build-bank build-migrator

build-auth:
	cd ./auth && go build -o ../dist/auth main.go

run-bank:
	./dist/bank --port $(BANK_PORT)

build-bank:
	cd ./bank && go build -o ../dist/bank ./restapi/main.go

build-all: build-auth build-bank

build-n-run-migrator: build-migrator run-migrator

run-migrator:
	./dist/migrator --db-name $(LOCAL_DB) \
					--db-user $(LOCAL_DB_USER) \
					--db-pass $(LOCAL_DB_PASSWORD) \
					--db-host $(LOCAL_DB_HOST) \
					--db-port $(LOCAL_DB_PORT)

build-migrator:
	cd gatepay && go build -o ../dist/migrator ./migrator/main.go

docker-create-db:
	 docker run -d -p $(LOCAL_DB_PORT):$(LOCAL_DB_PORT) --rm -it --name $(DB_CONTAINER) \
        -e POSTGRES_PASSWORD=$(LOCAL_DB_PASSWORD) \
        -e POSTGRES_USER=$(LOCAL_DB_USER) \
        -e POSTGRES_DB=$(LOCAL_DB) postgres:latest

docker-run-all:
	docker-compose up --build