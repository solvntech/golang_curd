# Golang Demo API

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://go.dev/)
[![Gin reference](https://img.shields.io/badge/Gin-dependency-blue?logo=github&logoColor=white)](https://github.com/gin-gonic/gin)
[![Gorm reference](https://img.shields.io/badge/Gorm-dependency-blue?logo=github&logoColor=white)](https://github.com/go-gorm/gorm)
[![Docker reference](https://img.shields.io/badge/Docker-tool-102C66?logo=docker&logoColor=white)](https://www.docker.com/)

This repository contains CRUD example
demonstrate connect database, find data, create, update and delete

### Clone the project

```
$ git https://github.com/solvntech/golang_curd.git
$ cd golang_curd
```

### Config environment

_Create `.env` file and copy `.env.example` to `.env`, after that modify according your configuration_

Example:

```.dotenv
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_db_pass
POSTGRES_DB_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=db_name
```

### Setup docker

```
$ docker compose up -d
```

### Install dev environment and run project

```
# install live host server at global

$ go install github.com/cosmtrek/air@latest

# install dependencies

$ go mod download && go mod tidy

# run project

$ air
```

![Screenshot](images/img.png)
