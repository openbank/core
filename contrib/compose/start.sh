#!/bin/bash

REPO_ROOT=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd))
set -e

if [ ! -f "$REPO_ROOT/api/v1/_gunk/gen/sql/postgres.sql" ]; then
	echo "Schema not found. Preparing repository."
	pushd "$REPO_ROOT" &> /dev/null
	(set -x;
		./install.sh
		go generate
	)
	popd &> /dev/null
fi

pushd "$REPO_ROOT/contrib/compose"
(set -x
	cp ../../api/v1/_gunk/gen/sql/postgres.sql ./schema.sql
	docker-compose up -d --build
	docker-compose exec hydra \
		hydra clients create \
		--skip-tls-verify \
		--endpoint http://localhost:4445 \
		--id openbank-core \
		--secret openbank-core-hydra-secret \
		--grant-types authorization_code,refresh_token \
		--response-types code \
		--scope openid,offline \
		--callbacks http://127.0.0.1:5555/callback
)
popd &> /dev/null
