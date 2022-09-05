#!/bin/bash

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

set -e

# pushd $SRC &> /dev/null
# FILES=$(find api/v*{,/_gunk/gen} -iname \*.go)
# if [ ! -z "$FILES" ]; then
#   (set -x;
#     rm -rf $FILES
#   )
# fi

FILES=$(find tools/db -iname \*.xo.go)
if [ ! -z "$FILES" ]; then 
  (set -x;
    rm -rf $FILES
  )
fi

go run ./scripts/clean-services ./tools/services

popd &> /dev/null
