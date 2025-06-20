#!/bin/bash
set -euo pipefail

# Name of the image to build. Defaults to gitli-local
IMAGE_NAME=${1:-gitli-local}

# Build the Docker image from the deven Dockerfile

docker build -t "$IMAGE_NAME" deven

echo "Image built: $IMAGE_NAME"

printf '\nYou can push this image using:\n'
echo "  docker push $IMAGE_NAME"

printf '\nLaunching the CLI...\n'
# Run the CLI inside the container. Pass any additional arguments to the CLI

# Detect if we are running in an interactive terminal
if [ -t 1 ]; then
  DOCKER_FLAGS="-it"
else
  DOCKER_FLAGS="-i"
fi

docker run --rm $DOCKER_FLAGS "$IMAGE_NAME" ultimate-setup "${@:2}"
