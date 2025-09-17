# GLUI

GitLab Terminal UI - Navigate GitLab from your terminal with k9s-style interface.

## Status: M0 Foundation (40% Complete)

✅ **Working**: GitLab API client, testing framework, security audit  
🔄 **In Progress**: Core engine, CLI commands  
📋 **Next**: TUI interface, navigation workflows

## Quick Start

### Prerequisites
- Go 1.21+ ([install guide](https://golang.org/doc/install))
- GitLab token for API access

### Installation
```bash
git clone <repo-url>
cd glui
make deps          # Install dependencies
make build         # Build binary
```

### Usage
```bash
# Currently available (basic mode detection)
./glui             # TUI mode (not implemented yet)
./glui pipelines   # CLI mode (not implemented yet)

# Development
make test          # Run all tests
make test-e2e      # Run E2E tests
make audit         # Security audit
```

### Configuration
```bash
# Required for GitLab API access
export GITLAB_TOKEN="glpat-xxxxxxxxxxxxxxxxxxxx"
export GITLAB_URL="https://gitlab.com"  # optional, defaults to gitlab.com
```

## Planned Features

### Core Features
- **Pipelines** - List, monitor, create with custom variables
- **Merge Requests** - Browse, create from templates  
- **Jobs** - View logs with auto-refresh
- **Issues** - List and navigate
- **Navigation** - MR → Pipeline → Job → Logs workflow

### User Experience
- **Keyboard-first** - vim-like navigation (j/k, h/l)
- **Fast** - Caching with smart refresh
- **Offline-aware** - Works with cached data
- **Multi-instance** - Support cloud + on-prem GitLab

## Architecture

```
┌─────────────┐    ┌─────────────┐
│   CLI Mode  │    │   TUI Mode  │
└──────┬──────┘    └──────┬──────┘
       │                  │
       └────────┬─────────┘
                │
        ┌───────▼────────┐
        │  Core Engine   │  ✅ GitLab API Client
        └───────┬────────┘  🔄 Business Logic
                │           📋 Caching Layer
    ┌───────────┼───────────┐
    │           │           │
┌───▼───┐  ┌───▼────┐  ┌──▼───┐
│ Cache │  │ GitLab │  │ Config│
│       │  │  API   │  │       │
└───────┘  └────────┘  └──────┘
```

## Development

### Current Test Status
- ✅ All tests pass
- ✅ 81.8% code coverage
- ✅ No security vulnerabilities
- ✅ E2E testing framework ready

### Commands
```bash
make build         # Build binary
make test          # Unit tests
make test-e2e      # E2E tests
make test-all      # All tests
make audit         # Security audit
make lint          # Code formatting
make clean         # Clean artifacts
```

## Documentation

- [📖 Full Documentation](docs/README.md) - Architecture, principles, guides
- [🚀 Developer Guide](docs/developer-guide.md) - Setup, conventions, best practices
- [📋 TODO](TODO.md) - Current progress and next steps
- [🎯 Milestones](docs/milestones.md) - Development roadmap

## Inspiration

- [k9s](https://github.com/derailed/k9s) - Kubernetes TUI
- [glab-tui](https://github.com/gitlab-tui/glab-tui) - GitLab TUI reference

## Contributing

1. Follow [Developer Guide](docs/developer-guide.md)
2. Use conventional commits (`feat(scope): description`)
3. Add tests for new features
4. Update documentation

---

**Note**: This is early development. The GitLab API client is working, but CLI/TUI interfaces are not yet implemented. See [TODO.md](TODO.md) for current progress.
