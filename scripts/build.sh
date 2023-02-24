#!/bin/bash
set -e

export REPO='ghcr.io/wyvernzora'
export TAG='dev'
export PROJECT_ROOT="$(readlink -f .)"

source "$(dirname "$0")"/images.sh

build_images
docker image prune -f
