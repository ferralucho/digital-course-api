# digital-course-api

Go api for courses, clean architecture, postgre and firebase implementation

Firebase middleware is disabled for the moment, uncomment 	//handler.Use(middleware.AuthMiddleware) in router.

## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)
- [Dependency Injection](#dependency-injection)

## Quick start

For migrations: https://github.com/golang-migrate/migrate

Local development:
```sh
# Postgres
$ make compose-up
# Run app with migrations
$ make run
```

Integration tests (can be run in CI):
```sh
# DB, app + migrations, integration tests
$ make compose-up-integration-test
```

## Project structure
### `cmd/app/main.go`
Configuration and logger initialization. Then the main function "continues" in
`internal/app/app.go`.

### `config`
Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in yaml, or in environment variables).

### `docs`
Swagger documentation. Auto-generated by [swag](https://github.com/swaggo/swag) library.

### `integration-test`
Integration tests.
They are launched as a separate container, next to the application container.

### `internal/app`

This is where all the main objects are created.
Dependency injection occurs through the "New ..." constructors (see Dependency Injection).
This technique allows us to layer the application using the [Dependency Injection](#dependency-injection) principle.
This makes the business logic independent from other layers.

Next, we start the server and wait for signals in _select_ for graceful completion.
If `app.go` starts to grow, you can split it into multiple files.

For a large number of injections, [wire](https://github.com/google/wire) can be used.

The `migrate.go` file is used for database auto migrations.
It is included if an argument with the _migrate_ tag is specified.
For example:

```sh
$ go run -tags migrate ./cmd/app
```

### `internal/controller`
Server handler layer (MVC controllers). The template shows a  server:
- REST http (Gin framework)

Server routers are written in the same style:
- Handlers are grouped by area of application (by a common basis)
- For each group, its own router structure is created, the methods of which process paths
- The structure of the business logic is injected into the router structure, which will be called by the handlers

#### `internal/controller/http`
For v2, we will need to add the `http/v2` folder with the same content.
And in the file `internal/app` add the line:
```go
handler := gin.New()
v1.NewRouter(handler, t)
```

In `v1/router.go` and above the handler methods, there are comments for generating swagger documentation using [swag](https://github.com/swaggo/swag).

### `internal/entity`
Entities of business logic (models) can be used in any layer.
There can also be methods, for example, for validation.

### `internal/usecase`
Business logic.
- Methods are grouped by area of application (on a common basis)
- Each group has its own structure
- One file - one structure

Repositories and other business logic structures are injected into business logic structures
(see [Dependency Injection](#dependency-injection)).

#### `internal/usecase/repo`
A repository is an abstract storage (database) that business logic works with.

## Dependency Injection
In order to remove the dependence of business logic on external packages, dependency injection is used.

For example, through the New constructor, we inject the dependency into the structure of the business logic.
This makes the business logic independent (and portable).
We can override the implementation of the interface without making changes to the `usecase` package.

```go
package usecase

import (
    // Nothing!
)

type Repository interface {
    Get()
}

type UseCase struct {
    repo Repository
}

func New(r Repository) *UseCase{
    return &UseCase{
        repo: r,
    }
}

func (uc *UseCase) Do()  {
    uc.repo.Get()
}
```

It will also allow us to do auto-generation of mocks and easily write unit tests.  https://github.com/golang/mock

> We are not tied to specific implementations in order to always be able to change one component to another.
> If the new component implements the interface, nothing needs to be changed in the business logic.


