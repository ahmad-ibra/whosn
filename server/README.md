# Server
The backend of whosn

## Testing App Locally
To start a local container, run the following commands:
```
❯ cd whosn-core
❯ docker build -t whosn-core .
❯ docker run -p 8080:8080 -d whosn-core
```
You can then test the server using curl or postman:
```
❯ curl localhost:8080/api/ping
{"message":"pong"}%
```