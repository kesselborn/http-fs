#!/bin/sh
BASIC_AUTH=${BASIC_AUTH:?please set \$BASIC_AUTH env var (something like "username:password")}
HOST=${HOST:?please set \$HOST env var}

IFS="
"; for f in $(find . -type f|grep -v "./$(basename $0)$"|sed 's/^\.//g')
do
  (set -x; curl -T".$f" -k -u "${BASIC_AUTH}" "https://${HOST}$f")
done

