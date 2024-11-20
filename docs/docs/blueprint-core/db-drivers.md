To extend the project with database functionality, users can choose from a variety of Go database drivers. Each driver is tailored to work with specific database systems, providing flexibility based on project requirements:

1. [Mongo](https://go.mongodb.org/mongo-driver): Provides necessary tools for connecting and interacting with MongoDB databases.
2. [Mysql](https://github.com/go-sql-driver/mysql): Enables seamless integration with MySQL databases.
3. [Postgres](https://github.com/jackc/pgx/): Facilitates connectivity to PostgreSQL databases.
4. [Redis](https://github.com/redis/go-redis): Provides tools for connecting and interacting with Redis.
5. [Sqlite](https://github.com/mattn/go-sqlite3): Suitable for projects requiring a lightweight, self-contained database. and interacting with Redis
6. [ScyllaDB](https://github.com/scylladb/gocql): Facilitates connectivity to ScyllaDB databases.

## Updated Project Structure

Integrating a database adds a new layer to the project structure, primarily in the `internal/database` directory:

```bash
/(Root)
├── /cmd
│   └── /api
│       └── main.go
├── /internal
│   ├── /database
│   │   ├── database_test.go
│   │   └── database.go
│   └── /server
│       ├── routes.go
│       ├── routes_test.go
│       └── server.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Database Driver Implementation

Users can select the desired database driver based on their project's specific needs. The chosen driver is then imported into the project, and the `database.go` file is adjusted accordingly to establish a connection and manage interactions with the selected database.

## Integration Tests for Database Operations

For all the database drivers but the `Sqlite`, integration tests are automatically generated to ensure that the database connection is working correctly. It uses [Testcontainers for Go](https://golang.testcontainers.org/) to spin up a containerized instance of the database server, run the tests, and then tear down the container.

[Testcontainers for Go](https://golang.testcontainers.org/) is a Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.


### Requirements

You need a container runtime installed on your machine. Testcontainers supports Docker and any other container runtime that implements the Docker APIs.

To install Docker:

```bash
curl -sLO get.docker.com
```

### Running the tests

Go to the `internal/database` directory and run the following command:

```bash
go test -v
```

or just run the following command from the root directory:

```bash
make itest
```

Testcontainers automatically pulls the required Docker images and start the containers. The tests run against the containers, and once the tests are done, the containers are stopped and removed. For further information, refer to the [official documentation](https://golang.testcontainers.org/).

## Docker-Compose for Quick Database Spinup

To facilitate quick setup and testing, a `docker-compose.yml` file is provided. This file defines a service for the chosen database system with the necessary environment variables. Running `docker-compose up` will quickly spin up a containerized instance of the database, allowing users to test their application against a real database server.

This Docker Compose approach simplifies the process of setting up a database for development or testing purposes, providing a convenient and reproducible environment for the project.
