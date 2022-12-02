DOCKER_COMPOSE_FOLDER := ./deployments
DOCKER_FILE := $(DOCKER_COMPOSE_FOLDER)/docker-compose.yaml
DOCKER_TEST_FILE := $(DOCKER_COMPOSE_FOLDER)/docker-compose.test.yaml
INTEGRATION_TEST_PREFIX := docker-compose -f $(DOCKER_FILE) -f $(DOCKER_TEST_FILE)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'
.PHONY: help

run: ## docker up
	cd $(DOCKER_COMPOSE_FOLDER) && docker-compose up -d
.PHONY: run

build: ## build service anti bruteforce
	go build -o ./bin/antiBruteForceService ./cmd/server
.PHONY: build

docker.rebuild: ## docker rebuild
	cd $(DOCKER_COMPOSE_FOLDER) && docker-compose up -d --build
.PHONY: docker.rebuild

down: ## docker down
	cd $(DOCKER_COMPOSE_FOLDER) && docker-compose down --remove-orphans

test: ## run test
	go test -race -count 100 ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps ## run linter
	golangci-lint run ./...

pre-commit: lint ## Pre commit handles
	go mod tidy
	go mod verify
	gofmt -w -s .
.PHONY: pre-commit

generate: ## Run go generate
	protoc -I api --go_out=proto/api --go_opt=paths=source_relative \
        --go-grpc_out=proto/api --go-grpc_opt=paths=source_relative \
        api/*.proto
	go generate ./...

test-packages: ## Test all packages
	go test all

wire: ## Run DI generating
	wire cmd/server/wire.go

integration-tests: ## Run integration tests
	$(INTEGRATION_TEST_PREFIX) up --build integration_test
	$(INTEGRATION_TEST_PREFIX) down --volumes --remove-orphans

full: wire pre-commit generate test
