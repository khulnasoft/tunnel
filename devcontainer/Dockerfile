# Use a lightweight Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /workspace

# Install any system dependencies you need
RUN apk add --no-cache \
    git \
    bash

# Set Go environment variables for Go modules
ENV GOPATH=/workspace/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Copy the go modules files and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Expose ports if needed
EXPOSE 8080

# Default command to run your Go application
CMD ["go", "run", "src/main.go"]
