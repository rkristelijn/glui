# GLUI

GitLab Terminal UI - Navigate GitLab from your terminal with k9s-style interface.

## 1. Context

Terminal interface for GitLab operations - browse pipelines, merge requests, issues, and jobs with keyboard navigation. Designed for developers who prefer CLI/TUI tools over web interfaces.

**Status**: Early development - GitLab API client working, CLI/TUI interfaces in progress.

## 2. Prerequisites

- Go 1.21+ ([install guide](https://golang.org/doc/install))
- GitLab personal access token ([create token](https://gitlab.com/-/profile/personal_access_tokens))

## 3. Local Setup

```bash
# Clone and build
git clone <repo-url>
cd glui
make build

# Configure GitLab access
export GITLAB_TOKEN="glpat-xxxxxxxxxxxxxxxxxxxx"
export GITLAB_URL="https://gitlab.com"  # optional

# Run (currently shows "not implemented yet")
./glui              # Interactive TUI mode
./glui pipelines    # CLI mode
```

## 4. Validation

```bash
# Run all tests
make test-all

# Security audit
make audit

# Test coverage
make test-coverage
```

## 5. Release

```bash
# Development releases
git tag v0.1.0
git push --tags

# Production releases (future)
# Will use GitHub Actions for automated releases
```

## 6. More Information

- [📖 Developer Documentation](docs/README.md) - Architecture, contributing, technical details
- [🚀 Developer Guide](docs/developer-guide.md) - Setup, conventions, best practices  
- [📋 Current Progress](TODO.md) - What's working, what's next
- [🎯 Roadmap](docs/milestones.md) - Development phases and timeline

---

**Inspiration**: [k9s](https://github.com/derailed/k9s) (Kubernetes TUI) • [glab-tui](https://github.com/gitlab-tui/glab-tui) (GitLab TUI)
