# TODO

Current focus: **M0 Foundation** - Get one feature working end-to-end with tests.

## M0: Foundation (Current Sprint)

### 1. Project Setup
- [ ] Initialize Go module with proper structure
  ```bash
  go mod init github.com/user/glui
  mkdir -p {cmd,internal/{core,gitlab,cache,config},test}
  ```
- [ ] Setup testing framework (testify)
- [ ] Add Makefile with common tasks (build, test, lint)
- [ ] Setup CI/CD pipeline (GitHub Actions)

### 2. GitLab API Client (TDD)
- [ ] Write test for `GetPipelines()` with mock response
- [ ] Implement minimal GitLab client
  ```go
  type Client interface {
      GetPipelines(repo string) ([]Pipeline, error)
  }
  ```
- [ ] Add authentication (token from env/config)
- [ ] Add basic error handling (network, auth, not found)
- [ ] Test with real GitLab instance

### 3. Core Engine Foundation
- [ ] Write test for core pipeline service
- [ ] Implement core interface
  ```go
  type Core interface {
      ListPipelines(repo string) ([]Pipeline, error)
  }
  ```
- [ ] Add basic caching (in-memory for now)
- [ ] Test error propagation from API to core

### 4. CLI Command
- [ ] Setup cobra CLI framework
- [ ] Implement `glui pipelines <repo>` command
- [ ] Add output formatting (table format)
- [ ] Add `--help` and basic flags
- [ ] Test CLI integration

### 5. Validation & Documentation
- [ ] Update architecture.md if design changed
- [ ] Add API.md with GitLab endpoints used
- [ ] Verify all tests pass
- [ ] Demo: `glui pipelines my-repo` works

## M1: Core Engine (.soon)
- State management
- Full caching layer
- All CLI commands
- Comprehensive error handling
- Configuration management

## M2: Basic TUI (.soon)
- TUI framework integration
- List views with navigation
- Keyboard shortcuts
- Status indicators

## M3: Advanced Navigation (.soon)
- Drill-down workflows
- Context preservation
- Auto-refresh
- Search/filtering

## M4: Polish & Performance (.soon)
- Performance optimization
- Multi-instance support
- Documentation polish
- Release preparation

## M5: Advanced Features (.soon)
- Pipeline creation
- MR templates
- Notifications
- Artifact downloads

---

## Current Blockers
*None - ready to start M0*

## Decisions Needed
- GitLab API library: custom vs existing (go-gitlab)
- TUI framework: bubbletea vs tview vs custom
- Config format: YAML vs TOML vs JSON

## Notes
- Study k9s architecture in `~/git/hub/k9s` for TUI patterns
- Keep each step small and testable
- Update this TODO after each completed item
