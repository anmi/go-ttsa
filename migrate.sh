#!/usr/bin/env bash

export POSTGRESQL_URL='postgres://myusername:mypassword@localhost:5432/tinytracker?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up
docker exec -it postgresql pg_dump -U myusername -s tinytracker > ./sqlc/schema.sql
docker run --rm -v $(pwd):/src -w /src/sqlc kjconroy/sqlc generate
