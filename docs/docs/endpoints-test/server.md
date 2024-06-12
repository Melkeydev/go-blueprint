## Testing Endpoints with CURL and WebSocat

Testing endpoints is an essential part of ensuring the correctness and functionality of your app. Depending on what options are used for go-blueprint project creation, you have various endpoints for testing your init application status.


Before proceeding, ensure you have the following tools installed:

- [CURL](https://curl.se/docs/manpage.html): A command-line tool for transferring data with URLs.
- [WebSocat](https://github.com/vi/websocat): A command-line WebSocket client.

You can utilize alternative tools that support the WebSocket protocol to establish connections with the server. WebSocat is an open-source CLI tool, while [POSTMAN](https://www.postman.com/) serves as a GUI tool specifically designed for testing APIs and WebSocket functionality.

## Hello World Endpoint

To test the Hello World endpoint, execute the following curl command:

```bash
curl http://localhost:PORT
```

Sample Output:
```json
{"message": "Hello World"}
```
If the server is running and it is healthy, you should see the message 'Hello World' in the response.
Also, depending on the framework you are using, there will be logs in the terminal:

```bash
make run
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> websocket-test/internal/server.(*Server).HelloWorldHandler-fm (3 handlers)
[GIN-debug] GET    /health                   --> websocket-test/internal/server.(*Server).healthHandler-fm (3 handlers)
[GIN-debug] GET    /websocket                --> websocket-test/internal/server.(*Server).websocketHandler-fm (3 handlers)
[GIN] 2024/05/28 - 17:44:31 | 200 |       27.93µs |       127.0.0.1 | GET      "/"
```
