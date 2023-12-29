PROJECT_NAME=coffeeshop
MODULE_NAME=coffeeshop

.DEFAULT_GOAL := build

.PHONY: build
build:
	@go build . && sudo cp ./auto_coffeeshop /usr/local/bin/runcoffee

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out

.PHONY: get
get:
	@go mod download

.PHONY: docker
docker:
	@docker build -t coffeeshop:latest .

