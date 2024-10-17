To test the SQL DB Health Check endpoint, use the following curl command:

```bash
curl http://localhost:PORT/health
```
## Health Function

The `Health` function checks the health of the database connection by pinging the database and retrieving various statistics. It returns a map with keys indicating different health metrics.

### Functionality

**Ping the Database**: The function pings the database to ensure it is reachable.

   - If the database is down, it logs the error, sets the status to "down," and terminates the program.
   - If the database is up, it proceeds to gather additional statistics.

**Collect Database Statistics**: The function retrieves the following statistics from the database connection:

   - `open_connections`: Number of open connections to the database.
   - `in_use`: Number of connections currently in use.
   - `idle`: Number of idle connections.
   - `wait_count`: Number of times a connection has to wait.
   - `wait_duration`: Total time connections have spent waiting.
   - `max_idle_closed`: Number of connections closed due to exceeding idle time.
   - `max_lifetime_closed`: Number of connections closed due to exceeding their lifetime.

**Evaluate Statistics**: Evaluates the collected statistics to provide a health message. Based on predefined thresholds, it updates the health message to indicate potential issues, such as heavy load or high wait events.

### Sample Output

The `Health` function returns a JSON-like map structure with the following keys and example values:

```json
{
  "idle": "1",
  "in_use": "0",
  "max_idle_closed": "0",
  "max_lifetime_closed": "0",
  "message": "It's healthy",
  "open_connections": "1",
  "status": "up",
  "wait_count": "0",
  "wait_duration": "0s"
}
```

## Code Implementation

```go
func (s *service) Health() map[string]string {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    stats := make(map[string]string)

    err := s.db.PingContext(ctx)
    if err != nil {
        stats["status"] = "down"
        stats["error"] = fmt.Sprintf("db down: %v", err)
        log.Fatalf("db down: %v", err)  
        return stats
    }

    stats["status"] = "up"
    stats["message"] = "It's healthy"

    dbStats := s.db.Stats()
    stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
    stats["in_use"] = strconv.Itoa(dbStats.InUse)
    stats["idle"] = strconv.Itoa(dbStats.Idle)
    stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
    stats["wait_duration"] = dbStats.WaitDuration.String()
    stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
    stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

    if dbStats.OpenConnections > 40 { 
        stats["message"] = "The database is experiencing heavy load."
    }

    if dbStats.WaitCount > 1000 {
        stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
    }

    if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
        stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
    }

    if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
        stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
    }

    return stats
}
```
