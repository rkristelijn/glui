# GLUI

GitLab Terminal UI - Navigate GitLab from your terminal with k9s-style interface.

Browse pipelines, merge requests, issues, and jobs with keyboard navigation. Designed for developers who prefer CLI/TUI tools over web interfaces.

**Status**: Early development - GitLab API client working, CLI/TUI interfaces in progress.

## Installation

### Download Binary (Current)

```bash
# Download latest release
curl -L https://github.com/nlrxk0145/glui/releases/latest/download/glui-macos -o glui
chmod +x glui
sudo mv glui /usr/local/bin/
```

### Build from Source

```bash
git clone <repo-url>
cd glui
make build
sudo mv glui /usr/local/bin/
```

## Setup

### 1. Create GitLab Token

Create a personal access token at: https://gitlab.com/-/profile/personal_access_tokens

Required scopes: `api`, `read_user`, `read_repository`

### 2. Configure GLUI

**Option A: Environment Variables**
```bash
export GITLAB_TOKEN="glpat-xxxxxxxxxxxxxxxxxxxx"
export GITLAB_URL="https://gitlab.com"  # optional, defaults to gitlab.com
```

**Option B: Configuration File**
```bash
# Create .env file in your project or home directory
echo "GITLAB_TOKEN=glpat-xxxxxxxxxxxxxxxxxxxx" > .env
echo "GITLAB_URL=https://gitlab.com" >> .env
```

## Usage

```bash
# Interactive TUI mode
glui

# CLI mode (specific commands)
glui pipelines
glui merge-requests
glui issues
```

**Note**: Currently shows "not implemented yet" - TUI interface in development.

## More Information

- [📖 Developer Documentation](docs/README.md) - Contributing, architecture, technical details
- [📋 Current Progress](TODO.md) - What's working, what's next
- [🎯 Roadmap](docs/milestones.md) - Development phases and timeline

---

**Inspiration**: [k9s](https://github.com/derailed/k9s) (Kubernetes TUI) • [glab-tui](https://github.com/gitlab-tui/glab-tui) (GitLab TUI)
