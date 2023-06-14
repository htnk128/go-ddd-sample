#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [ "$#" -ne 2 ]; then
  echo "usage: $0 <database-name> <command>"
  exit 1
fi

DATABASE_NAME=$1
COMMAND=$2

migrate -source file://./db/$DATABASE_NAME/migrations -database mysql://user:password@tcp\(127.0.0.1:3306\)/$DATABASE_NAME $COMMAND
