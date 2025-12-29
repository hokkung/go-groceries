FROM golang:1.22-alpine3.19 AS builder

RUN apk update && \
    apk add --no-cache  \
    git

WORKDIR /app

COPY go.mod go.sum ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '' -a -installsuffix cgo -o /output/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '' -a -installsuffix cgo -o /output/internal-api ./cmd/internal-api

FROM alpine:latest

RUN apk update && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache tzdata

ENV TZ=Asia/Bangkok

WORKDIR  /app

COPY --from=builder /output/api /app/api
COPY --from=builder /output/internal-api /app/internal-api
COPY --from=builder /app/script/entrypoint.sh /script/entrypoint.sh

ENTRYPOINT ["/script/entrypoint.sh"]
CMD ["api"]
