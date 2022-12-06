# Build stage
FROM golang:1.19.3-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go && apk add curl && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start-db.sh .
COPY wait-for-db.sh .
COPY ./db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start-db.sh" ]