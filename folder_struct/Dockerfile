# Use the official Golang image as a base image
FROM golang:1.21.3-alpine3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Make directory to store logs
RUN mkdir .logs

# Download dependencies
RUN go mod download

# Copy the entire application to the container
COPY . .

# Build the Go application
RUN go build -o main cmd/server/main.go

# Expose port 8080 for the Golang application
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
