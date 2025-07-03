# Justfile Project Management

`just` is an alternative task runner to Make, focused on simplicity and ease of use with a more modern syntax. This project includes a `Justfile` that mirrors much of the Makefile functionality for managing, building, and testing a Go application, with support for optional frontend tooling like HTMX and Tailwind CSS.

## Overview

The `justfile` provides shorthand commands for common development tasks, making it easy to:

- Build and run the project
- Install necessary tools (`templ`, `tailwindcss`, etc.)
- Manage frontend assets
- Run tests and clean the build environment
- Use Docker for database containers
- Enable live reload for development

## Recipes

### `all`

Builds the Go application and runs tests. Equivalent to running `just build` followed by `just test`.

### `templ-install`

Installs the Go-based templating tool `templ`, using OS-specific logic:

- **Unix-based systems**: Uses shell commands to check and install.
- **Windows**: Uses PowerShell to install if not present.

### `tailwind-install`

Downloads and installs the correct version of `tailwindcss`:

- **Linux/macOS**: Downloads the appropriate binary based on OS.
- **Windows**: Uses PowerShell to fetch the executable.

### `build`

Builds the Go application and generates frontend assets:

- Uses `templ` for template generation.
- Compiles Tailwind CSS if enabled.

### `run`

Runs the main Go application from `cmd/api/main.go`. If the React flag is active, it also runs `npm install` and `npm run dev`.

### `docker-run` and `docker-down`

Manages a Docker container for the database:

- **Unix**: Prefers Docker Compose V2, falls back to V1.
- **Windows**: Uses Docker Compose without fallback logic.

### `test`

Runs unit tests with `go test`.

### `itest`

Executes integration tests for projects that use databases other than SQLite.

### `clean`

Removes the built binary (`main` or `main.exe`, depending on platform).

### `watch`

Starts a development server with live reload using the `air` tool:

- **Unix**: Checks for `air` and installs if missing.
- **Windows**: Uses PowerShell to handle `air` setup and execution.

---

For developers who prefer `just` over `make`, this `Justfile` offers a clean and minimal way to manage the entire Go project lifecycle.
