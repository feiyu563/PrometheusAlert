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

all: format vet test build

.PHONY: format
format:
	@echo ">> formatting code"
	go fmt $(pkgs)

.PHONY: vet
vet:
	@echo ">> vetting code"
	go vet $(pkgs)

.PHONY: test
test:
	@echo ">> running short tests"
	go test -short $(pkgs)

.PHONY: build
build:
	@echo ">> building code"
	go mod tidy
	go mod vendor
	GO11MODULE=on GO111MODULE=on GOPROXY=https://goproxy.io \
	  go build -ldflags "$(VERSION_LDFLAGS)" -o PrometheusAlert

.PHONY: docker
docker:
	@echo ">> building docker image"
	docker build -t "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" .
	docker tag "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" "$(DOCKER_IMAGE_NAME):latest"

.PHONY: docker-push
docker-push:
	@echo ">> pushing docker image"
	docker push "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)"
	docker push "$(DOCKER_IMAGE_NAME):latest"

.PHONY: docker-test
docker-test:
	@echo ">> testing docker image and PrometheusAlert's health"
	cmd/test_image.sh "$(DOCKER_IMAGE_NAME):$(TAG_VERSION)" 8080
