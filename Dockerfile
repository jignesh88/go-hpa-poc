FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY main.go .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a small image for the final container
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./app"]