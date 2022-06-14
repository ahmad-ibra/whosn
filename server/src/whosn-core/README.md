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
❯ curl localhost:8080/api/v1/events

[
  {
    "id": 1,
    "name": "Volleyball",
    "owner_id": 1,
    "start_time": "0001-01-01T00:00:00Z",
    "location": "6Pack",
    "min_users": 10,
    "max_users": 12,
    "price": 120,
    "is_flat_rate": false,
    "link": "www.somepage.com/abasdcasdfasdf/1",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": 1,
    "name": "Soccer",
    "owner_id": 1,
    "start_time": "0001-01-01T00:00:00Z",
    "location": "Tom binnie",
    "min_users": 10,
    "max_users": 22,
    "price": 155,
    "is_flat_rate": false,
    "link": "www.somepage.com/abasdcasdfasdf/2",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": 1,
    "name": "Movie",
    "owner_id": 2,
    "start_time": "0001-01-01T00:00:00Z",
    "location": "Landmarks Guildford",
    "min_users": 1,
    "max_users": 10,
    "price": 12,
    "is_flat_rate": true,
    "link": "www.somepage.com/abasdcasdfasdf/3",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
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
