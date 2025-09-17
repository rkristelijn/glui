.PHONY: build test lint clean audit install-security

# Build the binary
build:
	go build -o glui main.go

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Lint code
lint:
	go vet ./...
	go fmt ./...

# Security audit
audit:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

# Install security tools
install-security:
	go install golang.org/x/vuln/cmd/govulncheck@latest

# Clean build artifacts
clean:
	rm -f glui coverage.out coverage.html

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run the application
run:
	go run main.go
