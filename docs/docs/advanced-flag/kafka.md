# Kafka Consumer

The Kafka consumer feature adds a complete Kafka consumer implementation to your Go Blueprint project using the `github.com/segmentio/kafka-go` library. This feature creates a structured, production-ready Kafka consumer that can be easily integrated into your application.

## Features

- **Consumer Implementation**: Complete Kafka consumer with configurable settings
- **Handler-Based Architecture**: Flexible message processing through handler functions
- **Environment Configuration**: Configurable via environment variables
- **Error Handling**: Robust error handling and logging
- **Testing**: Comprehensive unit tests included
- **Graceful Shutdown**: Supports context-based cancellation
- **Standalone Binary**: Dedicated consumer binary for independent execution

## Generated Files

When you select the Kafka consumer feature, the following files are generated:

### `pkg/kafka/segmentio/consumer.go`
- Core consumer implementation using `github.com/segmentio/kafka-go`
- `MessageHandler` function type for custom business logic
- Context-based cancellation support
- Built-in error handling and logging

### `pkg/kafka/segmentio/consumer_test.go`
- Unit tests for the consumer
- Tests for both default and custom message handlers
- Integration test template

### `cmd/consumer/main.go`
- **Standalone consumer binary** that can be compiled and run independently
- Command-line flags for configuration
- Environment variable support
- Graceful shutdown handling

### Environment Variables (added to `.env`)
```env
# Kafka Configuration
KAFKA_BROKERS=localhost:9092
KAFKA_GROUP_ID=<your-project-name>-consumer-group
KAFKA_TOPIC=<your-project-name>-topic
```

## Usage

### Command Line

Generate a project with Kafka consumer support:

```bash
go-blueprint create --name my-project --framework gin --feature kafka
```

### Standalone Consumer Binary (Recommended)

Build and run the consumer as a standalone binary:

```bash
# Build the consumer binary
go build -o consumer ./cmd/consumer

# Run with environment variables from .env file
./consumer

# Run with command-line flags
./consumer -brokers localhost:9092 -group-id my-group -topic my-topic

# Run directly with Go
go run ./cmd/consumer
```

The consumer binary supports the following command-line flags:
- `-brokers`: Kafka broker addresses (comma-separated)
- `-group-id`: Consumer group ID
- `-topic`: Kafka topic to consume from

### Integration with Your Application

Import and use the consumer in your Go application:

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "your-project-name/pkg/kafka/segmentio"
    "github.com/segmentio/kafka-go"
)

func main() {
    // Create consumer
    consumer := segmentio.NewConsumer(
        []string{"localhost:9092"},
        "my-group",
        "my-topic",
    )
    
    // Create custom message handler
    handler := func(msg kafka.Message) error {
        log.Printf("Processing message: %s", string(msg.Value))
        
        // Add your business logic here
        // For example: parse JSON, save to database, etc.
        
        return nil
    }
    
    // Create context with cancellation
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Handle graceful shutdown
    go func() {
        sigChan := make(chan os.Signal, 1)
        signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
        <-sigChan
        log.Println("Shutting down consumer...")
        cancel()
    }()
    
    // Start consuming
    log.Println("Starting consumer...")
    if err := consumer.Consume(ctx, handler); err != nil {
        log.Printf("Consumer error: %v", err)
    }
}
```

### Custom Message Processing

The consumer uses a handler-based architecture for flexible message processing:

```go
// Define your custom message handler
customHandler := func(msg kafka.Message) error {
    // Parse the message
    var data MyDataStruct
    if err := json.Unmarshal(msg.Value, &data); err != nil {
        return fmt.Errorf("failed to parse message: %w", err)
    }
    
    // Process the data
    if err := processBusinessLogic(data); err != nil {
        return fmt.Errorf("failed to process data: %w", err)
    }
    
    log.Printf("Successfully processed message: %s", msg.Key)
    return nil
}

// Use the handler with the consumer
consumer.Consume(ctx, customHandler)
```

### Configuration

Configure the consumer through environment variables in your `.env` file:

```env
# Kafka Configuration
KAFKA_BROKERS=localhost:9092,localhost:9093
KAFKA_GROUP_ID=my-app-consumer-group
KAFKA_TOPIC=my-app-events
```

Or through command-line flags:

```bash
./consumer -brokers localhost:9092 -group-id my-group -topic my-topic
```

### Testing

Run the tests:

```bash
# Run all tests
go test ./pkg/kafka/segmentio/...

# Run tests with verbose output
go test -v ./pkg/kafka/segmentio/...

# Run integration tests (requires running Kafka instance)
go test -v ./pkg/kafka/segmentio/... -short=false
```

## Development Setup

For local development, you can run Kafka using Docker:

```bash
# Create docker-compose.yml
cat > docker-compose.yml << EOF
version: '3.8'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
EOF

# Start Kafka and Zookeeper
docker-compose up -d

# Create a topic
docker exec -it <kafka-container> kafka-topics --create --topic your-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1

# Produce test messages
docker exec -it <kafka-container> kafka-console-producer --topic your-topic --bootstrap-server localhost:9092

# Consumer messages (optional for testing)
docker exec -it <kafka-container> kafka-console-consumer --topic your-topic --from-beginning --bootstrap-server localhost:9092
```

## API Reference

### Consumer

#### `NewConsumer(brokers []string, groupID, topic string) *Consumer`
Creates a new Kafka consumer instance.

- `brokers`: List of Kafka broker addresses
- `groupID`: Consumer group ID
- `topic`: Kafka topic to consume from

#### `Consume(ctx context.Context, handler MessageHandler) error`
Starts consuming messages from Kafka using the provided handler.

- `ctx`: Context for cancellation
- `handler`: Function to process each message

#### `Close() error`
Closes the consumer and releases resources.

### MessageHandler

```go
type MessageHandler func(msg kafka.Message) error
```

Function type for processing Kafka messages. Return an error to indicate processing failure.

### DefaultMessageHandler

```go
func DefaultMessageHandler(msg kafka.Message) error
```

A default message handler that logs the message content.

## Production Considerations

1. **Error Handling**: Implement proper error handling in your message handlers
2. **Monitoring**: Add metrics and monitoring for consumer lag and processing rates
3. **Scaling**: Run multiple consumer instances for horizontal scaling
4. **Security**: Configure SSL/TLS and authentication for production environments
5. **Logging**: Enhance logging with structured logging libraries
6. **Dead Letter Queues**: Implement dead letter queues for failed messages
7. **Backpressure**: Handle backpressure when processing is slower than consumption

## Resources

- [Kafka-go Documentation](https://github.com/segmentio/kafka-go)
- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
- [Go Blueprint Documentation](https://github.com/melkeydev/go-blueprint)
