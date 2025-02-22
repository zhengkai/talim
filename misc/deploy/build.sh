#!/bin/bash

BRANCH="${1:-master}"

cd "$(dirname "$(readlink -f "$0")")" || exit 1

IMAGE="talim-image-${BRANCH}"

sudo docker build \
	-t "$IMAGE" \
	-f ../docker/Dockerfile \
	..

FILE="output/${IMAGE}.tar"

sudo docker save "$IMAGE" -o "$FILE" >/dev/null 2>&1
sudo chown "${USER}:${USER}" "$FILE"
