To test the Redis Health Check endpoint, use the following curl command:

```bash
curl http://localhost:PORT/health
```

## Health Function

The `Health` function orchestrates the health assessment of the Redis server by invoking the `checkRedisHealth` function and returning the collected statistics.

### Functionality

**Check Redis Health**: The function pings the Redis server to check its availability and adds the response to the stats map.

   - If the ping fails, it logs the error and terminates the program.
   - If the ping succeeds, it proceeds to retrieve additional information.

**Retrieve Redis Information**: The function retrieves information about the Redis server, including version, mode, connected clients, memory usage, uptime, etc.

   - If an error occurs during info retrieval, it updates the health message accordingly.

**Evaluate Redis Statistics**: The function evaluates the collected statistics to identify potential issues and updates the health message accordingly.

   - It checks for high number of connected clients, stale connections, memory usage, recent restart, high idle connections, and high connection pool utilization.

### Sample Output

The `Health` function returns a JSON-like map structure with various keys representing different health metrics and their corresponding values.

```json
{
  "redis_active_connections": "0",
  "redis_connected_clients": "1",
  "redis_hits_connections": "1",
  "redis_idle_connections": "1",
  "redis_max_memory": "0",
  "redis_message": "Redis has been recently restarted",
  "redis_misses_connections": "1",
  "redis_mode": "standalone",
  "redis_ping_response": "PONG",
  "redis_pool_size_percentage": "0.42%",
  "redis_stale_connections": "0",
  "redis_status": "up",
  "redis_timeouts_connections": "0",
  "redis_total_connections": "1",
  "redis_uptime_in_seconds": "55",
  "redis_used_memory": "980704",
  "redis_used_memory_peak": "980704",
  "redis_version": "7.2.4"
}
```

### Code Implementation

```go
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Default is now 5s
	defer cancel()

	stats := make(map[string]string)

	stats = s.checkRedisHealth(ctx, stats)

	return stats
}

func (s *service) checkRedisHealth(ctx context.Context, stats map[string]string) map[string]string {
	pong, err := s.db.Ping(ctx).Result()
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	stats["redis_status"] = "up"
	stats["redis_message"] = "It's healthy"
	stats["redis_ping_response"] = pong

	info, err := s.db.Info(ctx).Result()
	if err != nil {
		stats["redis_message"] = fmt.Sprintf("Failed to retrieve Redis info: %v", err)
		return stats
	}

	redisInfo := parseRedisInfo(info)

	poolStats := s.db.PoolStats()

	stats["redis_version"] = redisInfo["redis_version"]
	stats["redis_mode"] = redisInfo["redis_mode"]
	stats["redis_connected_clients"] = redisInfo["connected_clients"]
	stats["redis_used_memory"] = redisInfo["used_memory"]
	stats["redis_used_memory_peak"] = redisInfo["used_memory_peak"]
	stats["redis_uptime_in_seconds"] = redisInfo["uptime_in_seconds"]
	stats["redis_hits_connections"] = strconv.FormatUint(uint64(poolStats.Hits), 10)
	stats["redis_misses_connections"] = strconv.FormatUint(uint64(poolStats.Misses), 10)
	stats["redis_timeouts_connections"] = strconv.FormatUint(uint64(poolStats.Timeouts), 10)
	stats["redis_total_connections"] = strconv.FormatUint(uint64(poolStats.TotalConns), 10)
	stats["redis_idle_connections"] = strconv.FormatUint(uint64(poolStats.IdleConns), 10)
	stats["redis_stale_connections"] = strconv.FormatUint(uint64(poolStats.StaleConns), 10)
	stats["redis_max_memory"] = redisInfo["maxmemory"]

	activeConns := uint64(math.Max(float64(poolStats.TotalConns-poolStats.IdleConns), 0))
	stats["redis_active_connections"] = strconv.FormatUint(activeConns, 10)

	poolSize := s.db.Options().PoolSize
	connectedClients, _ := strconv.Atoi(redisInfo["connected_clients"])
	poolSizePercentage := float64(connectedClients) / float64(poolSize) * 100
	stats["redis_pool_size_percentage"] = fmt.Sprintf("%.2f%%", poolSizePercentage)

	return s.evaluateRedisStats(redisInfo, stats)
}

func (s *service) evaluateRedisStats(redisInfo, stats map[string]string) map[string]string {
	poolSize := s.db.Options().PoolSize
	poolStats := s.db.PoolStats()
	connectedClients, _ := strconv.Atoi(redisInfo["connected_clients"])
	highConnectionThreshold := int(float64(poolSize) * 0.8)

	if connectedClients > highConnectionThreshold {
		stats["redis_message"] = "Redis has a high number of connected clients"
	}

	minStaleConnectionsThreshold := 500
	if int(poolStats.StaleConns) > minStaleConnectionsThreshold {
		stats["redis_message"] = fmt.Sprintf("Redis has %d stale connections.", poolStats.StaleConns)
	}

	usedMemory, _ := strconv.ParseInt(redisInfo["used_memory"], 10, 64)
	maxMemory, _ := strconv.ParseInt(redisInfo["maxmemory"], 10, 64)
	if maxMemory > 0 {
		usedMemoryPercentage := float64(usedMemory) / float64(maxMemory) * 100
		if usedMemoryPercentage >= 90 {
			stats["redis_message"] = "Redis is using a significant amount of memory"
		}
	}

	uptimeInSeconds, _ := strconv.ParseInt(redisInfo["uptime_in_seconds"], 10, 64)
	if uptimeInSeconds < 3600 {
		stats["redis_message"] = "Redis has been recently restarted"
	}

	idleConns := int(poolStats.IdleConns)
	highIdleConnectionThreshold := int(float64(poolSize) * 0.7)
	if idleConns > highIdleConnectionThreshold {
		stats["redis_message"] = "Redis has a high number of idle connections"
	}

	poolUtilization := float64(poolStats.TotalConns-poolStats.IdleConns) / float64(poolSize) * 100
	highPoolUtilizationThreshold := 90.0
	if poolUtilization > highPoolUtilizationThreshold {
		stats["redis_message"] = "Redis connection pool utilization is high"
	}

	return stats
}
```
