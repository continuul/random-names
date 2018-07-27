#!/usr/bin/dumb-init /bin/sh
set -e

CONFIG_DIR=/random-names/config

if [ "${1:0:1}" = '-' ]; then
    set -- random-names "$@"
fi

if [ "$1" = 'server' ]; then
    shift
    set -- random-names server \
    --config-dir="$CONFIG_DIR" \
    "$@"
elif [ "$1" = 'version' ]; then
    set -- random-names "$@"
elif random-names "$1" --help 2>&1 | grep -q "random-names $1"; then
    set -- random-names "$@"
fi

# don't run as root...
if [ "$1" = 'random-names' ]; then
    set -- su-exec fun:fun "$@"
fi

exec "$@"
