# Makefile for File Browser

.PHONY: build run clean

# Binary name
BINARY_NAME=filebrowser

# Build the application
build:
	go build -o $(BINARY_NAME) *.go

# Run the application
run: build
	./$(BINARY_NAME)

# Clean build files
clean:
	go clean
	rm -f $(BINARY_NAME)
