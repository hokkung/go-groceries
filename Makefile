.DEFAULT_GOAL := build

.PHONY: api
api:
	go run ./cmd/api/main.go

.PHONY: gen
gen:
	@echo "start generating code"
	go generate ./...
	@echo "start generating code"

.PHONY: lint
lint:
	@echo "start linting"
	golangci-lint run
	@echo "finished linting"

.PHONY: test
test: build
	@echo "start running test"
	GOARCH=amd64 go test -cover ./...
	@echo "finished running test"

.PHONY: test-cov
test-cov:
	@echo "start running test coverage 😱😱😱"
	GOARCH=amd64 go test ./... -coverprofile=coverage.out.tmp
	cat coverage.out.tmp | grep -v "mock" > coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
	@echo "done test! good job bro! 🎉🎉🎉"

.PHONY: build
build: gen lint
	@echo "start building application"
	go build -o bin/api/main ./cmd/api/main.go
	@echo "finished building application"

.PHONY: wire
wire:
	@echo "start wiring"
	wire ./...
	@echo "finished generating wire"
