#! /usr/bin/bash

set -e

source ./scripts/log.sh

if ! docker info > /dev/null 2>&1; then
    log_err "Docker is not running"
fi

if [ "$1" == "local" ]; then
    log_info "Local env chosen"
    dockerfile="./docker/Dockerfile.api"
elif [ "$1" == "debug" ]; then
    log_info "Debug env chosen"
    dockerfile="./docker/Dockerfile.api.debug"
else
    log_err "Invalid first parameter (must be either local or debug): $1"
fi

log_info "Building images"

docker build -t cellapi:1.0 -f $dockerfile .
docker build -t celldb:1.0 -f ./docker/Dockerfile.mysql .

log_info "Running containers"

if [ "$( docker container inspect -f '{{.State.Running}}' cellapi_instance)" == "true" ]; then
    log_err "Container cellapi_instance already exists"
    exit 1
fi

docker run -p 8880:8080 -p 2345:2345 -tid --name cellapi_instance cellapi:1.0

if [ "$( docker container inspect -f '{{.State.Running}}' celldb_instance)" == "true" ]; then
    log_err "Container celldb_instance already exists"
    exit 1
fi

docker run -p 3310:3306 -tid --name celldb_instance celldb:1.0

log_info "Creating network"

docker network create cell_net
docker network connect --alias cellapi cell_net cellapi_instance
docker network connect --alias celldb cell_net celldb_instance

log_ok "Done"
