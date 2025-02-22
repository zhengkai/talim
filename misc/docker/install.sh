#!/bin/bash

TARGET="Freya"

if [ "$HOSTNAME" != "$TARGET" ]; then
	>&2 echo only run in server "$TARGET"
	exit 1
fi

sudo docker stop talim
sudo docker rm talim
sudo docker rmi talim

sudo cat /tmp/docker-talim.tar | sudo docker load

sudo docker run -d --name talim \
	--mount type=bind,source=/www/talim/log,target=/log \
	--mount type=bind,source=/www/talim/static,target=/static \
	--restart always \
	talim
