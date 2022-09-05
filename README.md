# About

Contains the API and database definitions, server implementations, et al, for
the OpenBank Core Engine.

[View API docs][docs].

[redoc]: https://api.corebank.brankas.dev/

## Using

```sh
# Generate docs
$ pushd docs
$ pnpm install
$ pnpm build
$ popd

# install dependencies
$ ./install.sh

# building
$ ./build.sh

# run
$ ./openbank-core &
$ ./openbank-core-gw &

# grpcurl
$ grpcurl localhost:9090 list
```

## Docker Compose

```sh
# generate files with gunk
$ pushd api/v1
$ gunk generate ./...
$ popd

# start services
$ cd compose
$ docker-compose up -d --build
```
