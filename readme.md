# @go-groceries

An example of microservice application written by Golang.
This project is implemented by applying clean code architecture.

![GitHub Workflow Status](https://github.com/hokkung/go-groceries/actions/workflows/build.yml/badge.svg)

```
├── cmd
│   └── api
│       └── main.go
├── config
├── go.mod
├── go.sum
├── internal
│   ├── client
│   ├── di
│   ├── entity
│   ├── handler
│   ├── repository
│   ├── server
│   └── service
├── pkg
└── utils
```

## Table of Contents
- [Prerequisites](#prerequisites)
- [Getting started](#getting-started)
- [Environment](#environment)
- [Testing](#testing)
- [Internal Library](#internal-library)
- [Reference](#reference)

---
## Prerequisites
1. install Go 
2. Install Wire [Official-Site](https://github.com/google/wire)
3. Install Mockgen [Official-Site](https://github.com/golang/mock)

---
## Getting Started
1. Start Admin, MySQL container
   - `docker compose -f /docker-compose/{file_name} up`
2. Create a `grocery` database 
3. Run `make api`

---
## Environment
```
export APP_SERVER_ADDR=:8081
```

---
## Testing
1. Simply run command `make test`

---
## Internal Library
1. HTTP Server (https://github.com/hokkung/srv)
2. Gorm Repository (https://github.com/hokkung/gorem)
3. Redis (https://github.com/hokkung/redis)

---
## Reference
- [Cat APIs](https://developers.thecatapi.com/view-account/ylX4blBYT9FaoVd6OhvR?report=bOoHBz-8t) 
