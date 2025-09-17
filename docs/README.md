# GLUI Developer Documentation

## Project Status: Working Prototype ✅

GLUI now has a **working end-to-end prototype** that connects to real GitLab APIs and displays pipeline data in a k9s-style TUI.

## Quick Start

```bash
# 1. Set up GitLab token
export GITLAB_TOKEN="glpat-your-token"
export GITLAB_PROJECT_ID="your-group/your-project"

# 2. Build and run
make build
./glui
```

## Architecture Overview

```
GitLab API → Adapter → Core Service → TUI
     ↓           ↓          ↓         ↓
  HTTP Client  Convert   Cache    k9s-style
  Auth/Error   Models   30s TTL   Navigation
```

## Documentation Structure

### Core Documentation
- **[Architecture](architecture.md)** - System design and component interactions
- **[Use Cases](use-cases.md)** - User workflows and interaction patterns
- **[Developer Guide](developer-guide.md)** - Contributing and development setup

### Implementation Details
- **[UI Mockups](ui-mockups.md)** - k9s-style interface designs
- **[Code Principles](this-code-principles.md)** - Development guidelines
- **[Milestones](milestones.md)** - Project roadmap and phases

### Project Management
- **[GitHub Setup](github-setup.md)** - Repository configuration
- **[Documentation Orchestration](documentation-orchestration.md)** - Doc management

## Current Implementation

### ✅ Working Features
- **GitLab API Client** - HTTP client with auth and error handling
- **Core Service** - Business logic with 30-second caching
- **k9s-style TUI** - Pipeline navigation with j/k keys
- **Configuration** - .env file loading
- **Status Indicators** - Color-coded pipeline states

### 🔄 Next Features (M1)
- Merge requests view
- Issues view
- Pipeline drill-down
- Job details and logs

## Development

```bash
# Run tests
make test

# Build
make build

# Run with real GitLab data
./glui
```

## Contributing

1. Read [Developer Guide](developer-guide.md)
2. Follow [Code Principles](this-code-principles.md)
3. Check [TODO.md](../TODO.md) for current tasks
4. Submit PRs with tests

---

**Status**: M0 Foundation complete, M1 Core Engine in progress
