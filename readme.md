# @go-groceries

An example of microservice application written by Golang.
This project is implemented by applying clean code architecture. 

## Table of Contents
- [Prerequisites](#prerequisites)
- [Getting started](#getting-started)
- [Environment](#environment)
- [Testing](#testing)
- [Internal Library](#internal-library)

---
## Prerequisites
1. install Go 
2. Install Wire [Official-Site](https://github.com/google/wire)
3. Install Mockgen [Official-Site](https://github.com/golang/mock)

---
## Getting Started
1. Start MYSQL
   - `docker compose -f /docker-compose up`
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
1. Manage HTTP Server (https://github.com/hokkung/srv)
2. Manage Repository (https://github.com/hokkung/gorem)
3. Manage Redis (https://github.com/hokkung/redis)
