# Build stage
FROM golang:1.22.11-alpine AS builder

# Set working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Runtime stage
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port (Render will override this)
EXPOSE 8080

# Start application
CMD ["./main"]
