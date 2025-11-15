.PHONY: build clean docker gen ent wire api openapi run app help

GOPATH ?= $(shell go env GOPATH)
# GOVERSION is the current go version, e.g. go1.9.2
GOVERSION ?= $(shell go version | awk '{print $$3;}')

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT	:= awk '{ print } END { if (NR > 0) { exit 1 } }'

GO_CMD			:= GO111MODULE=on go
GIT_CMD			:= git
DOCKER_CMD		:= docker

ARCH			:= "`uname -s`"
LINUX			:= "Linux"
MAC				:= "Darwin"

DEFAULT_VERSION	:= 0.0.1

ifeq ($(OS),Windows_NT)
    IS_WINDOWS	:= TRUE
endif

ifneq (git,)
	GIT_EXIST	:= TRUE
endif

ifneq ("$(wildcard .git)", "")
	HAS_DOTGIT	:= TRUE
endif

ifeq ($(GIT_EXIST),TRUE)
ifeq ($(HAS_DOTGIT),TRUE)
	# CUR_TAG is the last git tag plus the delta from the current commit to the tag
	# e.g. v1.5.5-<nr of commits since>-g<current git sha>
	CUR_TAG ?= $(shell git describe --tags --first-parent)

	# LAST_TAG is the last git tag
    # e.g. v1.5.5
    LAST_TAG ?= $(shell git describe --match "v*" --abbrev=0 --tags --first-parent)

    # VERSION is the last git tag without the 'v'
    # e.g. 1.5.5
    VERSION ?= $(shell git describe --match "v*" --abbrev=0 --tags --first-parent | cut -c 2-)
endif
endif

CUR_TAG		?= $(DEFAULT_VERSION)
LAST_TAG	?= v$(DEFAULT_VERSION)
VERSION		?= $(DEFAULT_VERSION)

# GOFLAGS is the flags for the go compiler.
GOFLAGS		?= -ldflags "-X main.version=$(VERSION)"

PROJECT_NAME		:= kratos-admin
APP_RELATIVE_PATH	:= $(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
SERVICE_NAME		:= $(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b)
APP_NAME			:= $(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")

# build golang application
build: api openapi
	@go build -ldflags "-X main.version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# build docker image
docker:
	@docker build -t $(PROJECT_NAME)/$(APP_NAME) \
				  --build-arg SERVICE_NAME=$(SERVICE_NAME) \
				  --build-arg APP_VERSION=$(APP_VERSION) \
				  -f ../../../Dockerfile ../../../

# generate code
gen: ent wire api openapi

# generate ent code
ent:
ifneq ("$(wildcard ./internal/data/ent)","")
	@ent generate \
				--feature privacy \
				--feature entql \
				--feature sql/modifier \
				--feature sql/upsert \
				--feature sql/lock \
				./internal/data/ent/schema
endif

# generate wire code
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

# generate protobuf api go code
api:
	@cd ../../../api && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../../api && \
	buf generate --template buf.admin.openapi.gen.yaml

# run application
run: api openapi
	-@go run ./cmd/server -conf ./configs

# build service app
app: api openapi wire ent build

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
