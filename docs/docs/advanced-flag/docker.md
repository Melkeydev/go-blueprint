The Docker advanced flag provides the app's Dockerfile configuration and creates or updates the docker-compose.yml file, which is generated if a DB driver is used.
The Dockerfile includes a two-stage build, and the final config depends on the use of advanced features. In the end, you will have a smaller image without unnecessary build dependencies.

## Dockerfile

```dockerfile
FROM golang:1.23-alpine AS build

RUN apk add --no-cache curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate && \
    curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss && \
    chmod +x tailwindcss && \
    ./tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
```

Docker config if React flag is used:

```dockerfile
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]


FROM node:20 AS frontend_builder
WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install
COPY frontend/. .
RUN npm run build

FROM node:20-slim AS frontend
RUN npm install -g serve
COPY --from=frontend_builder /frontend/dist /app/dist
EXPOSE 5173
CMD ["serve", "-s", "/app/dist", "-l", "5173"]
```
## Docker compose
Docker and docker-compose.yml pull environment variables from the .env file.

Example if the Docker flag is used with the MySQL DB driver:
```yaml
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

or

```bash
docker compose build --no-cache && docker compose up
```
