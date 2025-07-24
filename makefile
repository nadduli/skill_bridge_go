BINARY_NAME=server
OUTPUT_DIR=bin
SOURCE_PATH=./cmd/server/main.go

.PHONY: run build clean

# Build the project
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SOURCE_PATH)

# Run with live-reload using CompileDaemon
run:
	CompileDaemon \
		-build="go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SOURCE_PATH)" \
		-command="$(OUTPUT_DIR)/$(BINARY_NAME)" \
		-pattern="(.+\.go|.env)"

# Clean build artifacts
clean:
	rm -rf $(OUTPUT_DIR)