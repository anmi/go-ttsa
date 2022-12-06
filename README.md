## Update deps

```
go mod tidy
```

## Run locally

```
go run main.go
```

## SQLC

```
docker run --rm -v $(pwd):/src -w /src/sqlc kjconroy/sqlc generate
```

## OpenAPI

```
go run github.com/ogen-go/ogen/cmd/ogen --target api  --clean openapi.yml
```
or
```
go generate
```

## POSTGRES

```
docker exec -it postgresql psql -U myusername tinytracker
```

Dump schema
```
docker exec -it postgresql pg_dump -U myusername -s tinytracker > ./sqlc/schema.sql
```

## Migrations

https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

### Installing
Mac OS
```
brew install golang-migrate
```

### Create migration
```
migrate create -ext sql -dir db/migrations -seq create_users_table
```
### Run migration
```
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations up
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```
