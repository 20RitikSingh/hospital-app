# Makefile for Hospital App
# Usage:
#   make build   # Build the application
#   make run     # Build and run the application
#   make test    # Run all tests
#   make clean   # Remove build artifacts

# Variables
BINARY=./bin/hospital-app
SRC=./cmd/hospital-app

.PHONY: build run test clean

build:
	@echo "Building the application..."
	@go build -o $(BINARY) $(SRC)

run: build
	@echo "Running the application..."
	$(BINARY)

test:
	@echo "Running tests..."
	@go test ./... -v

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY)
