SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng.js"

$(shell cp --update=none .config.default config.ini)
-include config.ini

local:
	ulimit -n 51200; \
	$(NG) serve --watch --poll 2000 --no-hmr --port $(port) --allowed-hosts $(domain) --host 127.0.0.1

prod:
	./dist/build.sh talim.9farm.com prod
	cp ../public/favicon.ico dist/prod/browser
	cp ../public/robots.txt dist/prod/browser
	cp ../public/twitter-copier.user.js dist/prod/browser
	ssh doll 'mkdir -p /www/talim/prod'
	rsync --partial -vzrtopg -e ssh ./dist/prod/browser/ doll:/www/talim/prod
