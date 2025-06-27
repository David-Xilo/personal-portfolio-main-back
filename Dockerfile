# Build stage
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install swag tool for generating Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation to the correct location (module root/docs)
RUN swag init -g src/internal/controllers/controller_manager.go -o docs

# Build the application
RUN go build -o main ./src/cmd/api

# Production stage
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=build /app/main .

# Copy Swagger docs to the expected location (module root/docs)
COPY --from=build /app/docs ./docs

# Expose port and run
EXPOSE 4000
CMD ["./main"]