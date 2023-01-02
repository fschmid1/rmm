.PHONY: clean build run env
clean:
	rm -rf ./build

build.backend:
	bazel build //projects/rmm/go/backend/...

build.client:
	bazel build //projects/rmm/go/client/...

linux: clean
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFLAGS) main.go

run.backend: build.backend
	. ./go/backend/.env && bazel run //projects/rmm/go/backend:RMM_API

run.client: build.client
	bazel run //projects/rmm/go/client:RMM_CLIENT

docker.build:
	docker-compose build

docker.push:
	buildx docker-compose.yml -p -a