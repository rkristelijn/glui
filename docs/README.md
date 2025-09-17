# GLUI Developer Documentation

Technical documentation for contributors and maintainers.

## 1. Context (Technical)

| Document | Purpose |
|----------|---------|
| [Architecture](architecture.md) | C4 model, design decisions, tech stack |
| [Code Principles](this-code-principles.md) | GLUI-specific principles (CLI UX, error handling) |
| [General Principles](general-principles.md) | Universal coding principles (KISS, YAGNI, etc.) |
| [Milestones](milestones.md) | Development roadmap and validation checkpoints |

**Current Architecture**: Go application with GitLab API client, planned CLI/TUI interfaces using cobra/bubbletea.

## 2. Prerequisites (Development)

- **Go 1.21+** with modules support
- **Security tools**: `make install-security` (govulncheck)
- **Testing tools**: testify, expect (for E2E)
- **GitLab token** for integration testing
- **Development environment**: See [Developer Guide](developer-guide.md)

## 3. Local Setup (Development)

```bash
# Full development setup
make deps              # Install dependencies
make install-security  # Install security tools
cp .env.example .env   # Configure environment
make build            # Build binary
make test-all         # Run all tests
```

**Development workflow**: TDD approach, conventional commits, pre-commit hooks enforced.

## 4. Validation (Technical)

| Test Type | Command | Coverage |
|-----------|---------|----------|
| Unit | `make test` | 81.8% (GitLab client) |
| Integration | `make test-integration` | Real GitLab API |
| E2E | `make test-e2e` | User workflows |
| Security | `make audit` | Vulnerability scan |
| Golden Files | `make update-golden` | TUI snapshots |

**Quality gates**: All tests pass, no vulnerabilities, conventional commits.

## 5. Release (Technical)

```bash
# Development process
git commit -m "feat(scope): description"  # Conventional commits
make test-all                            # Validate changes
git tag v0.1.0                          # Version tagging
git push --tags                         # Trigger release

# Future: Automated releases via GitHub Actions
```

**Versioning**: Semantic versioning, automated changelog generation.

## 6. More Information (Technical)

### Core Documentation
| Document | Purpose |
|----------|---------|
| [Developer Guide](developer-guide.md) | Security, conventions, Go best practices, setup |
| [Code Template](code-template.md) | Template for new files with documentation links |
| [Documentation Orchestration](documentation-orchestration.md) | Rules for user vs developer docs |
| [GitHub Setup](github-setup.md) | Repository setup, CI/CD, secrets |

### Project Management
| Document | Purpose |
|----------|---------|
| [TODO](../TODO.md) | Current progress and next steps |
| [Milestones](milestones.md) | Development phases and validation |

### Code Organization
- **Package headers**: All files link to relevant documentation
- **Testing**: Unit, integration, E2E, golden file testing
- **Security**: Vulnerability scanning, secret management
- **Architecture**: Clean separation of concerns (API, Core, UI)

---

**For end users**: See [root README.md](../README.md) for quick start and usage instructions.
