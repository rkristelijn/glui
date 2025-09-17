.PHONY: build test test-integration test-all lint clean audit install-security deps

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build the binary
build:
	go build -o glui main.go

# Run unit tests
test:
	go test ./internal/...

# Run integration tests (requires GITLAB_TOKEN)
test-integration:
	@if [ -z "$(GITLAB_TOKEN)" ]; then \
		echo "GITLAB_TOKEN not set, skipping integration tests"; \
		exit 0; \
	fi
	go test -v ./internal/gitlab/... -run Integration

# Run all tests
test-all: test test-integration

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./internal/...
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

# Run the application
run:
	go run main.go
