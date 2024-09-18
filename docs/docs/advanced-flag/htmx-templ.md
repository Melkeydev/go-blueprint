The WEB directory contains the web-related components and assets for the project. It leverages [htmx](https://github.com/bigskysoftware/htmx) and [templ](https://github.com/a-h/templ) in Go for dynamic web content generation.

## Structure

```
web/
│
│
├── assets/
│   └── js/
│       └── htmx.min.js     # htmx library for dynamic HTML content
│
├── base.templ              # Base template for HTML structure
├── base_templ.go           # Generated Go code for base template
├── efs.go                  # Embeds static files into the Go binary
│
├── hello.go                # Handler for the Hello Web functionality
├── hello.templ             # Template for rendering the Hello form and post data
└── hello_templ.go          # Generated Go code for hello template
```

## Usage

- **Navigate to Project Directory:**
```bash
cd my-project
```

- **Install Templ CLI:**
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

- **Generate Templ Function Files:**
```bash
templ generate
```

- **Start Server:**
```bash
make run
```

## Makefile

Automates templ with Makefile entries, which are automatically created if the htmx advanced flag is used.
It detects if templ is installed or not and generates templates with the make build command.
Both Windows and Unix-like OS are supported.

```bash
all: build

templ-install:
	@if ! command -v templ > /dev/null; then \
		read -p "Go's 'templ' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/a-h/templ/cmd/templ@latest; \
			if [ ! -x "$$(command -v templ)" ]; then \
				echo "templ installation failed. Exiting..."; \
				exit 1; \
			fi; \
		else \
			echo "You chose not to install templ. Exiting..."; \
			exit 1; \
		fi; \
	fi

build: templ-install
	@echo "Building..."
	@templ generate
	@go build -o main cmd/api/main.go
```

## Templating

Templates are generated using the `templ generate` command after project creation. These templates are then compiled into Go code for efficient execution.

You can test HTMX functionality on `localhost:PORT/web` endpoint.
