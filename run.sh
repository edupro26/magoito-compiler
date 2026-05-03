#!/bin/bash
set -e

IMAGE_NAME="magoito"
docker build -t "$IMAGE_NAME" .
docker run -it --rm --name "magoito-compiler" "$IMAGE_NAME"