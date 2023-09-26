#!/usr/bin/bash

BASE_DIR=$(dirname "$0")

IMAGE_NAME="go-context-examples"

if [[ ! -z "${DOCKER_IMAGE_NAME}" ]]; then
  IMAGE_NAME="${DOCKER_IMAGE_NAME}"
fi

docker build -t "$IMAGE_NAME" "$BASE_DIR/../"
docker run --rm "$IMAGE_NAME"