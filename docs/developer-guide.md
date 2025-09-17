# Developer Guide

Essential information for contributing to GLUI safely and effectively.

## Security First

### Secrets Management
- **Never commit tokens** - use environment variables or config files (gitignored)
- **No hardcoded URLs** - make GitLab instances configurable
- **Validate all inputs** - especially repo names and user data
- **Use HTTPS only** - no plain HTTP for API calls

```bash
# Safe way to test
export GITLAB_TOKEN="your-token-here"
export GITLAB_URL="https://gitlab.com"
./glui pipelines my-repo
```

### Dependencies
- Pin dependency versions in go.mod
- Regularly update dependencies: `go get -u ./...`
- Review new dependencies before adding
- **Security audit**: `make audit` (uses govulncheck)

```bash
# Install security tools
make install-security

# Run security audit
make audit

# Check for outdated dependencies
go list -u -m all
```

## Git Conventions

### Commit Messages
Format: `type(scope): description`

**Types**: feat, fix, docs, style, refactor, test, chore  
**Scopes**: core, cli, tui, api, cache, config, docs

```bash
# Good examples
git commit -m "feat(api): add merge request listing"
git commit -m "fix(core): handle empty pipeline response"
git commit -m "docs: update architecture diagram"
git commit -m "test(cli): add integration tests"

# Bad examples  
git commit -m "stuff"
git commit -m "WIP"
git commit -m "fix bug"
```

### Branch Naming
- `feature/short-description`
- `fix/issue-description`
- `docs/update-readme`

## Project Setup

### Prerequisites
- Go 1.21+ ([install guide](https://golang.org/doc/install))
- Git
- Make (optional but recommended)

### Quick Start
```bash
# Clone and setup
git clone <repo-url>
cd glui
make deps          # Install dependencies
make test          # Run tests
make build         # Build binary

# Development workflow
make test          # Run tests before committing
make lint          # Format and vet code
./glui             # Test the binary
```

### Environment Setup
```bash
# Option 1: Use .env file (recommended for development)
cp .env.example .env
# Edit .env with your GitLab token

# Option 2: Export environment variables
export GITLAB_TOKEN="glpat-xxxxxxxxxxxxxxxxxxxx"
export GITLAB_URL="https://gitlab.com"  # or your instance

# Optional
export GLUI_LOG_LEVEL="debug"
```

**Note**: GLUI automatically loads `.env` file if present. Never commit `.env` with real tokens!

## Go Best Practices We Follow

### Code Organization
- **Interfaces in consumer packages** - define interfaces where you use them
- **Small interfaces** - prefer many small interfaces over large ones
- **Package naming** - short, lowercase, no underscores
- **Error handling** - always handle errors, wrap with context
- **Documentation links** - use [code template](code-template.md) for all new files

```go
// Good - interface where it's used
package core
type GitLabClient interface {
    GetPipelines(repo string) ([]Pipeline, error)
}

// Good - small, focused interface
type PipelineLister interface {
    GetPipelines(repo string) ([]Pipeline, error)
}

// Good - error wrapping
if err := client.GetPipelines(repo); err != nil {
    return fmt.Errorf("failed to fetch pipelines for %s: %w", repo, err)
}
```

### Testing
- **Test-driven development** - write tests first
- **Table-driven tests** for multiple scenarios
- **Mock external dependencies** - use interfaces
- **Test error cases** - not just happy path
- **Golden file testing** - for TUI snapshots
- **E2E testing** - full user workflows

```bash
# Unit tests
make test

# Integration tests (requires GITLAB_TOKEN)
export GITLAB_TOKEN="your-token"
make test-integration

# E2E tests (full user workflows)
make test-e2e

# All tests
make test-all

# Update golden files (TUI snapshots)
make update-golden
```

#### TUI Testing Patterns

**Golden File Testing** - Capture TUI output as snapshots:
```go
func TestPipelineList(t *testing.T) {
    output := renderPipelineList(mockData)
    golden.AssertTUISnapshot(t, output, "pipeline-list-step1")
}

// Run with -update to create/update golden files
go test ./... -update
```

**E2E Testing** - Full user workflows with expect:
```bash
# Install expect for interactive testing
brew install expect

# E2E tests simulate real user interactions
./test/e2e/run-tests.sh
```

**Integration Testing** - Real GitLab API calls:
```go
func TestRealGitLab(t *testing.T) {
    if os.Getenv("GITLAB_TOKEN") == "" {
        t.Skip("GITLAB_TOKEN not set")
    }
    // Test with real API
}
```

```go
func TestGetPipelines(t *testing.T) {
    tests := []struct {
        name    string
        repo    string
        want    []Pipeline
        wantErr bool
    }{
        {"valid repo", "owner/repo", mockPipelines, false},
        {"empty repo", "", nil, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Performance
- **Use context.Context** for cancellation and timeouts
- **Avoid premature optimization** - profile first
- **Pool expensive resources** - HTTP clients, connections
- **Stream large responses** - don't load everything into memory

### Security
- **Validate inputs** - especially user-provided data
- **Use prepared statements** - if we add database later
- **Sanitize outputs** - especially for terminal display
- **Timeout network calls** - prevent hanging

## Go Learning Resources

### Official
- [Go Tour](https://tour.golang.org/) - Interactive introduction
- [Effective Go](https://golang.org/doc/effective_go.html) - Idiomatic Go
- [Go by Example](https://gobyexample.com/) - Practical examples

### Books
- "The Go Programming Language" - Donovan & Kernighan
- "Go in Action" - Kennedy, Ketelsen, St. Martin

### Tools
- `go fmt` - Format code
- `go vet` - Find suspicious code
- `go test -race` - Race condition detection
- `go mod tidy` - Clean up dependencies

## Common Patterns We Use

### Error Handling
```go
// Wrap errors with context
return fmt.Errorf("operation failed: %w", err)

// Check for specific error types
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

### HTTP Clients
```go
// Reuse HTTP clients
var httpClient = &http.Client{
    Timeout: 10 * time.Second,
}

// Always set timeouts
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### Configuration
```go
// Use struct tags for config
type Config struct {
    GitLabURL   string `yaml:"gitlab_url" env:"GITLAB_URL"`
    Token       string `yaml:"token" env:"GITLAB_TOKEN"`
    CacheDir    string `yaml:"cache_dir" env:"GLUI_CACHE_DIR"`
}
```

## Debugging

### Logging
```go
// Use structured logging
log.Printf("fetching pipelines: repo=%s, page=%d", repo, page)

// Debug mode
if os.Getenv("GLUI_DEBUG") == "true" {
    log.Printf("debug: %+v", data)
}
```

### Testing
```bash
# Run specific test
go test -run TestGetPipelines ./internal/gitlab/

# Run with verbose output
go test -v ./...

# Run with race detection
go test -race ./...

# Generate coverage report
make test-coverage
open coverage.html
```
