.PHONY: build run clean build-all release-local test docker-build docker-run
.DEFAULT_GOAL := build

# Build for current platform
build:
	go build -o quickforge

# Run the application
run:
	./quickforge

# Clean build artifacts
clean:
	rm -f quickforge quickforge-*

# Build for all supported platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o quickforge-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o quickforge-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o quickforge-linux-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o quickforge-linux-arm64

# Test the application (unit tests)
test:
	go test ./...

# Test with Docker (has limitations - see test.sh for details)
test-docker: docker-build
	./test.sh

# Test locally in dry-run mode (safe)
test-dry-run: build
	./quickforge --dry-run --non-interactive

# Test locally in non-interactive mode (caution: may install software)
test-local: build
	./quickforge --non-interactive

# Build Docker image
docker-build:
	docker build -t quickforge .

# Run Docker container
docker-run: docker-build
	docker run -it --rm quickforge

# Create a local release (for testing)
release-local: build-all
	mkdir -p release
	cp quickforge-* release/
	@echo "Local release created in ./release/"