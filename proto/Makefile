SHELL:=/bin/bash

PBJS=../client/node_modules/protobufjs-cli/bin/pbjs
PBTS=../client/node_modules/protobufjs-cli/bin/pbts

server:
	protoc --go_out=../server/src *.proto

client:
	$(PBJS) -t static-module -w es6 --es6 --force-long -o ../client/src/pb/pb.js *.proto
	$(PBTS) -o ../client/src/pb/pb.d.ts ../client/src/pb/pb.js

all: server client
