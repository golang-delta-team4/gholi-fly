# Build stage
FROM golang:1.23 AS builder

# Environment variables for Go build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory for the build
WORKDIR /gholi-fly

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o ./hotel-api ./cmd/main.go

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk add --no-cache ca-certificates tzdata

# Set timezone
ENV TZ=Asia/Tehran

# Set the working directory for the application
WORKDIR /gholi-fly

# Copy the binary from the builder stage
COPY --from=builder /gholi-fly/hotel-api ./hotel-api

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./hotel-api", "--config", "/etc/config.json"]
