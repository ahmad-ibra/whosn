whosn_dir := $(abspath $(shell git rev-parse --show-toplevel))

mocks:
	mockery --case underscore --name=Storer --dir ./internal/data --output ./internal/data/mocks

test:
	go test ./...

exec:
	go run cmd/main.go

build:
	GOOS=linux go build cmd/main.go
	mv main bin/whosn-core

deploy:
	git push https://git.heroku.com/whosn-core.git HEAD:main

.PHONY: test exec build deploy
