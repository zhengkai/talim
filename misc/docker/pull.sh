#!/bin/bash

cd /www/publish/talim/misc/docker || exit 1
git pull
make build

cd "$(dirname "$(readlink -f "$0")")" || exit 1

#sudo docker load -i "./talim-image-master.tar"

mkdir -p "/www/talim/static/log"
mkdir -p "/www/talim/static/tmp"

sudo docker stop "talim-master"
sudo docker rm "talim-master"
sudo docker run \
	-d --name "talim-master" \
	-p "127.0.0.1:51828:80" \
	--mount "type=bind,source=/www/talim/static,target=/static" \
	--restart always \
	"talim"
