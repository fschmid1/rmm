.PHONY: clean build run env bazel
clean:
	rm -rf ./build

build.backend: bazel
	bazel build //projects/backend:RMM_API

build.backend.linux: bazel
	bazel build //projects/backend:RMM_API_LINUX

build.client: bazel
	bazel build //projects/client:RMM_CLIENT

build.client.linux: bazel
	bazel build //projects/client:RMM_CLIENT_LINUX

build.client.arm: bazel
	bazel build //projects/client:RMM_CLIENT_ARM

run.backend: build.backend
	. ./projects/backend/.env && bazel run //projects/backend:RMM_API

run.frontend:
	cp projects/frontend/src/vars.dev.ts projects/frontend/src/vars.ts
	npx turbo dev

build.frontend:
	cp projects/frontend/src/vars.prod.ts projects/frontend/src/vars.ts
	npx turbo build

bazel:
	bazel run //:gazelle

run.client: build.client
	sudo bazel-bin/projects/client/RMM_CLIENT_/RMM_CLIENT
