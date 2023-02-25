FROM golang:1.20-alpine AS base
RUN apk add gcc musl-dev make

FROM base AS builder
WORKDIR /app
ADD . /app
ADD build/Makefile /app
RUN go mod download && go mod verify
RUN --mount=type=cache,target=/root/.cache/go-build \
    make test && make binaries LDFLAGS='-s -w -linkmode external -extldflags "-static"'

FROM telegraf:alpine
WORKDIR /plugins
COPY --from=builder /app/bin/telegraf-pihole /plugins/
