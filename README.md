# GRPC Microservices starter

## Setup

This repository is used in the `make` command, and selected when choosing which command you wish to use.

### Database

Create an `.env` file in root directory. Example content of it is in `.env.example` file.

Execute gorm migrate to populate database:

```
$ make migrate-db
```

### Develop

To develop all apps and packages, run the following command:

```
# To update this dev setup to work properly
cd go-grpc-microservices
make run
```

### Build

To build the apps, run the following command:

```
cd go-grpc-microservices
make docker-build
```
