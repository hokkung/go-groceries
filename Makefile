.PHONY: api
api:
	go run ./cmd/api/main.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: lint
	go build -o bin/api/main ./cmd/api/main.go
