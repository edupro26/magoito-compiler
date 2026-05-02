#!/bin/bash
set -e

IMAGE_NAME="magoito"
docker build -t "$IMAGE_NAME" .
docker run -it --name "magoito-compiler" "$IMAGE_NAME"