# Commands for running the API and MySQL database by using docker directly
.PHONY: setup restart restart_debug up up_debug down dc_restart dc_restart_debug dc_up dc_up_debug dc_down mysql protos

setup:
	chmod 755 -R ./scripts

restart: down up

restart_debug: down up_debug

up: protos
	@./scripts/docker.sh up local

up_debug: protos
	@./scripts/docker.sh up debug

down:
	@./scripts/docker.sh down i

dc_restart: dc_down dc_up

dc_restart_debug: dc_down dc_up_debug

dc_up:
	@./scripts/dcompose.sh up local

dc_up_debug:
	@./scripts/dcompose.sh up debug

dc_down:
	@./scripts/dcompose.sh down i

mysql:
	mysql -h 127.0.0.1 -P 3310 -u celluser -p CELLDB

protos:
	@./scripts/protos.sh
