When SQLC support is enabled, the project includes a configuration file and example SQL files. Running `sqlc generate` produces type-safe Go code for executing your queries using the standard `database/sql` interface by default. However, SQLC  is also compatible with `pgx` v4 and v5.

SQLC supports PostgreSQL, MySQL, and SQLite.

## Project Structure

The generated project will follow this structure:

```bash
/ (Root)
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── database/
│   │   ├── database.go
│   │   ├── database_test.go
│   │   ├── repository/
│   │   │   ├── db.go
│   │   │   ├── models.go
│   │   │   ├── querier.go
│   │   │   └── query.sql.go
│   │   └── sql/
│   │       ├── query.sql
│   │       └── schema.sql
│   └── server/
│       ├── routes.go
│       ├── routes_test.go
│       └── server.go
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── sqlc.yaml
```

## SQLC Configuration

The `sqlc.yaml` configuration file is automatically created at the root of the project. It defines how SQLC generates Go code from your SQL schema and query files.

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "internal/database/sql/query.sql"
    schema: "internal/database/sql/schema.sql"
    gen:
      go:
        package: "repository"
        out: "internal/database/repository"
        emit_json_tags: true
        emit_prepared_queries: false
        emit_interface: true
        emit_empty_slices: true
```

## Example SQL Files

### Schema 

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### Queries

```sql
-- name: GetUserByID :one
SELECT id, name, email FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;
```

## Makefile Integration

The Makefile includes entries for installing SQLC and generating code.


```bash
sqlc-install: 
	# Installation logic goes here

build: sqlc-install
	@echo "Building..."
	@sqlc generate
	@go build -o main cmd/api/main.go
```

## Use SQLC in your project

By default, the project includes SQL schema and query files.
To customize them:

1. Modify the files under `internal/database/sql/`.
2. Run `sqlc generate` to regenerate the Go code reflecting your changes.
