#! /usr/bin/env bash

set -x
rm *.xo.*
xo schema file://../../db.sqlite -o . -d templates/storage --exclude "schema_migrations" $@
