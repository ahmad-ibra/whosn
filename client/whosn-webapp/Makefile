exec:
	cp .env_dev .env
	npm run start-dev

build:
	cp .env_prod .env
	npm run build

deploy:
	git push https://git.heroku.com/whosn-webapp.git HEAD:main

.PHONY: exec build deploy
