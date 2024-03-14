## HTMX and Templ Setup

The WEB directory contains the web-related components and assets for the project. It leverages htmx and tmpl in Go for dynamic web content generation.

## Structure

```
web/
│
├── js/
│   └── htmx.min.js         # htmx library for dynamic HTML content
│
├── base.templ              # Base template for HTML structure
├── base_templ.go           # Generated Go code for base template
├── efs.go                  # Embeds static files (e.g., JavaScript) into the Go binary
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

## Templating

Templates are generated using the `templ generate` command after project creation. These templates are then compiled into Go code for efficient execution.

You can test HTMX functionality on `localhost:port/web` endpoint.