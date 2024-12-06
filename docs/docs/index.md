---
hide:
  - toc
---
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
│       ├── go-test.yml           # GitHub Actions workflow for running tests.
│       └── release.yml           # GitHub Actions workflow for releasing the application.
├── cmd/
│   ├── api/
│   │   └── main.go               # Main file for starting the server.
│   └── web/
│       ├── styles/               # only for generating css will not be served public
│       │   └── input.css         # Tailwind input file for compiling output.css with CLI when HTMX is used
│       ├── assets/
│       │   ├── css/
│       │   │   └── output.css    # Generated CSS file.
│       │   └── js/
│       │       └── htmx.min.js   # HTMX library for dynamic HTML content.
│       ├── base.templ            # Base HTML template file.
│       ├── base_templ.go         # Generated Go code for base template.
│       ├── efs.go                # Includes assets into compiled binary.
│       ├── hello.go              # Logic for handling "hello" form.
│       ├── hello.templ           # Template file for the "hello" endpoint.
│       └── hello_templ.go        # Generated Go code for the "hello" template.
├── frontend/                     # React advanced flag. Excludes HTMX.
│   ├── node_modules/             # Node dependencies.
│   ├── public/
│   │   ├── index.html
│   │   └── favicon.ico
│   ├── src/                      # React source files.
│   │   ├── App.tsx               # Main React component.
│   │   ├── assets/               # React assets directory
│   │   │   └── logo.svg
│   │   ├── components/           # React components directory.
│   │   │   ├── Header.tsx
│   │   │   └── Footer.tsx
│   │   ├── styles/               # CSS/SCSS styles directory.
│   │   │   └── global.css
│   │   └── index.tsx             # Main entry point for React
│   ├── eslint.config.js          # ESLint configuration file.
│   ├── index.html                # Base HTML template.
│   ├── package.json              # Node.js package configuration.
│   ├── package-lock.json         # Lock file for Node.js dependencies.
│   ├── README.md                 # README file for the React project.
│   ├── tsconfig.app.json         # TypeScript configuration for the app.
│   ├── tsconfig.json             # Root TypeScript configuration.
│   ├── tsconfig.node.json        # TypeScript configuration for Node.js.
│   └── vite.config.ts            # Vite configuration file.
├── internal/
│   ├── database/
│   │   ├── database_test.go      # File containing integration tests for the database operations.
│   │   └── database.go           # File containing functions related to database operations.
│   └── server/
│       ├── routes.go             # File defining HTTP routes.
│       ├── routes_test.go        # Test file for testing HTTP handlers.
│       └── server.go             # Main server logic.
├── .air.toml                     # Configuration file for Air, a live-reload utility.
├── docker-compose.yml            # Docker Compose configuration.
├── Dockerfile                    # Dockerfile configuration for the Go project.
├── .env                          # Environment configuration file.
├── .gitignore                    # File specifying which files and directories to ignore in Git.
├── go.mod                        # Go module file for managing dependencies.
├── .goreleaser.yml               # Configuration file for GoReleaser, a tool for building and releasing binaries.
├── go.sum                        # Go module file containing checksums for dependencies.
├── Makefile                      # Makefile for defining and running commands.
├── tailwind.config.js            # Tailwind CSS configuration file for HTMX.
└── README.md                     # Project's README file containing essential information about the project.

```

This structure provides a comprehensive organization of your project, separating source code, tests, configurations and documentation.
