# ------------------------ Base Stage ------------------------
    FROM golang:1.23 AS base

    ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        GOPROXY=https://goproxy.cn,direct
    
    WORKDIR /gholi-fly
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    RUN go install github.com/air-verse/air@v1.61.5
    # ------------------------ Development Stage ------------------------
    FROM base AS dev
    
    COPY . .
    
    CMD ["air", "-c", ".air.toml"]
    
    # ------------------------ Builder Stage ------------------------
    FROM base AS builder
    
    COPY . .
    RUN go build -o ./bank-api ./cmd/main.go
    
    # ------------------------ Production Stage ------------------------
    FROM alpine:latest AS prod
    
    RUN apk add --no-cache ca-certificates tzdata curl
    ENV TZ=Asia/Tehran
    
    WORKDIR /gholi-fly
    
    COPY --from=builder /gholi-fly/bank-api ./bank-api
    CMD ["./bank-api", "--config", "/etc/config.json"]
    