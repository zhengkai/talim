SHELL:=/bin/bash

TARGET_HOST:=freya

build: git
	sudo docker build -t talim -f Dockerfile ../..

run: build
	sudo docker run --env \
		--mount type=bind,source=/www/talim/static,target=/static \
		--mount type=bind,source=/www/talim/server/dist/prod/log,target=/log \
		talim

install: build
	sudo docker save talim > docker-talim.tar
	scp docker-talim.tar $(TARGET_HOST):/tmp
	scp install.sh $(TARGET_HOST):/tmp
	ssh $(TARGET_HOST) "chmod +x /tmp/install.sh && /tmp/install.sh && rm /tmp/install.sh"

git:
	../../server/build/git-hash.sh > ../../server/build/.git-hash

clean:
	sudo docker stop talim || :
	sudo docker rm talim || :
	sudo docker rmi talim || :
	sudo docker image prune --force
