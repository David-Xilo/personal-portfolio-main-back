# Start with an Alpine-compatible Go base image
FROM golang:1.22-alpine AS build

# Install build dependencies
RUN apk add --no-cache git gcc g++ libc-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application as a static binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main ./src/main.go

# Create a new stage for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the previous build stage
COPY --from=build /app/main .

# Copy Swagger docs to the expected location
COPY --from=build /app/src/docs /app/docs

# Expose the port that the application listens on
EXPOSE 4000

# Define the command to run the application
CMD ["./main"]