# Worker Background Jobs

The Worker feature enables background job processing in your Go application using the [hibiken/asynq](https://github.com/hibiken/asynq) library. This feature creates a separate worker command that can process tasks asynchronously using Redis as the message broker.

## Overview

When you enable the Worker feature, go-blueprint generates:

- A dedicated worker command (`cmd/worker/main.go`)
- Example task definitions (`cmd/worker/tasks/`)
- Redis configuration in environment variables
- Proper dependency management with `go.mod`

## Usage

### Command Line

```bash
# Enable worker feature along with other options
go-blueprint create --name myapp --framework gin --feature worker

# Worker can be combined with other features
go-blueprint create --name myapp --framework gin --feature worker --feature docker
```

### Interactive Mode

When running `go-blueprint create` interactively, the worker option will appear in the advanced features selection:

```
 What advanced features would you like to enable?
 
 [x] Worker
 Background job processing with Redis and asynq
```

## Generated Structure

The worker feature generates the following files:

```
your-project/
├── cmd/
│   └── worker/
│       ├── main.go                    # Worker server setup
│       └── tasks/
│           └── hello_world_task.go    # Example task implementation
├── .env                               # Includes Redis configuration
└── go.mod                            # Includes asynq dependencies
```

## Configuration

### Environment Variables

The worker feature adds the following environment variables to your `.env` file:

```env
# Worker Configuration
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

### Redis Setup

The worker requires Redis to be running. You can start Redis using:

```bash
# Using Docker
docker run -d -p 6379:6379 redis:alpine

# Using Homebrew on macOS
brew install redis
brew services start redis

# Using apt on Ubuntu/Debian
sudo apt install redis-server
sudo systemctl start redis-server
```

## Worker Command

The generated worker command (`cmd/worker/main.go`) includes:

- **Redis Connection**: Configured using environment variables
- **Task Multiplexer**: Uses `asynq.NewServeMux()` for routing tasks
- **Signal Handling**: Graceful shutdown on SIGINT/SIGTERM
- **Task Registration**: Automatic registration of task handlers

### Running the Worker

```bash
# Start the worker server
go run cmd/worker/main.go

# Or build and run
go build -o worker cmd/worker/main.go
./worker
```

## Example Task

The generated `hello_world_task.go` demonstrates:

```go
package tasks

import (
    "context"
    "encoding/json"
    "log"
    "github.com/hibiken/asynq"
)

const TypeHelloWorld = "hello_world"

type HelloWorldPayload struct {
    Name string `json:"name"`
}

func HandleHelloWorldTask(ctx context.Context, t *asynq.Task) error {
    var payload HelloWorldPayload
    if err := json.Unmarshal(t.Payload(), &payload); err != nil {
        return err
    }
    
    log.Printf("Hello, %s!", payload.Name)
    return nil
}
```

## Creating Tasks

To create and enqueue tasks from your main application:

```go
package main

import (
    "encoding/json"
    "log"
    "github.com/hibiken/asynq"
)

func main() {
    client := asynq.NewClient(asynq.RedisClientOpt{
        Addr: "localhost:6379",
    })
    defer client.Close()

    // Create a task
    payload := map[string]string{"name": "World"}
    data, _ := json.Marshal(payload)
    task := asynq.NewTask("hello_world", data)

    // Enqueue the task
    info, err := client.Enqueue(task)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Enqueued task: %+v", info)
}
```

## Advanced Usage

### Task Options

You can enqueue tasks with various options:

```go
// Delay task execution
task := asynq.NewTask("hello_world", payload)
client.Enqueue(task, asynq.ProcessIn(5*time.Minute))

// Set retry attempts
client.Enqueue(task, asynq.MaxRetry(3))

// Combine options
client.Enqueue(task, 
    asynq.ProcessIn(10*time.Second),
    asynq.MaxRetry(5),
    asynq.Queue("high"),
)
```

### Multiple Queues

```go
// In worker setup
mux := asynq.NewServeMux()
mux.HandleFunc("hello_world", HandleHelloWorldTask)

srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: "localhost:6379"},
    asynq.Config{
        Queues: map[string]int{
            "critical": 6, // Higher priority
            "default":  3,
            "low":      1,
        },
    },
)
```

## Integration with Other Features

The worker feature integrates seamlessly with other go-blueprint features:

- **Docker**: Redis service is automatically added to `docker-compose.yml`
- **Environment**: Worker variables are included in `.env` templates
- **Testing**: Worker tasks can be unit tested using asynq's testing utilities

## Dependencies

The worker feature automatically adds these dependencies to your `go.mod`:

- `github.com/hibiken/asynq`: Core async task queue
- `github.com/joho/godotenv`: Environment variable loading

## Monitoring

Asynq provides built-in monitoring capabilities:

```go
// Enable monitoring server
srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: "localhost:6379"},
    asynq.Config{
        // ... other config
        MonitoringAddr: ":8080", // Enable monitoring on port 8080
    },
)
```

Access the monitoring dashboard at `http://localhost:8080`

## Best Practices

1. **Error Handling**: Always handle errors appropriately in task handlers
2. **Timeouts**: Set appropriate timeouts for long-running tasks
3. **Retries**: Configure retry logic for transient failures
4. **Monitoring**: Use asynq's monitoring features in production
5. **Testing**: Write unit tests for your task handlers
6. **Security**: Secure your Redis instance in production environments

## Troubleshooting

### Common Issues

1. **Redis Connection Failed**
   - Ensure Redis is running and accessible
   - Check REDIS_ADDR environment variable
   - Verify network connectivity

2. **Tasks Not Processing**
   - Check worker logs for errors
   - Verify task handlers are registered
   - Ensure correct queue configuration

3. **Performance Issues**
   - Monitor Redis memory usage
   - Adjust concurrency settings
   - Use appropriate queue priorities

### Debugging

Enable debug logging in your worker:

```go
srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: "localhost:6379"},
    asynq.Config{
        Logger: asynq.DefaultLogger,
        LogLevel: asynq.DebugLevel,
    },
)
```

## Resources

- [Asynq Documentation](https://pkg.go.dev/github.com/hibiken/asynq)
- [Redis Documentation](https://redis.io/documentation)
- [Go-Blueprint Advanced Features](../advanced-flag.md)
