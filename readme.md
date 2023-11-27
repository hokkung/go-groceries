
### Environment
```
export APP_SERVER_ADDR=:8081
```

### How to start service
1. Start MYSQL
- `docker compose -f /docker-compose up`
1. create a `grocery` database 
2. go run ./cmd/api/main.go


### How to run test
1. Make sure you already create mockgen if dependency is required.
   - `go generate ./...` to create mockgen file
2. Run `go test ./...`

### Thrid Party
1. Manage HTTP Server (https://github.com/hokkung/srv)
2. Manage Repository (https://github.com/hokkung/gorem)
3. Manage Redis (https://github.com/hokkung/redis)