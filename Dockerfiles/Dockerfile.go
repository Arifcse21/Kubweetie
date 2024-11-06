# Use the official Golang image as a build environment
FROM golang:bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the rest of the Go source code to the container
COPY . ./

# Build the Go application
RUN go build -o socket_client_load socket_client_load.go

# Start a new, minimal image to run the app
FROM alpine:latest

# Install necessary dependencies to run the app (like ca-certificates for HTTPS requests)
RUN apk add --no-cache ca-certificates

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=builder /app/socket_client_load ./

# Run the compiled Go binary
CMD ["./socket_client_load"]
