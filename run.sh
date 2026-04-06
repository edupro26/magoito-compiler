#!/bin/bash
set -e

IMAGE_NAME="magoito"
docker build -t "$IMAGE_NAME" .
docker run --rm -it "$IMAGE_NAME"