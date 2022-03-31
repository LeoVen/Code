#! /usr/bin/bash

set -e

source ./scripts/log.sh

if [ "$1" == "i" ]; then
    log_warn "Ignoring errors"
    ignore=true
else
    ignore=false
fi

log_info "Stopping docker containers"

docker stop cellapi_instance || $ignore
docker stop celldb_instance || $ignore

log_info "Removing docker containers"

docker container rm cellapi_instance || $ignore
docker container rm celldb_instance || $ignore

log_info "Removing docker images"

docker image rm cellapi:1.0 || $ignore
docker image rm celldb:1.0 || $ignore

log_info "Removing network"

docker network rm cell_net || $ignore

log_ok "Done"
