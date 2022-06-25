whosn_dir := $(abspath $(shell git rev-parse --show-toplevel))

exec:
	npm run start

deploy:
	git push https://git.heroku.com/whosn-webapp.git HEAD:main

.PHONY: exec deploy
