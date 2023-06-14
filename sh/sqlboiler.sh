#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [ "$#" -ne 1 ]; then
	echo "usage: $0 <database-name>"
	exit 1
fi

DATABASE_NAME=$1
OUTPUT_PATH="pkg/$DATABASE_NAME/adapter/gateway/db/model"

cat <<EOF > sqlboiler.toml 
pkgname="model"
output="$OUTPUT_PATH"

[mysql]
  dbname = "$DATABASE_NAME"
  host   = "127.0.0.1"
  port   = 3306
  user   = "user"
  pass   = "password"
  sslmode = "false"
  blacklist = ["schema_migrations"]
EOF

sqlboiler mysql --add-enum-types --no-tests --wipe
rm sqlboiler.toml
