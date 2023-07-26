SHELL:=/bin/sh
.PHONY: all vet test build clean docker docker-push docker-test

export GO111MODULE=on
export GOPROXY=https://goproxy.io

# path related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}/build/bin

pkgs	= $(shell go list ./... | grep -v vendor/)

DOCKER_IMAGE_NAME ?= feiyu563/prometheus-alert

BRANCH 		?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDDATE   ?= $(shell date -I'seconds')
BUILDUSER   ?= $(shell whoami)@$(shell hostname)
REVISION    ?= $(shell git rev-parse HEAD)
TAG_VERSION ?= $(shell git describe --tags --abbrev=0)

VERSION_LDFLAGS := \
	-X main.Version=$(TAG_VERSION) \
	-X main.Revision=$(REVISION) \
	-X main.BuildUser=$(BUILDUSER) \
	-X main.BuildDate=$(BUILDDATE)

# go source files, ignore vendor directory
SOURCE = $(shell find ${MKFILE_DIR} -type f -name "*.go")
TARGET = ${RELEASE_DIR}/PrometheusAlert

all: ${TARGET}

${TARGET}: ${SOURCE}
	@echo ">> building code"
	go mod tidy
	go mod vendor
	mkdir -p ${RELEASE_DIR}
	go build -ldflags "$(VERSION_LDFLAGS)" -o ${TARGET}

format:
	@echo ">> formatting code"
	go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	go vet $(pkgs)

test:
	@echo ">> running short tests"
	go test -short $(pkgs)

build: all

clean:
	@echo ">> cleaning build"
	rm -rf ${MKFILE_DIR}build

docker:
	@echo ">> building docker image"
	docker build -t "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" .
	docker tag "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" "$(DOCKER_IMAGE_NAME):latest"

docker-push:
	@echo ">> pushing docker image"
	docker push "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)"
	docker push "$(DOCKER_IMAGE_NAME):latest"

docker-test:
	@echo ">> testing docker image and PrometheusAlert's health"
	cmd/test_image.sh "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" 8080
