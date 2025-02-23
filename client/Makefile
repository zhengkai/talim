SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng.js"

$(shell cp --update=none .config.default config.ini)
-include config.ini

local:
	ulimit -n 51200; \
	$(NG) serve --watch --poll 2000 --no-hmr --port $(port) --allowed-hosts $(domain) --host 127.0.0.1

