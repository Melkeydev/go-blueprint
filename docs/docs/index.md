## Go Blueprint - Ultimate Golang Blueprint Library

![logo](./public/logo.png)


Powerful CLI tool designed to streamline the process of creating Go projects with a robust and standardized structure. Not only does Go Blueprint facilitate project initialization, but it also offers seamless integration with popular Go frameworks, allowing you to focus on your application's code from the very beginning.

## Why Choose Go Blueprint?

- **Easy Setup and Installation**: Go Blueprint simplifies the setup process, making it a breeze to install and get started with your Go projects.

- **Pre-established Go Project Structure**: Save time and effort by having the entire Go project structure set up automatically. No need to worry about directory layouts or configuration files.

- **HTTP Server Configuration Made Easy**: Whether you prefer Go's standard library HTTP package, Chi, Gin, Fiber, HttpRouter, Gorilla/mux or Echo, Go Blueprint caters to your server setup needs.

- **Focus on Your Application Code**: With Go Blueprint handling the project scaffolding, you can dedicate more time and energy to developing your application logic.

## Project Structure

Here's an overview of the project structure created by Go Blueprint when all options are utilized:

```bash
/ (Root)
├── .github/
│   └── workflows/
│       ├── go-test.yml     # GitHub Actions workflow for running tests.
│       └── release.yml     # GitHub Actions workflow for releasing the application.
├── cmd/
│   ├── api/            
│   │   └── main.go         # Main file for starting the server.
│   └── web/             
│       ├── js/         
│       │   └── htmx.min.js # HTMX library for dynamic HTML content 
│       ├── base.templ      # Base HTML template file.
│       ├── base.templ.go   # Generated Go code for base template
│       ├── efs.go          # File for handling file system operations.
│       ├── hello.go        # Handler for serving "hello" endpoint.
│       └── hello.templ     # Template file for the "hello" endpoint.
|       └── hello.templ.go  # Generated Go code for the "hello" template. 
├── internal/   
│   ├── database/           
│   │   └── database.go     # File containing functions related to database operations.
│   └── server/             
│       ├── routes.go       # File defining HTTP routes.
│       └── server.go       # Main server logic.
├── tests/    
│   └── handler_test.go     # Test file for testing HTTP handlers.
├── .air.toml               # Configuration file for Air, a live-reload utility.
├── docker-compose.yml      # Docker Compose configuration for defining DB config.
├── .env                    # Environment configuration file.
├── .gitignore              # File specifying which files and directories to ignore in Git.
├── go.mod                  # Go module file for managing dependencies.
├── .goreleaser.yml         # Configuration file for GoReleaser, a tool for building and releasing binaries.
├── go.sum                  # Go module file containing checksums for dependencies.
├── Makefile                # Makefile for defining and running commands.
└── README.md               # Project's README file containing essential information about the project.
```

This structure provides a comprehensive organization of your project, separating source code, tests, configurations and documentation.




