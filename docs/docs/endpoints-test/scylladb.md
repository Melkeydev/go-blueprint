To test the ScyllaDB Health Check endpoint, use the following curl command:

```bash
curl http://localhost:PORT/health
```
## Health Function

The `Health` function checks the health of the ScyllaDB Cluster by pinging the [Coordinator Node](https://opensource.docs.scylladb.com/stable/architecture/architecture-fault-tolerance.html). It returns a simple map containing a health message.

### Functionality

**Ping ScyllaDB Server**: The function pings the ScyllaDB through server to check its availability.

   - If the ping fails, it logs the error and terminates the program.
   - If the ping succeeds, it returns a health message indicating that the server with some .

### Sample Output

The `Health` returns a JSON-like map structure with a single key indicating the health status:

```json
{
  "message": "It's healthy",
  "status": "up",
  "scylla_active_conns": "5",
  "scylla_cluster_size": "1",
  "scylla_current_datacenter": "datacenter1",
  "scylla_current_time": "2024-11-04 22:59:21.69 +0000 UTC",
  "scylla_health_check_duration": "16.896976ms",
  "scylla_keyspaces": "6"
}
```

## Code implementation

```go
func (s *service) checkScyllaHealth(ctx context.Context, stats map[string]string) map[string]string {
    startedAt := time.Now()
    
    // Execute a simple query to check connectivity
    query := "SELECT now() FROM system.local"
    iter := s.Session.Query(query).WithContext(ctx).Iter()
    var currentTime time.Time
    if !iter.Scan(&currentTime) {
        if err := iter.Close(); err != nil {
            stats["status"] = "down"
            stats["message"] = fmt.Sprintf("Failed to execute query: %v", err)
            return stats
        }
    }
    if err := iter.Close(); err != nil {
        stats["status"] = "down"
        stats["message"] = fmt.Sprintf("Error during query execution: %v", err)
        return stats
    }
    
    // ScyllaDB is up
    stats["status"] = "up"
    stats["message"] = "It's healthy"
    stats["scylla_current_time"] = currentTime.String()
    
    // Retrieve cluster information
    // Get keyspace information
    getKeyspacesQuery := "SELECT keyspace_name FROM system_schema.keyspaces"
    keyspacesIterator := s.Session.Query(getKeyspacesQuery).Iter()
    
    stats["scylla_keyspaces"] = strconv.Itoa(keyspacesIterator.NumRows())
    if err := keyspacesIterator.Close(); err != nil {
        log.Fatalf("Failed to close keyspaces iterator: %v", err)
    }
    
    // Get cluster host information
    var currentDatacenter string
    clusterNodesIterator := s.Session.Query("SELECT data_center FROM system.peers").Iter()
    clusterNodesIterator.Scan(&currentDatacenter)
    
    // +1 is because the default connection (coordinator) is not included in the query
    stats["scylla_cluster_size"] = strconv.Itoa(clusterNodesIterator.NumRows() + 1)
    stats["scylla_current_datacenter"] = currentDatacenter
    if err := clusterNodesIterator.Close(); err != nil {
        log.Fatalf("Failed to close cluster nodes iterator: %v", err)
    }
    
    // Retrieve Connected Sessions
    connectedSessionsIterator := s.Session.Query("SELECT connection_stage as connected_sessions FROM system.clients").Iter()
    stats["scylla_active_conns"] = strconv.Itoa(connectedSessionsIterator.NumRows())
    
    if err := connectedSessionsIterator.Close(); err != nil {
        log.Fatalf("Failed to close cluster nodes iterator: %v", err)
    }
    
    // Calculate the time taken to perform the health check
    stats["scylla_health_check_duration"] = time.Since(startedAt).String()
    return stats
}

```

## Note

Scylladb does not support advanced health check functions like SQL databases or Redis. 
Implementation is based on queries at `system` related keyspaces.
