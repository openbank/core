#!/bin/bash

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

set -e

pushd $SRC &> /dev/null
CMDS=$(find cmd -mindepth 1 -maxdepth 1 -type d|sort)
for i in $CMDS; do
  CMD=$(basename $i)
  (set -x;
    go build -o $SRC/$CMD ./$i
  )
done
popd &> /dev/null
