BINARY=crawler

# Build the binary
build:
	@echo "Building..."
	go build -o ${BINARY}

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean up
clean:
	@echo "Cleaning..."
	rm -f ${BINARY}

# Clean up
clean-data:
	@echo "Cleaning data..."
	rm ./data/*.json

# Run the application
run: build
	@echo "Running application..."
	./$(BINARY)

# Help
help:
	@echo "Makefile commands:"
	@echo "build 			- Compiles the Go application"
	@echo "test  			- Runs the Go tests"
	@echo "clean 			- Removes the compiled binary"
	@echo "clean-data - Removes generated json files from data directory"
	@echo "run   			- Builds and runs the application"
	@echo "help  			- Displays this help message"

.PHONY: build test clean clean-data run help
