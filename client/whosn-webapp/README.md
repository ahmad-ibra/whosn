# whosn-webapp

A webapp client of the whosn

## Setup
Create a `.env_prod` and `.env_dev` file, each defining the `REACT_APP_BACKEND_ADDRESS` as shown in the `.env_example`.
Note that the `.env_prod` should be set to the prod whosn-core backend address.

## Running locally
Run the following command to run the app locally:

```
❯ make exec
```

## Deploying The App
Run the following command to deploy the app:

```
❯ make deploy
```
