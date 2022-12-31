.PHONY: clean build run

APP_NAME = RMM_CLIENT
BUILD_DIR = $(PWD)/build

GOFLAGS = -ldflags="-w -s"

clean:
	rm -rf ./build

build:
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) main.go

arm: clean
	env GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFLAGS) main.go
make linux: clean
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFLAGS) main.go

run: build
	sudo $(BUILD_DIR)/$(APP_NAME)
