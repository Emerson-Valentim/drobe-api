# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git and build dependencies
RUN apk add --no-cache git

# Install air for hot reloading (only for development)
RUN go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main cmd/api/main.go

# Final stage
FROM builder AS runner

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /go/bin/air /bin/air

# Copy air configuration (we'll create this next)
COPY .air.toml .

# Expose port
EXPOSE 8080
