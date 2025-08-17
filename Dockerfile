# Stage 1: Build Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod & go.sum first, install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy source code
COPY . .

# Build binary
RUN go build -o main .

# Stage 2: Minimal image
FROM alpine:latest

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run server
CMD ["./main"]