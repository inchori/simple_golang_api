# grpc_identity

Build a Simple Blog Service with gRPC and goFiber

## TODO
### REST API
- [X] Configure a database with [entgo](https://entgo.io/)
- [X] User CRUD API
- [X] Blog CRUD API
- [X] Configure Edges Between User and Post
- [X] User Login API with JWT

### gRPC
- [X] Support GRPC
- [X] gRPC Interceptor for JWT
- [X] Apply gRPC-gateway

### ETC
- [X] GitHub Actions CI
- [X] Error Handle with logging
- [X] Swagger Docs

## Build From Source

### Prerequisite

- Go version 1.20 or greater
- MySQL

1. Clone this Repository
```bash
git clone `{this repository URL}`

cd grpc_identity
```
2. Configure .env file
```bash
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=@
DB_NAME=blog

# grpc or http
SERVER=grpc
HTTP_PORT=
GRPC_PORT=
GATEWAY_PORT=
GATEWAY_ENABLED= # If you want to enable, grpc-gateway write `true`

```

3. Build and Run
```bash
make build

make clean build // If the codes updated, use this command

make test // Check the Unit test

./main // Run Go Binary File
```

## Swagger Docs

If you configure the server to `http` in `.env`, you can check swagger docs in `http://localhost:{your_http_port}/swagger/index.html`.
