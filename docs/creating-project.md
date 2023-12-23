# Creating a Project

After installing the Go-Blueprint CLI tool, you can create a new project with the default settings by running the following command:

```sh
go-blueprint create
```

This command will interactively guide you through the project setup process, allowing you to choose the project name, framework, and database driver.

## Using Flags for Non-Interactive Setup

For a non-interactive setup, you can use flags to provide the necessary information during project creation. Here's an example:

```
go-blueprint create --name my-project --framework gin --driver postgres
```

In this example:

- `--name`: Specifies the name of the project (replace "my-project" with your desired project name).
- `--framework`: Specifies the Go framework to be used (e.g., "gin").
- `--driver`: Specifies the database driver to be integrated (e.g., "postgres").

Customize the flags according to your project requirements.

## Advanced Flag

By including the `--advanced` flag, users can choose one or both of the advanced features, HTMX support and GitHub Actions for CI/CD, during the project creation process. The flag enhances the simplicity of Blueprint while offering flexibility for users who require additional functionality.

To recreate the project using the same configuration non-interactively, use the following command:
```bash
go-blueprint create --name my-project --framework chi --driver mysql --advanced true
```


## Makefile Project Management

**Build the Application:**
Compiles the application and generates the executable.
```bash
make build
```

**Run the Application:**
Executes the application using `go run`
```bash
make run
```

**Create DB Container:**
Utilizes Docker Compose to set up the database container. It includes a fallback for Docker Compose V1
```bash
make docker-run
```

**Shutdown DB Container:**
Stops and removes the database container. It also has a fallback for Docker Compose V1
```bash
make docker-down
```

**Test the Application:**
Executes tests defined in the `./tests` directory
```bash
make test
```

**Clean the Binary:**
Removes the generated binary file
```bash
make clean
```

**Live Reload:**
Monitors file changes and automatically rebuilds and restarts the application using `air`
```bash
make watch
```

Using this Makefile simplifies common development tasks, making it easier to build, run, test, and manage dependencies in a Go project.