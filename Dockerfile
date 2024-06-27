FROM golang:1.22-alpine3.19 AS builder

RUN apk update && \
    apk add --no-cache  \
    git

WORKDIR /app

COPY go.mod go.sum ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '' -a -installsuffix cgo -o go-groceries-api ./cmd/api

FROM alpine:latest

RUN apk update && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache tzdata

ENV TZ=Asia/Bangkok

WORKDIR  /root/

COPY --from=builder /app/go-groceries-api /root/api/main

EXPOSE 8081

ENTRYPOINT ["./api/main"]
