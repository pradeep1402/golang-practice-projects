# Start from the official Golang base image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the server application
RUN go build -o /bookstore-server ./server

# Expose the port the gRPC server runs on
EXPOSE 50052

# Command to run the server
CMD ["/bookstore-server"]