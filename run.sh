#!/bin/bash
set -e

IMAGE_NAME="magoito"
docker build -q -t "$IMAGE_NAME" .
docker run --rm -it "$IMAGE_NAME"