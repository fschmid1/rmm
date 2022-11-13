.PHONY: clean build run env

APP_NAME = RMM_API
BUILD_DIR = $(PWD)/build
DATABASE_URL = postgres://postgres:password@cgapp-postgres/postgres?sslmode=disable

SHELL := /bin/bash

GOFLAGS = -ldflags="-w -s"

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	. ./.env && $(BUILD_DIR)/$(APP_NAME)

docker.run: docker.build
	docker-compose up --force-recreate

docker.build:
	docker-compose build
