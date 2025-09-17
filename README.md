# GLUI

GitLab Terminal UI - Navigate GitLab from your terminal with k9s-style interface.

Browse pipelines, merge requests, issues, and jobs with keyboard navigation. Designed for developers who prefer CLI/TUI tools over web interfaces.

**Status**: ✅ **Working Prototype** - End-to-end pipeline viewing with real GitLab API integration.

## Features

✅ **Pipeline Monitoring** - View real-time pipeline status with k9s-style navigation  
✅ **GitLab API Integration** - Connects to any GitLab instance (gitlab.com or self-hosted)  
✅ **Smart Caching** - 30-second cache to reduce API calls  
✅ **Status Indicators** - Color-coded pipeline states (running, success, failed)  
✅ **Keyboard Navigation** - j/k navigation, r to refresh, q to quit  

## Installation

### Download Binary (Coming Soon)

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
export GITLAB_PROJECT_ID="mygroup/myproject"
```

**Option B: Configuration File**
```bash
# Create .env file in your project directory
echo "GITLAB_TOKEN=glpat-xxxxxxxxxxxxxxxxxxxx" > .env
echo "GITLAB_URL=https://gitlab.com" >> .env
echo "GITLAB_PROJECT_ID=mygroup/myproject" >> .env
```

## Usage

```bash
# Interactive TUI mode (current)
glui

# CLI mode (planned)
glui pipelines
glui merge-requests
glui issues
```

### TUI Controls

- `j/k` or `↑/↓` - Navigate pipeline list
- `r` - Refresh data from GitLab
- `q` - Quit application

## Current Status

**Working:**
- ✅ Real GitLab API connection
- ✅ Pipeline list view with live data
- ✅ k9s-style keyboard navigation
- ✅ Status indicators and formatting
- ✅ Configuration via .env file

**Planned:**
- 🔄 Merge requests view
- 🔄 Issues view
- 🔄 Pipeline job details
- 🔄 Log viewing
- 🔄 Multi-project support

## More Information

- [📖 Developer Documentation](docs/README.md) - Contributing, architecture, technical details
- [📋 Current Progress](TODO.md) - What's working, what's next
- [🎯 Use Cases](docs/use-cases.md) - User workflows and interaction patterns

---

**Inspiration**: [k9s](https://github.com/derailed/k9s) (Kubernetes TUI) • [glab-tui](https://github.com/gitlab-tui/glab-tui) (GitLab TUI)
