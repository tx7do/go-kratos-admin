# Not Support Windows

.PHONY: help wire gen ent build api openapi init all vendor dep test cover vet lint docker

ifeq ($(OS),Windows_NT)
    IS_WINDOWS := 1
endif

CURRENT_DIR	:= $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR	:= $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

SRCS_MK		:= $(foreach dir, app, $(wildcard $(dir)/*/*/Makefile))

# initialize develop environment
init: plugin cli

# install protoc plugin
plugin:
	# go
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/menta2k/protoc-gen-redact/v3@latest
	# Dart
	@flutter pub global activate protoc_plugin
	# Typescript
	@npm install -g ts-proto

# install cli tools
cli:
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	@go install github.com/google/gnostic@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/tx7do/kratos-cli/config-exporter/cmd/cfgexp@latest
	@go install github.com/tx7do/kratos-cli/sql-orm/cmd/sql2orm@latest
	@go install github.com/tx7do/kratos-cli/sql-proto/cmd/sql2proto@latest
	@go install github.com/tx7do/kratos-cli/sql-kratos/cmd/sql2kratos@latest

# use docker compose to run backend services and all its dependency services like redis, mysql, etc.
compose-up:
	@docker compose up -d --force-recreate

compose-down:
	@docker compose down

# download dependencies of module
dep:
	@go mod download

# create vendor
vendor:
	@go mod vendor

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	@go vet

# run lint
lint:
	@golangci-lint run

# generate wire code
wire:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make wire;\
    )

# generate code by go:generate
gen:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make gen;\
    )

# generate ent code
ent:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make ent;\
    )

# generate protobuf api go code
api:
	@cd api && \
	buf generate

# generate OpenAPI v3 docs.
openapi:
	@cd api && \
	buf generate --template buf.admin.openapi.gen.yaml

# generate typescript.
ts:
	@cd api && \
	buf generate --template buf.admin.typescript.gen.yaml

# build all service applications
build:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make build;\
    )

# build docker image
docker:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make docker;\
    )

# export configuration to etcd
export:
	@cfgexp \
		--type=etcd \
		--addr=localhost:2379 \
		--proj=kratos_admin

# generate & build all service applications
all:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make app;\
    )

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
