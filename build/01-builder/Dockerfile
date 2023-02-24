ARG BASE_NAME
ARG TAG=dev
ARG IMAGE_ROOT
FROM ${BASE_NAME}-base:$TAG

WORKDIR /app
ADD . /app
ADD $IMAGE_ROOT/Makefile /app
RUN go mod download && go mod verify
RUN --mount=type=cache,target=/root/.cache/go-build \
    make binaries LDFLAGS='-s -w -linkmode external -extldflags "-static"'
