# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git and build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main cmd/migrations/main.go

# Final stage
FROM builder AS runner

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .
