# Bank

This is an example application that simulates a bank, albieit a very bad one, with event sourcing and CQRS using Kafka and Go.

## Features

- Accounts
- Transfers

## Dependencies
This is what you need to be able to run this application.
- Docker
- Bazel

## How to run

First compile the application using bazel and then publish the docker images to local development using this following command:

```shell
bazel run //deployments/local:docker
```

Then run docker compose with the compose file found in deployments/local-compose.yaml

```shell
docker-compose -f deployments/local/docker-compose.yaml up -d
```

## Technology
- Docker
- Kafka
- Go
- Bazel
- Linkerd
- Postgres
- gRPC

# TODO

- [] Create ui
- [] HTTP API
- [] Integrate with EventStore