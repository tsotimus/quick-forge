# Start from a minimal Go environment
FROM golang:1.22-alpine

# Install basic tools
RUN apk add --no-cache bash git

# Set the working directory
WORKDIR /app

# Copy your Go files into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the CLI
RUN go build -o quickforge main.go

# Run the CLI as the default command
CMD ["./quickforge"]