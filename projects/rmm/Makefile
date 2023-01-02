.PHONY: clean build run env
clean:
	rm -rf ./build

build.backend:
	bazel build //projects/rmm/go/backend:RMM_API

build.backend.linux:
	bazel build //projects/rmm/go/backend:RMM_API_LINUX

build.client:
	bazel build //projects/rmm/go/client:RMM_CLIENT

build.client.linux:
	bazel build //projects/rmm/go/client:RMM_CLIENT_LINUX

build.client.arm:
	bazel build //projects/rmm/go/client:RMM_CLIENT_ARM

linux: clean
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFLAGS) main.go

run.backend: build.backend
	. ./go/backend/.env && bazel run //projects/rmm/go/backend:RMM_API

run.client: build.client
	bazel run //projects/rmm/go/client:RMM_CLIENT
