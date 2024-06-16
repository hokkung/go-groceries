build:
	go build -o bin/api/main ./cmd/api/main.go

api:
	go run ./cmd/api/main.go

test:
	go test -cover ./...