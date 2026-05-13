# Swagger

The Swagger advanced feature adds Swaggo-generated API documentation and serves Swagger UI from the generated application.

```bash
go-blueprint create --advanced --feature swagger
```

After starting the generated API, open:

```bash
http://localhost:8080/swagger/index.html
```

To regenerate the docs after changing route comments:

```bash
make swagger
```
