# Go Blueprint - Ultimate Golang Blueprint Library

![logo](./public/logo.png)


Powerful CLI tool designed to streamline the process of creating Go projects with a robust and standardized structure. Not only does Go Blueprint facilitate project initialization, but it also offers seamless integration with popular Go frameworks, allowing you to focus on your application's code from the very beginning.

## Why Choose Go Blueprint?

- **Easy Setup and Installation**: Go Blueprint simplifies the setup process, making it a breeze to install and get started with your Go projects.

- **Pre-established Go Project Structure**: Save time and effort by having the entire Go project structure set up automatically. No need to worry about directory layouts or configuration files.

- **HTTP Server Configuration Made Easy**: Whether you prefer Go's standard library HTTP package, Chi, Gin, Fiber, HttpRouter, Gorilla/mux, Echo or Bone, Go Blueprint caters to your server setup needs.

- **Focus on Your Application Code**: With Go Blueprint handling the project scaffolding, you can dedicate more time and energy to developing your application logic.

## Project Structure

Here's an overview of the project structure created by Go Blueprint when all options are utilized:

```textfile
/ (Root)
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── web/
│       ├── js/
│       │   └── htmx.min.js
│       ├── base.templ
│       ├── efs.go
│       ├── hello.go
│       └── hello.templ
├── .github/
│   └── workflows/
│       ├── go-test.yml
│       └── release.yml
├── internal/
│   ├── database/
│   │   └── database.go
│   └── server/
│       ├── routes.go
│       └── server.go
├── tests/
│   └── handler_test.go
├── .air.toml
├── docker-compose.yml
├── .env
├── .gitignore
├── go.mod
├── .goreleaser.yml
├── go.sum
├── Makefile
└── README.md
```

- **`cmd/`**: Contains the entry points for your application.

  - **`api/`**: Main file (`main.go`).

  - **`web/`**: Main files for the web part of your application, including JavaScript (`js/`), templates (`base.templ`, `efs.go`, `hello.go`, `hello.templ`).

- **`.github/`**: GitHub Actions workflows for testing and releasing (`go-test.yml`, `release.yml`).

- **`internal/`**: Internal packages or modules of your application.

  - **`database/`**: `database.go` file for handling database-related functionality.

  - **`server/`**: Files related to the server, such as `routes.go` for defining routes and `server.go` for server-related logic.

- **`tests/`**: Test files, with `handler_test.go` as an example.

- **`.air.toml`**: Configuration file for [Air](https://github.com/cosmtrek/air), a live-reload utility for Go.

- **`docker-compose.yml`**: Configuration file for Docker Compose, defining DB config.

- **`.env`**: Environment configuration file.

- **`.gitignore`**: Gitignore file specifying which files and directories to ignore.

- **`go.mod` and `go.sum`**: Files used by Go modules to manage dependencies.

- **`.goreleaser.yml`**: Configuration file for [GoReleaser](https://goreleaser.com/), a tool for building and releasing Go binaries.

- **`Makefile`**: A Makefile for defining and running tasks or commands.

- **`README.md`**: The project's README file, containing essential information about the project, how to run it, and any other relevant details.

This structure provides a comprehensive organization of your project, separating source code, tests, configurations and documentation.




