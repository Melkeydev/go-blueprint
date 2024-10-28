To test the MongoDB Health Check endpoint, use the following curl command:

```bash
curl http://localhost:PORT/health
```
## Health Function

The `Health` function checks the health of the MongoDB by pinging it. It returns a simple map containing a health message.

### Functionality

**Ping MongoDB Server**: The function pings the MongoDB thru server to check its availability.

   - If the ping fails, it logs the error and terminates the program.
   - If the ping succeeds, it returns a health message indicating that the server is healthy.

### Sample Output

The `Health` returns a JSON-like map structure with a single key indicating the health status:

```json
{
  "message": "It's healthy"
}
```

## Code implementation

```go
func (s *service) Health() map[string]string {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    err := s.db.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("db down: %v", err) 
    }

    return map[string]string{
        "message": "It's healthy",
    }
}
```

## Note

MongoDB does not support advanced health check functions like SQL databases or Redis. Implementation is basic, providing only a simple ping response to indicate if the server is reachable and DB connection healthy.
