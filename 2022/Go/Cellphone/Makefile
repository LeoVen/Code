# Commands for running the API and MySQL database by using docker directly
.PHONY: restart restart_debug up up_debug down mysql protos

restart: down up

restart_debug: down up_debug

up:
	@./scripts/docker.up.sh local

up_debug:
	@./scripts/docker.up.sh debug

down:
	@./scripts/docker.down.sh i

mysql:
	mysql -h 127.0.0.1 -P 3310 -u celluser -p CELLDB

protos:
	@./scripts/protos.sh
