#!/bin/sh
docker-compose exec hydra \
	hydra token user \
	--client-id openbank-core \
	--client-secret openbank-core-hydra-secret \
	--endpoint http://127.0.0.1:4444/ \
	--port 5555 \
	--scope openid,offline
