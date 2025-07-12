# Safehouse Backend

REST API backend for a personal portfolio application built with Go.

## Table of Contents

- [Overview](#overview)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
- [Development](#development)
- [Testing](#testing)
- [Deployment](#deployment)
- [Security Features](#security-features)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)

## Overview

Safehouse Backend is a Go application that serves as the backend for a personal portfolio website. It features:

- **JWT-based Authentication** with configurable expiration -> the project does not use credentials at the moment but that might change in the future
- **Multi-layered Security** including rate limiting, input validation, and security headers
- **PostgreSQL Database** with GORM for type-safe queries
- **Dual Secret Management** (Google Cloud Secret Manager for production, local files/env for development)
- **Cloud Native** with Docker support and Google Cloud Run deployment
- **Testing** 

## Technologies Used

### **Core Technologies**
- **Go 1.24** - Programming language
- **Gin Framework** - HTTP web framework
- **PostgreSQL** - Primary database
- **GORM** - ORM for database operations
- **JWT** - Authentication tokens

### **Security & Middleware**
- **Rate Limiting** - IP-based request throttling
- **Input Validation** - XSS, SQL injection, and attack prevention
- **Security Headers** - CSP, CORS, X-Frame-Options, etc.
- **Google Cloud Secret Manager** - Secure secret storage

### **Infrastructure & DevOps**
- **Docker** - Containerization
- **Google Cloud Run** - Serverless deployment
- **Google Container Registry** - Image storage
- **GitHub Actions** - CI/CD pipeline
- **Trivy** - Container vulnerability scanning

### **Development Tools**
- **Swagger/OpenAPI** - API documentation
- **Testify** - Testing framework

## API Endpoints

### **Authentication**
```
POST /auth/token
```
- **Description**: Generate JWT token for frontend authentication (not very useful without proper credentials, but easier to change in the future)
- **Body**: `{"auth_key": "your-auth-key"}`
- **Response**: `{"token": "jwt-token", "expires_in": 1800}`

### **About & Contact**
```
GET /about/contact
```
- **Description**: Get contact information
- **Authentication**: Required (JWT)
- **Response**: Contact details from database

```
GET /about/personal-reviews
```
- **Description**: Get personal reviews carousel data
- **Authentication**: Required (JWT)
- **Response**: Array of personal reviews

### **Portfolio Projects**
```
GET /tech/projects
```
- **Description**: Get technology projects
- **Authentication**: Required (JWT)
- **Response**: Array of tech project groups

```
GET /finance/projects
```
- **Description**: Get finance projects
- **Authentication**: Required (JWT)
- **Response**: Array of finance project groups

```
GET /games/projects
```
- **Description**: Get game projects
- **Authentication**: Required (JWT)
- **Response**: Array of game project groups

```
GET /games/played
```
- **Description**: Get recently played games
- **Authentication**: Required (JWT)
- **Response**: Array of recently played games

### **System**
```
GET /health
```
- **Description**: Health check endpoint
- **Authentication**: Not required
- **Response**: `{"status": "healthy"}`

### **Documentation**
```
GET /
GET /swagger/*
```
- **Description**: Swagger API documentation
- **Authentication**: Not required (development only)

## Getting Started

### **Prerequisites**
- Go 1.24 or higher
- PostgreSQL 12+ database
- Docker
- Google Cloud SDK (for production deployment)

## Development

### **Development Environment Setup**

For development environment please check https://github.com/David-Xilo/safehouse-orchestration

## Testing

### **Essential Go Commands**

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests for specific package
go test ./src/internal/config

# Run specific test
go test -run TestLoadConfig ./src/internal/config

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Clean test cache
go clean -testcache

# Build the application
go build -o main ./src/cmd/api

# Format code
go fmt ./...

# Lint code
go vet ./...

# Security scan
gosec ./...

# Check for vulnerabilities
go list -json -m all | nancy sleuth

# Update dependencies
go mod tidy
go mod verify
```

## Deployment

### **Docker Deployment**

```bash
# Build locally
docker build -t safehouse-backend .

# Run container
docker run -p 8080:8080 \
  -e ENV=development \
  -e GCP_PROJECT_ID=your-project \
  safehouse-backend
```

### **Local Development**

```bash
# Set development environment
export ENV=development

# Run with development secrets
go run src/cmd/api/main.go
```

## Security Features

### **Multi-Layer Security Architecture**

- **JWT Authentication** with HMAC-SHA256 signing
- **Rate Limiting** (5 requests/second, 30 burst)
- **Input Validation** (XSS, SQL injection, path traversal protection)
- **Security Headers** (CSP, HSTS, X-Frame-Options, etc.)
- **Attack Prevention** (null byte, control character filtering)
- **Security Logging** with attack attempt tracking
- **SSL/TLS** required for database connections
- **Container Security** (non-root user, minimal base image)

### **Security Scanning**

The CI/CD pipeline includes:
- **Trivy** vulnerability scanning
- **Google Cloud** container image scanning


## Related Projects

- **Frontend**: https://github.com/David-Xilo/safehouse-main-front
- **Infrastructure**: https://github.com/David-Xilo/safehouse-orchestration
- **Database**: https://github.com/David-Xilo/safehouse-db-schema

---
