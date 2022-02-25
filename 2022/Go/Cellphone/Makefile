
up: down
	docker build -t cellapi:1.0 -f ./Dockerfile.api .
	docker build -t celldb:1.0 -f ./Dockerfile.mysql .

	docker run -p 8080:8080 -tid --name cellapi_instance cellapi:1.0
	docker run -tid --name celldb_instance celldb:1.0

	docker network create cell_net
	docker network connect --alias cellapi cell_net cellapi_instance
	docker network connect --alias celldb cell_net celldb_instance

	echo Success

down:
	docker stop cellapi_instance
	docker stop celldb_instance

	docker rm cellapi_instance
	docker rm celldb_instance

	docker image rm cellapi:1.0
	docker image rm celldb:1.0

	docker network rm cell_net

	echo Success
