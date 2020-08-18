.PHONY: all build run test lint up down db-up db-in clean mig-cr mig-up mig-down mig-v

all: test bin/server

run: clean
	DOCKER_BUILDKIT=1 docker-compose up --build server

bin/server:
	DOCKER_BUILDKIT=1 docker build --target bin-cloud . \
	--output bin/

# TODO dockerイメージをビルドするように書き換える
build:
	go build

test:
	go test ./... -v -coverpkg=./...

lint:
	golangci-lint run

up: clean db-up
	docker-compose up --build gorm-test

down:
	docker-compose down --rmi all --volumes

clean:
	docker system prune --volumes -f

db-up:
	DOCKER_BUILDKIT=1 docker-compose up --build -d db

db-in:
	docker-compose exec db bash

mig-cr:
	migrate create -ext sql -dir db/migrations -seq ${name}

mig-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

mig-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

mig-v:
	migrate -database ${POSTGRESQL_URL} -path db/migrations version
