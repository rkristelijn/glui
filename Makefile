.PHONY: build test test-integration test-e2e test-all lint clean audit install-security update-golden docs-sync docs-validate deps

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build the binary
build:
	go build -o glui main.go

# Run unit tests
test:
	go test ./...

# Run integration tests (requires GITLAB_TOKEN)
test-integration:
	go test -tags=integration ./...

# Run E2E tests
test-e2e: build
	./test/e2e/run-tests.sh

# Run all tests
test-all: test test-integration test-e2e

# Update golden files for snapshot testing
update-golden:
	go test ./... -update

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

# Run the application
run:
	go run main.go

# Check documentation sync
docs-sync: build
	./scripts/check-docs-sync.sh

# Validate documentation
docs-validate:
	@command -v markdown-link-check >/dev/null 2>&1 || { echo "Installing markdown-link-check..."; npm install -g markdown-link-check; }
	find . -name "*.md" -exec markdown-link-check {} \;
