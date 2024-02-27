## Database Drivers

To extend the project with database functionality, users can choose from a variety of Go database drivers. Each driver is tailored to work with specific database systems, providing flexibility based on project requirements:

1. [Mysql](https://github.com/go-sql-driver/mysql): Enables seamless integration with MySQL databases.
2. [Postgres](https://github.com/jackc/pgx/): Facilitates connectivity to PostgreSQL databases.
3. [Sqlite](https://github.com/mattn/go-sqlite3): Suitable for projects requiring a lightweight, self-contained database.
4. [Mongo](https://go.mongodb.org/mongo-driver): Provides necessary tools for connecting and interacting with MongoDB databases.
5. [Redis](https://github.com/redis/go-redis): Provides tools for connectiong and interacting with Redis

## Updated Project Structure

Integrating a database adds a new layer to the project structure, primarily in the `internal/database` directory:

```bash
/(Root)
├── /cmd
│   └── /api
│       └── main.go
├── /internal
│   ├── /database
│   │   └── database.go
│   └── /server
│       ├── routes.go
│       └── server.go
├── /tests
│   └── handler_test.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Database Driver Implementation

Users can select the desired database driver based on their project's specific needs. The chosen driver is then imported into the project, and the `database.go` file is adjusted accordingly to establish a connection and manage interactions with the selected database.

## Docker-Compose for Quick Database Spinup

To facilitate quick setup and testing, a `docker-compose.yml` file is provided. This file defines a service for the chosen database system with the necessary environment variables. Running `docker-compose up` will quickly spin up a containerized instance of the database, allowing users to test their application against a real database server.

This Docker Compose approach simplifies the process of setting up a database for development or testing purposes, providing a convenient and reproducible environment for the project.
