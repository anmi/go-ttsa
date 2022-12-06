#!/usr/bin/env bash

docker exec -it postgresql pg_dump -U myusername -s tinytracker > ./sqlc/schema.sql