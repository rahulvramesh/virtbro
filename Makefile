# Makefile

# Variables
APP_NAME = virtbro
SRC_DIR = ./cmd
MAIN_FILE = ./main.go

# Default target
all: build

# Build the binary
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) $(MAIN_FILE)

# Run the application
run: build
	@echo "Running the application..."
	./$(APP_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt $(SRC_DIR)/*.go
	go fmt $(MAIN_FILE)

# Lint the code (requires golangci-lint)
lint:
	@echo "Linting the code..."
	golangci-lint run

# Test the code
test:
	@echo "Running tests..."
	go test ./...

.PHONY: all build run deps clean fmt lint test
