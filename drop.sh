#!/usr/bin/env bash

export POSTGRESQL_URL='postgres://myusername:mypassword@localhost:5432/tinytracker?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations drop
