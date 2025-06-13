# Run all tests in the project
.PHONY: test
test:
	go test -v ./...

# Build the Go application
.PHONY: build
build:
	go build