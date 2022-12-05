#! /usr/bin/bash

set -e

scriptName=`basename "$0"`

log()
{
    echo $'\e[1;90m' $(date "+%Y/%m/%d %H:%M:%S") $scriptName $'\e[0m' $1
}

log_ok()
{
    echo [$'\e[1;92m' OK \ \ $'\e[0m]' $(log "$@")
}

log_err()
{
    echo [$'\e[1;91m' ERR \ $'\e[0m]' $(log "$@")
}

log_info()
{
    echo [$'\e[1;94m' INFO $'\e[0m]' $(log "$@")
}

log_warn()
{
    echo [$'\e[1;93m' WARN $'\e[0m]' $(log "$@")
}

log_info "Starting installation"
log_err "Could not install"
log_warn "Not found in path. This could cause problems."
log_ok "Installation successfull"
