# GLUI Documentation

GitLab Terminal UI - Navigate GitLab from your terminal with k9s-style interface.

## Quick Start

```bash
# Interactive mode
glui

# CLI mode  
glui pipelines my-repo
glui mrs --status=opened
glui jobs 12345 --follow
```

## Documentation

| Document | Purpose |
|----------|---------|
| [Developer Guide](developer-guide.md) | Security, conventions, Go best practices, setup |
| [Architecture](architecture.md) | C4 model, design decisions, tech stack |
| [Milestones](milestones.md) | Development roadmap and validation checkpoints |
| [General Principles](general-principles.md) | Universal coding principles (KISS, YAGNI, etc.) |
| [Code Principles](this-code-principles.md) | GLUI-specific principles (CLI UX, error handling) |

## Features

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

## Architecture Overview

```
┌─────────────┐    ┌─────────────┐
│   CLI Mode  │    │   TUI Mode  │
└──────┬──────┘    └──────┬──────┘
       │                  │
       └────────┬─────────┘
                │
        ┌───────▼────────┐
        │  Core Engine   │
        └───────┬────────┘
                │
    ┌───────────┼───────────┐
    │           │           │
┌───▼───┐  ┌───▼────┐  ┌──▼───┐
│ Cache │  │ GitLab │  │ Config│
│       │  │  API   │  │       │
└───────┘  └────────┘  └──────┘
```

## Development

### Setup
```bash
git clone <repo>
cd glui
go mod tidy
go run main.go
```

### Testing
```bash
go test ./...                    # Unit tests
go test -tags=integration ./...  # Integration tests
```

### Building
```bash
go build -o glui main.go
```

## Configuration

Default config location: `~/.config/glui/config.yaml`

```yaml
gitlab:
  url: "https://gitlab.com"
  token: "your-token"
  
cache:
  ttl_pipelines: "30s"
  ttl_mrs: "2m"
  
ui:
  theme: "dark"
  vim_keys: true
```

## Contributing

1. Follow [General Principles](general-principles.md)
2. Follow [Code Principles](this-code-principles.md)  
3. Update docs for new features
4. Add tests for new functionality

## Inspiration

- [k9s](https://github.com/derailed/k9s) - Kubernetes TUI
- [glab-tui](https://github.com/gitlab-tui/glab-tui) - GitLab TUI reference
