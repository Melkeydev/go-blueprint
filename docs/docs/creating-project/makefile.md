## Makefile Project Management

This Makefile is included as a default after project creation. It offers a set of commands to simplify various development tasks for managing a Go project.

## Commands

- **Build the Application:**
Compiles the application and generates the executable.
```bash
make build
```

- **Run the Application:**
Executes the application using `go run`.
```bash
make run
```

- **Create DB Container:**
Utilizes Docker Compose to set up the database container. It includes a fallback for Docker Compose V1.
```bash
make docker-run
```

- **Shutdown DB Container:**
Stops and removes the database container. It also has a fallback for Docker Compose V1.
```bash
make docker-down
```

- **Test the Application:**
Executes tests defined in the `./tests` directory.
```bash
make test
```

- **Clean the Binary:**
Removes the generated binary file.
```bash
make clean
```

- **Live Reload:**
Monitors file changes and automatically rebuilds and restarts the application using `air`.
```bash
make watch
```

Makefile simplifies common development tasks, making it easier to build, run, test, and manage dependencies in a Go project. It enhances productivity by providing a standardized approach to project management.