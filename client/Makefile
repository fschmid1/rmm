.PHONY: clean build run

APP_NAME = RMM_CLIENT
BUILD_DIR = $(PWD)/build

GOFLAGS = -ldflags="-w -s"

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	sudo $(BUILD_DIR)/$(APP_NAME)
