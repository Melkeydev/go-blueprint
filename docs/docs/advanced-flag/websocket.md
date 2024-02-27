## Websocket

A `/websocket` endpoint is added in `routes.go` to facilitate websocket connections. Upon accessing this endpoint, the server establishes a websocket connection and begins transmitting timestamp messages at 2-second intervals. WS is utilized across all Go-blueprint supported frameworks. This simple implementation showcases how flexible project is.

## Testing with WebSocat
[WebSocat](https://github.com/vi/websocat) is a versatile tool for working with websockets from the command line. Below are some examples of using WebSocat to test the websocket endpoint:

```bash
# Start server
make run
``` 

```bash
# Connect to the websocket endpoint
$ websocat ws://localhost:PORT/websocket
```

Replace `PORT` with the port number on which your server is running.

## Expected Output
Upon successful connection, the client should start receiving timestamp messages from the server every 2 seconds.

```bash
server timestamp: 1709046650354893857
server timestamp: 1709046652355956336
server timestamp: 1709046654357101642
server timestamp: 1709046656357202535
server timestamp: 1709046658358258120
server timestamp: 1709046660359338389
server timestamp: 1709046662360422533
server timestamp: 1709046664361194735
server timestamp: 1709046666362308678
server timestamp: 1709046668363390475
server timestamp: 1709046670364477838
server timestamp: 1709046672365193667
server timestamp: 1709046674366265199
server timestamp: 1709046676366564490
server timestamp: 1709046678367646090
server timestamp: 1709046680367851980
server timestamp: 1709046682368920527
```

You can utilize alternative tools that support the WebSocket protocol to establish connections with the server. WebSocat is an open-source CLI tool, while [POSTMAN](https://www.postman.com/) serves as a GUI tool specifically designed for testing APIs and WebSocket functionality.

