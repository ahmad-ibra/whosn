whosn_dir := $(abspath $(shell git rev-parse --show-toplevel))

mocks:
	mockery --case underscore --name=Storer --dir ./internal/data --output ./internal/data/mocks

db:
	docker-compose up --build -d

db-clean:
	docker-compose down && docker-compose up --build -d

test:
	go test ./...

exec:
	go run cmd/main.go

build:
	GOOS=linux go build cmd/main.go
	mv main bin/whosn-core

deploy:
	git push https://git.heroku.com/whosn-core.git HEAD:main

.PHONY: mocks db db-clean test exec build deploy
