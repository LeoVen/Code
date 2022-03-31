# Commands for running the API and MySQL database by using docker directly
.PHONY: restart up up_debug down mysql protos

restart: down up

up:
	docker build -t cellapi:1.0 -f ./docker/Dockerfile.api .
	docker build -t celldb:1.0 -f ./docker/Dockerfile.mysql .

	docker run -p 8880:8080 -p 2345:2345 -tid --name cellapi_instance cellapi:1.0
	docker run -p 3310:3306 -tid --name celldb_instance celldb:1.0

	docker network create cell_net
	docker network connect --alias cellapi cell_net cellapi_instance
	docker network connect --alias celldb cell_net celldb_instance

up_debug:
	docker build -t cellapi:1.0 -f ./docker/Dockerfile.api.debug .
	docker build -t celldb:1.0 -f ./docker/Dockerfile.mysql .

	docker run -p 8880:8080 -p 2345:2345 -tid --name cellapi_instance cellapi:1.0
	docker run -p 3310:3306 -tid --name celldb_instance celldb:1.0

	docker network create cell_net
	docker network connect --alias cellapi cell_net cellapi_instance
	docker network connect --alias celldb cell_net celldb_instance

down:
	docker stop cellapi_instance
	docker stop celldb_instance

	docker container rm cellapi_instance
	docker container rm celldb_instance

	docker image rm cellapi:1.0
	docker image rm celldb:1.0

	docker network rm cell_net

mysql:
	mysql -h 127.0.0.1 -P 3310 -u celluser -p CELLDB

protos:
	@./scripts/protos.sh
