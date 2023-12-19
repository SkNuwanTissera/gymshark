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

RUN chmod +x /app/main  # Set executable permissions

EXPOSE 8080
CMD ["/app/main"]  # Update if main isnt the correct entry point
