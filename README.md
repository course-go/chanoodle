[![Go Report Card](https://goreportcard.com/badge/github.com/course-go/chanoodle)](https://goreportcard.com/report/github.com/course-go/chanoodle)
![Go version](https://img.shields.io/github/go-mod/go-version/course-go/chanoodle)

# Chanoodle

Chanoodle is an EPG, channel and event metadata service. It exposes this data using REST API.

<img src="assets/gopher-noodles.svg" alt="Gopher noodles" width="300">

Illustration credit: [MariaLetta](https://github.com/MariaLetta/free-gophers-pack)

## Architecture

The application internally uses the [Hexagonal architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>) (also known as the "Port and adapters", "Onion architecture" or "Clean architecture" - although these can deviate, they are basically the same thing).

<img src="assets/architecture.svg" alt="Chanoodle architecture" width="1200">

### Layers

If this is your first time dealing with such architecture, here is a brief overview of the layers and their responsibilities:

- The **API layer** is responsible for exposing the application functionality to the outside world. In Chanoodle, this is implemented using a REST API.
- The business layer (no layer is actually called that, it is just an umbrella term) is made up of the application and domain layers.
  - The **Domain layer** contains the core business logic and rules. It is independent of any other layers and should not have any dependencies on them. Apart from interfaces (ports) implemented by other layers, it defines:
    - Entities - the main data structures used in the application. All entities can be uniquely identified using an ID.
    - Value objects - data structures that do not have a unique identity and are defined by their attributes.
    - Domain services - stateless services that encapsulate business logic that does not naturally fit within an entity or value object.
  - The **Application layer** contains application-specific business logic. It orchestrates the interactions between the domain layer and the external layers (API, persistence). It defines use cases that represent the main functionalities of the application.
- In the **Infrastructure layer**, we have the implementations of the interfaces defined in the domain layer. In Chanoodle, this includes the persistence layer that interacts with the database.

## Assignment

Chanoodle already implements a REST API exposing the basic functionality.
However, as this project was only a POC (proof-of-concept) the developers only
implemented an in-memory data store.

You goal will be to implement an actual persistent database layer, containerize the
application and add database dependency using Docker compose.

As with the router technology in previous assignment, the choice
of the "persistence" library is again up to you.
Here are the technologies that were presented in lectures:

- [databases/sql](https://pkg.go.dev/database/sql) with a driver of your choice
- [sqlx](https://github.com/jmoiron/sqlx)
- [sqlc](https://github.com/sqlc-dev/sqlc)
- [GORM](https://github.com/go-gorm/gorm)

Whatever your choice will be, document it in the `REASONING.md` file.
Describe why you chose the library, what did you like about it, what
additional features it offers compared to the standard library etc.

### Specification

#### Containerization

##### Container image for API

Create a [Dockerfile](https://docs.docker.com/reference/dockerfile/) for
building the container image for the API.

The Dockerfile should be multistage to decrease the image size, that is:

- it should only contain the final binary
- it shouldn't contain any source code or Go tooling such as the compiler etc.

##### Compose for running the application

Create a [Docker Compose](https://docs.docker.com/compose/) for
running the application with all of its dependencies.

The compose should contain two service definitions:

- one service for the API built using the Dockerfile created earlier and
- one service for the database.

Note that you will need to mount a volume for the database to persist its data.

#### Persistence

Add a support for an actual relational database.
For this, you can choose whatever relational database system you like.
As previously mentioned, the library used is also up to you.

As with the library, provide a reasoning for the chosen database system in
the `REASONING.md` file.

Please do not use file based solutions like [SQLite](https://sqlite.org/).
Our goal here, among others, is to practice multi-service communication using Compose
and using file based solutions that do not run their own process beats the purpose.

## Requirements

The application is runnable by executing `docker compose up` command
and it persists its data using the database system and the library of your choice.
The storage implementation choice can also be configured using the configuration
file.

## Motivation

The goal of this homework is to practice implementing a persistence layer
using the Go ecosystem and its libraries and to containerize and setup a basic
compose file for the application and its dependencies.
