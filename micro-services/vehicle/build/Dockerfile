# Stage 1: Build the application
FROM golang:1.23 AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /vehicle cmd/main.go

# Stage 2: Runtime container
FROM debian:bullseye-slim

# Set working directory in the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /vehicle /app/vehicle

# Expose the application port
EXPOSE 8088

# Command to run the application
CMD ["./vehicle"]
