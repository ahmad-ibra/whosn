whosn_dir := $(abspath $(shell git rev-parse --show-toplevel))

build:
	GOOS=linux go build cmd/main.go
	mv main ../../bin/whosn-core

deploy:
	git push https://git.heroku.com/whosn-core.git HEAD:main

.PHONY: build deploy
