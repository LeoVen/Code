#! /usr/bin/bash

set -e

source ./scripts/log.sh

if [ "$1" == "i" ]; then
    log_warn "Ignoring errors"
    ignore=true
else
    ignore=false
fi

if [ "$1" == "up" ]; then

    if [ "$2" == "local" ]; then
        log_info "Local env chosen"
        dockerfile="./docker-compose.yaml"
    elif [ "$2" == "debug" ]; then
        log_info "Debug env chosen"
        dockerfile="./docker/docker-compose-debug.yaml"
    else
        log_err "Invalid first parameter (must be either local or debug): $1"
        exit 1
    fi

    log_info "Building with docker-compose"

    docker-compose -f "$dockerfile" build

    log_info "Pulling images with docker-compose"

    docker-compose -f "$dockerfile" pull

    log_info "Running containers with docker-compose"

    docker-compose -f $dockerfile up -d

    docker-compose ps

elif [ "$1" == "down" ]; then

    if [ "$2" == "i" ]; then
        log_warn "Ignoring errors"
        ignore=true
    else
        ignore=false
    fi

    log_info "Stopping services with docker-compose"

    docker-compose down || $ignore

else
    log_err "Expected up or down, got: $1"
fi

log_ok "Script end"
