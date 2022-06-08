whosn_dir := $(abspath $(shell git rev-parse --show-toplevel))

build:
	GOOS=linux go build -o $(whosn_dir)/bin

deploy:
	git push https://git.heroku.com/whosn-core.git HEAD:main

.PHONY: build deploy
