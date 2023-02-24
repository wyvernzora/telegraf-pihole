#!/bin/bash
set -e

build_images() {
  for stage in build/*; do
    build_image "$stage"
  done
}

build_image() {
  IMAGE_ROOT=$1
  NAME=${1##build/[0-9][0-9]-}
  FULL_NAME=$REPO/telegraf-pihole-$NAME:$TAG

  echo building "$FULL_NAME"
  echo image root "$IMAGE_ROOT"

  docker build \
    --build-arg BASE_NAME="ghcr.io/wyvernzora/telegraf-pihole" \
    --build-arg TAG="$TAG" \
    --build-arg IMAGE_ROOT="$IMAGE_ROOT" \
    -f "$IMAGE_ROOT/Dockerfile" \
    -t "$FULL_NAME" \
    .
}
