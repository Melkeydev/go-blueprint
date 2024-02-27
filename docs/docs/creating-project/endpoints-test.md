## Testing Endpoints with CURL and WebSocat

Testing endpoints is an essential part of ensuring the correctness and functionality of your app. Depending on what options are used for go-blueprint project creation, you have various endpoints for testing your init application status.


Before proceeding, ensure you have the following tools installed:

- [CURL](https://curl.se/docs/manpage.html): A command-line tool for transferring data with URLs.
- [WebSocat](https://github.com/vi/websocat): A command-line WebSocket client.



## Hello World Endpoint

To test the Hello World endpoint, execute the following curl command:

```bash
curl http://localhost:PORT
```

Expected Output:
```json
{"message": "Hello World"}
```

## DB Health Check Endpoint

To test the DB Health Check endpoint, use the following curl command:

```bash
curl http://localhost:PORT/health
```

Expected Output:
```json
{"message": "It's healthy"}
```

## WebSocket Endpoint

Initiate a WebSocket connection:

```bash
websocat ws://localhost:PORT/websocket
```

Expected Output:
```
server timestamp: 1709046650354893857
server timestamp: 1709046652355956336
server timestamp: 1709046654357101642
server timestamp: 1709046656357202535
```

## Testing /web Endpoint

To test the `/web` endpoint, you can simply open it in a web browser. This endpoint serves a simple HTML page with a form.
Navigate to `http://localhost:PORT/web`
This page contains a form with a single input field and a submit button. Upon submitting the form, "Hello, [input]" will be displayed.


