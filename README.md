# Drobe API

A Go-based API service for managing clothes and wardrobe.

## Project Structure

```
.
├── cmd/           # Application entry points
├── services/      # Business logic and services
├── docs/          # Documentation
└── item.go     # Core types and models
```

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Run `go run cmd/api/main.go` to start the server

## Development

This project follows a service-oriented architecture with:
- Services: Business logic layer
- Types: Core domain models
- Documentation: API and project documentation

## Commands

### Lint

`golangci-lint run ./... --fix`

### Generate models

`sqlc generate`

## Required tools

- github.com/golangci/golangci-lint
- github.com/golang-migrate/migrate
- github.com/sqlc-dev/sqlc