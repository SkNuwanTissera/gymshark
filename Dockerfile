# Build stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/start.sh .

EXPOSE 8080
CMD [ "/app/main" ]
