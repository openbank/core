#!/bin/bash

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

set -e

pushd $SRC &> /dev/null

# Setup some files embedded to allow for querying tools without errors.
mkdir -p ./api/v1/_gunk/gen/json/core
touch ./api/v1/_gunk/gen/json/core/all.swagger.json

# Workaround gunk bug: Directory must exist for docs.
mkdir -p docs/content/

# Go install all tools.
TOOLS=$(go list -tags tools -f '{{ join .Imports "\n" }}' tools/tools.go)
(set -x;
  go install $TOOLS
)

popd &> /dev/null
