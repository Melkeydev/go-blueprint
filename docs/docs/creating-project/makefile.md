## Makefile Proect Management

Makefile is designed for building, running, and testing a Go project. It includes support for advanced options like HTMX and Tailwind CSS, and handles OS-specific operations for Unix-based systems (Linux/macOS) and Windows.

## Targets

**_`all`_**

The default target that builds and test the application by running the `build` and `test` target.

**_`templ-install`_**

This target installs the Go-based templating tool, `templ`, if it is not already installed. It supports:

- **Unix-based systems**: Prompts the user to install `templ` if it is missing.
- **Windows**: Uses PowerShell to check for and install `templ`.

**_`tailwind-install`_**

This target downloads and sets up `tailwindcss`, depending on the user's operating system:

- **Linux**: Downloads the Linux binary.
- **macOS**: Downloads the macOS binary.
- **Windows**: Uses PowerShell to download the Windows executable.

**_`build`_**

Builds the Go application and generates assets with `templ` and `tailwind`, if the corresponding advanced options are enabled:

- Uses `templ` to generate templates.
- Runs `tailwindcss` to compile CSS.

**_`run`_**

Runs the Go application by executing the `cmd/api/main.go` file and npm install with run dev if React flag is used.

**_`docker-run`_** and **_`docker-down`_**

These targets manage a database container:

- **Unix-based systems**: Tries Docker Compose V2 first, falls back to V1 if needed.
- **Windows**: Uses Docker Compose without version fallback.

**_`test`_**

Runs unit tests for the application using `go test`.

**_`itest`_**

Runs integration tests if a database, with the exception of SQLite, is used.

**_`clean`_**

Removes the compiled binary (`main` or `main.exe` depending on the OS).

**_`watch`_**

Enables live reload for the project using the `air` tool:

- **Unix-based systems**: Checks if `air` is installed and prompts for installation if missing.
- **Windows**: Uses PowerShell to manage `air` installation and execution.
