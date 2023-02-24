ARG BASE_NAME
ARG TAG=dev
ARG IMAGE_ROOT
FROM ${BASE_NAME}-builder:$TAG AS builder

FROM telegraf:alpine
WORKDIR /plugins
COPY --from=builder /app/bin/telegraf-pihole /plugins/
