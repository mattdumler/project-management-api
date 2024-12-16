# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a lightweight image for the application
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./main"]

