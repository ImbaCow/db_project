#!/bin/bash

set -e

host="$1"
cmd="$@"

>&2 echo "Check host for available"

until curl http://"$host"; do
  >&2 echo "Host is unavailable - sleeping"
  sleep 1
done

>&2 echo "Host is up - executing command"

if [ $cmd != "" ]
then
  echo "Command: '${$cmd}'"
  exec $cmd
fi