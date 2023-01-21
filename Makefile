.PHONY: clean build run env bazel
clean:
	rm -rf ./build

build.backend:
	bazel build //projects/backend:RMM_API

build.backend.linux:
	bazel build //projects/backend:RMM_API_LINUX

build.client:
	bazel build //projects/client:RMM_CLIENT

build.client.linux:
	bazel build //projects/client:RMM_CLIENT_LINUX

build.client.arm:
	bazel build //projects/client:RMM_CLIENT_ARM

linux: clean
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) $(GOFLAGS) main.go

run.backend: build.backend
	. ./projects/backend/.env && bazel run //projects/backend:RMM_API

bazel:
	bazel run //:gazelle

run.client: build.client
	sudo bazel-bin/projects/client/RMM_CLIENT_/RMM_CLIENT
