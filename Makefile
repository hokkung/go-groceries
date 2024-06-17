.DEFAULT_GOAL := build

.PHONY: api
api:
	go run ./cmd/api/main.go


.PHONY: lint
lint:
	@echo "start linting"
	golangci-lint run
	@echo "finished linting"

.PHONY: test
test: build
	@echo "start running test"
	go test -cover ./...
	@echo "finished running test"


.PHONY: build
build: lint
	@echo "start building application"
	go build -o bin/api/main ./cmd/api/main.go
	@echo "finished building application"
