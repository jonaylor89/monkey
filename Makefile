
APP=monkey
REGISTRY?=jonaylor
COMMIT_SHA=$(shell git rev-parse --short HEAD)

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o ${APP} .

.PHONY: clean
## clean: cleans the binary
clean:
	@rm -rf ${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: install 
## install: install the binary 
install:
	go install

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: docker-build
## docker-build: build docker image and tag with latest commit
docker-build:
	docker build -t ${APP} .
	docker tag ${APP} ${REGISTRY}/${APP}:${COMMIT_SHA}

.PHONY: docker-push
## docker-push: push latest docker image to docker hub
docker-push: docker-build
	@docker push ${REGISTRY}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
