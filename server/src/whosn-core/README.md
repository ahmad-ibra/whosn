# Server
The backend of whosn

## Testing App Locally
As a pre-req you'll need to have go installed.

Once installed, run the following command to run the app locally:
```
❯ make exec
```
You can then test the server using curl or postman:
```
❯ curl localhost:8080/_hc

{"status":"healthy"}
```

## Creating The App
Run the following to first deploy to heroku (note this will only need to be done once):
```
❯ heroku create -a whosn-core
❯ heroku buildpacks:add -a whosn-core https://github.com/heroku/heroku-buildpack-multi-procfile
❯ heroku config:set -a whosn-core PROCFILE=server/src/whosn-core/Procfile
❯ git push https://git.heroku.com/whosn-core.git HEAD:main
```

## Deploying The App
You can leverage the Makefile in this directory to easily deploy the app.
To build the artifact which will be deployed, simply run:
```
❯ make build
```
This will create a binary which will be stored in the `whosn/server/bin` directory.
Next, commit and push the binary to github.
Finally, run the following command to deploy the app:
```
❯ make deploy
```

## Endpoints
Below is a list of all the available endpoints:
```
GET    /_hc
POST   /api/v1/login
POST   /api/v1/user
GET    /api/v1/secured/users
DELETE /api/v1/secured/user/:id
PUT    /api/v1/secured/user/:id
GET    /api/v1/secured/user/:id
GET    /api/v1/secured/events
GET    /api/v1/secured/events/owned
GET    /api/v1/secured/events/joined
DELETE /api/v1/secured/event/:id
PUT    /api/v1/secured/event/:id
GET    /api/v1/secured/event/:id
POST   /api/v1/secured/event
GET    /api/v1/secured/event_users
GET    /api/v1/secured/event/:id/join
GET    /api/v1/secured/event/:id/leave
```
All endpoints under `/api/v1/secured/` require an `Authorization` header with a valid token. This token is provided on login.
