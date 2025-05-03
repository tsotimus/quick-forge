# Start from a Debian-based Go environment
FROM golang:1.24

# Install basic tools needed for Homebrew and general usage
RUN apt-get update && \
    apt-get install -y bash git curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

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