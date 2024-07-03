[Testcontainers for Go](https://golang.testcontainers.org/) is a Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.


The project tree would look like this:
```bash
/ (Root)
├── .github/
│   └── workflows/
│       ├── go-test.yml           # GitHub Actions workflow for running tests.
│       └── release.yml           # GitHub Actions workflow for releasing the application.
├── cmd/
│   ├── api/
│   │   └── main.go               # Main file for starting the server.
│   └── web/
│       ├── assets/
│       │   ├── css/
│       │   │   ├── input.css     # Tailwind input file for compiling output.css with CLI
│       │   │   └── output.css    # Generated CSS file.
│       │   └── js/
│       │       └── htmx.min.js   # HTMX library for dynamic HTML content.
│       ├── base.templ            # Base HTML template file.
│       ├── base_templ.go         # Generated Go code for base template
│       ├── efs.go                # Includes assets into compiled binary.
│       ├── hello.go              # Logic for handling "hello" form.
│       ├── hello.templ           # Template file for the "hello" endpoint.
│       └── hello_templ.go        # Generated Go code for the "hello" template. 
├── internal/
│   ├── database/
│   │   └── database_test.go      # File containing integrations tests for the database operations.
│   │   └── database.go           # File containing functions related to database operations.
│   └── server/
│       ├── routes.go             # File defining HTTP routes.
│       └── server.go             # Main server logic.
├── tests/
│   └── handler_test.go           # Test file for testing HTTP handlers.
├── .air.toml                     # Configuration file for Air, a live-reload utility.
├── docker-compose.yml            # Docker Compose configuration for defining DB config.
├── .env                          # Environment configuration file.
├── .gitignore                    # File specifying which files and directories to ignore in Git.
├── go.mod                        # Go module file for managing dependencies.
├── .goreleaser.yml               # Configuration file for GoReleaser, a tool for building and releasing binaries.
├── go.sum                        # Go module file containing checksums for dependencies.
├── Makefile                      # Makefile for defining and running commands.
├── tailwind.config.js            # Tailwind CSS configuration file.
└── README.md                     # Project's README file containing essential information about the project.
```

## Requirements

You need a container runtime installed on your machine. Testcontainers supports Docker and any other container runtime that implements the Docker APIs.

To install Docker:

```bash
curl -sLO get.docker.com
```

## Running the tests

Go to the `internal/database` directory and run the following command:

```bash
go test -v
```

Testcontainers automatically downloads the required Docker images and start the containers. The tests run against the containers, and once the tests are done, the containers are stopped and removed. For further information, refer to the [official documentation](https://golang.testcontainers.org/).
