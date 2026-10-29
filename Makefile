.PHONY: all build test clean cross-build install

BINARY_NAME=nophoneqr
BUILD_DIR=dist
VERSION=1.0.0

all: test build

build:
	go build -ldflags="-s -w" -o $(BINARY_NAME) .

test:
	go test -v ./...

install:
	go install .

cross-build: clean
	mkdir -p $(BUILD_DIR)
	# macOS Intel
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	# macOS Apple Silicon
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	# Windows x64
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	# Linux x64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "All binaries built in $(BUILD_DIR)/"

clean:
	rm -rf $(BUILD_DIR) $(BINARY_NAME) $(BINARY_NAME).exe
