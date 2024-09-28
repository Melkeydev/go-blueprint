The Docker advanced flag provides the app's Dockerfile configuration and creates or updates the docker-compose.yml file, which is generated if a DB driver is used.
The Dockerfile includes a two-stage build that leverages Makefile configuration. In the end, you will have a smaller image without unnecessary build dependencies.

## Dockerfile

```dockerfile
FROM golang:1.23-alpine AS build

RUN apk add --no-cache make curl

WORKDIR /app

COPY . .
RUN go mod download

RUN make build

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
```
## Docker compose
Docker and docker-compose.yml pull environment variables from the .env file.

Example if the Docker flag is used with the MySQL DB driver:
```ymal
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
      target: prod
    restart: unless-stopped
    ports:
      - ${PORT}:${PORT}
    environment:
      APP_ENV: ${APP_ENV}
      PORT: ${PORT}
      BLUEPRINT_DB_HOST: ${BLUEPRINT_DB_HOST}
      BLUEPRINT_DB_PORT: ${BLUEPRINT_DB_PORT}
      BLUEPRINT_DB_DATABASE: ${BLUEPRINT_DB_DATABASE}
      BLUEPRINT_DB_USERNAME: ${BLUEPRINT_DB_USERNAME}
      BLUEPRINT_DB_PASSWORD: ${BLUEPRINT_DB_PASSWORD}
    depends_on:
      mysql_bp:
        condition: service_healthy
    networks:
      - blueprint
  mysql_bp:
    image: mysql:latest
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: ${BLUEPRINT_DB_DATABASE}
      MYSQL_USER: ${BLUEPRINT_DB_USERNAME}
      MYSQL_PASSWORD: ${BLUEPRINT_DB_PASSWORD}
      MYSQL_ROOT_PASSWORD: ${BLUEPRINT_DB_ROOT_PASSWORD}
    ports:
      - "${BLUEPRINT_DB_PORT}:3306"
    volumes:
      - mysql_volume_bp:/var/lib/mysql
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "${BLUEPRINT_DB_HOST}", "-u", "${BLUEPRINT_DB_USERNAME}", "--password=${BLUEPRINT_DB_PASSWORD}"]
      interval: 5s
      timeout: 5s
      retries: 3
      start_period: 15s
    networks:
      - blueprint

volumes:
  mysql_volume_bp:
networks:
  blueprint:
```

## Note
If you are testing more than one framework locally, be aware of Docker leftovers such as volumes.
For proper cleaning and building, use `docker compose down --volumes` and `docker compose up --build`.

