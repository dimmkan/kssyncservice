FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

ENV GOPROXY https://repository.rt.ru/repository/go-proxy.golang.org

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

WORKDIR /app

COPY --from=builder /app/main /app/main

COPY ./.env .

CMD ["./main"]