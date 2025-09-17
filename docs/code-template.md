# Code File Template

Use this template for all new Go files to maintain documentation links.

## Package Header Template

```go
// Package [name] provides [brief description]
//
// Documentation:
// - Architecture: docs/architecture.md ([relevant section])
// - Developer Guide: docs/developer-guide.md ([relevant section])
// - Code Principles: docs/this-code-principles.md ([relevant principle])
// - TODO: TODO.md ([relevant milestone/task])
package [name]
```

## Examples by Package

### Core Engine (`internal/core/`)
```go
// Package core provides business logic and state management
//
// Documentation:
// - Architecture: docs/architecture.md (Component Diagram - Core Engine)
// - Developer Guide: docs/developer-guide.md (Go Best Practices - Code Organization)
// - Code Principles: docs/this-code-principles.md (Context Awareness, Responsive by Default)
// - TODO: TODO.md (M0 Foundation - Core Engine Foundation)
package core
```

### CLI Commands (`cmd/`)
```go
// Package cmd provides CLI command implementations using cobra
//
// Documentation:
// - Architecture: docs/architecture.md (Container Diagram - CLI Parser)
// - Developer Guide: docs/developer-guide.md (Go Best Practices - Error Handling)
// - Code Principles: docs/this-code-principles.md (CLI-First Design, Fail Fast)
// - TODO: TODO.md (M0 Foundation - CLI Command)
package cmd
```

### TUI Components (`internal/tui/`)
```go
// Package tui provides terminal UI components using bubbletea
//
// Documentation:
// - Architecture: docs/architecture.md (Container Diagram - Terminal UI)
// - Developer Guide: docs/developer-guide.md (TUI Testing Patterns)
// - Code Principles: docs/this-code-principles.md (Responsive by Default, CLI-First Design)
// - TODO: TODO.md (M2 Basic TUI)
package tui
```

### Cache Layer (`internal/cache/`)
```go
// Package cache provides local caching with TTL support
//
// Documentation:
// - Architecture: docs/architecture.md (Component Diagram - Cache Service)
// - Developer Guide: docs/developer-guide.md (Performance - Pool expensive resources)
// - Code Principles: docs/this-code-principles.md (Offline Graceful, Responsive by Default)
// - TODO: TODO.md (M0 Foundation - Basic caching)
package cache
```

### Configuration (`internal/config/`)
```go
// Package config provides configuration management and .env loading
//
// Documentation:
// - Architecture: docs/architecture.md (Scalability Considerations - Config)
// - Developer Guide: docs/developer-guide.md (Security First - Secrets Management)
// - Code Principles: docs/this-code-principles.md (Context Awareness)
// - TODO: TODO.md (M0 Foundation - Configuration)
package config
```

## Benefits

1. **Traceability** - Developers know which docs to update
2. **Context** - Links to relevant architecture and principles
3. **Planning** - Connection to TODO items and milestones
4. **Onboarding** - New developers can quickly find relevant docs
5. **Maintenance** - Easy to keep docs in sync with code changes

## Usage

1. Copy the appropriate template when creating new files
2. Update the documentation links when modifying existing code
3. Add new sections to docs when adding new concepts
4. Update TODO.md when completing tasks
