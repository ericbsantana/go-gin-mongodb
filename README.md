![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ericbsantana/gogimon)
[![Go Report Card](https://goreportcard.com/badge/github.com/ericbsantana/gogimon)](https://goreportcard.com/report/github.com/ericbsantana/gogimon)

# gogimon

> Go + Gin + MongoDB = gogimon

## Overview

This is a simple and minimal template for building RESTful APIs using Gin as HTTP web framework and MongoDB as database in Go.

This template follows a dependency injection pattern to make the code modular and testable. This template will help you to spin up a new project quickly and focus on building your application logic.

It contains a sample CRUD structure with a `User` entity, including routes, handlers, repositories and data transfer objects (DTOs). Easily dettachable code for developing business logic and automated testing.

## Features

- [gin-gonic/gin](https://github.com/gin-gonic/gin): Uses Gin, a HTTP web framework for Go, useful for fast building APIs.
- [mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver): Integrates MongoDB for data storage.
- [testcontainers/testcontainers-go](https://github.com/testcontainers/testcontainers-go): Includes [Testcontainers](https://golang.testcontainers.org/) to simplify integration testing and enable testing with external services like databases, Kafka, or Redis. In this template, it is used to run a MongoDB container for testing.
- [go-playground/validator](https://github.com/go-playground/validator/): Uses validator to validate DTOs.

## Pre-requisites

- [Go](https://golang.org/dl/): Make sure you have Go installed on your machine.
- [Docker](https://www.docker.com/get-started): Required to run MongoDB container for both development and testing.
- [Air](https://github.com/cosmtrek/air/): Optional, but recommended for live reloading during development.

## Quick Start

You may want to run the application in a container for a fast setup:

```bash
git clone https://github.com/ericbsantana/gogimon.git
cd gogimon
docker-compose up -d
```

The application will be available at `http://localhost:8080`. The routes are defined in `main.go`:

```go
r.GET("/", func(c *gin.Context) {
  c.String(http.StatusOK, "OK")
})

r.GET("/users", userHandler.Find)
r.GET("/users/:id", userHandler.FindByID)
r.POST("/users", userHandler.Create)
r.PATCH("/users/:id", userHandler.Update)
r.DELETE("/users/:id", userHandler.Delete)
```

## Manual Setup

If you prefer to run the application locally, you need to have MongoDB running in a container or nativelly on your machine.

You have to set the environment variables `MONGODB_URI` to connect to your MongoDB instance. You can set them in a `.env` file in the root of the project:

```bash
# in project root
cp .env.example .env
```

Now you can run the application:

```bash
air

...
[GIN-debug] Listening and serving HTTP on :8080
...
```

Open `http://localhost:8080` and you should see a `200 OK` response.

## Project folder structure

```bash
.
├── Dockerfile # for building the application and run with compose
├── docker-compose.yml # to run the application with a MongoDB container
├── main.go # entry point of the application
└── internal
    ├── databases # contains mongo database connection
    ├── dtos # data transfer objects, createUserDto, for example
    ├── handlers # request handlers or controllers, call it what you want
    ├── models # data models, like User
    ├── repositories # data access layer to interact with the database
    ├── tests # integration tests
    │   ├── config # configuration of a testcontainer for MongoDB
    │   └── handlers # handlers tests
    └── validator # custom validators for DTOs
```

## Testing

This template includes integration tests using Testcontainers. To run the tests, use the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests to suggest improvements, report bugs, or add new features.

## Roadmap

Here's what's planned for future development of this template:

- **Authentication with JWT**: Implement JWT-based authentication to secure endpoints and manage user sessions.

- **CI/CD Pipeline**: Set up a continuous integration and continuous deployment (CI/CD) pipeline to automate the build, testing, and deployment process.

- **Logging Mechanism**: Introduce a logging mechanism to provide insight into the application's behavior.

## License

This project is licensed under the [MIT License](LICENSE).
