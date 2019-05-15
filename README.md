# Go playground [![Build Status](https://travis-ci.org/ppworks/go-playground.svg?branch=master)](https://travis-ci.org/ppworks/go-playground)

## Docker

docker-compose up

## Test

docker exec -i -t go-playground_app_1 go test -v ./...

## DB

[pressly/goose: Goose database migration tool - fork of https://bitbucket.org/liamstask/goose](https://github.com/pressly/goose)

### Migration

docker exec -i -t go-playground_app_1 goose -dir=database/migrations -v postgres "postgres://postgres:@db:5432/postgres?sslmode=disable" up
