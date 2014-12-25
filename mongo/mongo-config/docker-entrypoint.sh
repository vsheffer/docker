#!/bin/bash
set -e

which mongod
echo "Hello: $@"

if [ "${1:0:1}" = '-' ]; then
	set -- mongod "$@"
fi

if [ "$1" = 'mongod' ]; then
	chown -R mongodb /data/config
	exec gosu mongodb "$@"
fi

echo "Hello2: $@"
exec $@
