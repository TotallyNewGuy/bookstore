# Build stage
FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY app.env .
COPY wait-for.sh .
RUN ["chmod", "+x", "/app/wait-for.sh"]
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]