#!/bin/bash

cd "$(dirname "$(readlink -f "$0")")" || exit 1

sudo docker load -i "./talim-image-${BRANCH}.tar"

mkdir -p "${TALIM_DIR}/log"
mkdir -p "${TALIM_DIR}/tmp"

sudo docker stop "talim-${BRANCH}"
sudo docker rm "talim-${BRANCH}"
sudo docker run \
	-d --name "talim-${BRANCH}" \
	-p "127.0.0.1:${TALIM_PORT}:80" \
	--mount "type=bind,source=${TALIM_DIR},target=/static" \
	--restart always \
	"talim-image-${BRANCH}"
