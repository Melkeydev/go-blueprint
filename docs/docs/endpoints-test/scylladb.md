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
  "scylla_cluster_nodes_up": "3",
  "scylla_cluster_nodes_down": "0",
  "scylla_cluster_size": "1",
  "scylla_current_datacenter": "datacenter1",
  "scylla_current_time": "2024-11-04 22:59:21.69 +0000 UTC",
  "scylla_health_check_duration": "16.896976ms",
  "scylla_keyspaces": "6"
}
```

## Code implementation

```go
func (s *service) Health() map[string]string {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    stats := make(map[string]string)
    
    // Check ScyllaDB health and populate the stats map
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
    
    // Get cluster information
    var currentDatacenter string
    var currentHostStatus bool
    
    var clusterNodesUp uint
    var clusterNodesDown uint
    var clusterSize uint
    
    clusterNodesIterator := s.Session.Query("SELECT dc, up FROM system.cluster_status").Iter()
    for clusterNodesIterator.Scan(&currentDatacenter, &currentHostStatus) {
        clusterSize++
        if currentHostStatus {
            clusterNodesUp++
        } else {
            clusterNodesDown++
        }
    }
    
    if err := clusterNodesIterator.Close(); err != nil {
        log.Fatalf("Failed to close cluster nodes iterator: %v", err)
    }
    
    stats["scylla_cluster_size"] = strconv.Itoa(int(clusterSize))
    stats["scylla_cluster_nodes_up"] = strconv.Itoa(int(clusterNodesUp))
    stats["scylla_cluster_nodes_down"] = strconv.Itoa(int(clusterNodesDown))
    stats["scylla_current_datacenter"] = currentDatacenter
    
    // Calculate the time taken to perform the health check
    stats["scylla_health_check_duration"] = time.Since(startedAt).String()
    return stats
}

```

## Note

Scylladb does not support advanced health check functions like SQL databases or Redis. 
The current implementation is based on queries at `system` related keyspaces.
