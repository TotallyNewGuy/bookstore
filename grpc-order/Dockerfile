# Build stage
FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./
COPY app.env .
COPY start.sh .
RUN ["chmod", "+x", "/app/start.sh"]
COPY wait-for.sh .
RUN ["chmod", "+x", "/app/wait-for.sh"]
COPY db/migration ./migration
RUN ["echo", "开始运行order里面的dockerfile"]

EXPOSE 9092
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]