# TODO

Current focus: **M1 Core Engine** - Expand TUI functionality and add more views.

## M0: Foundation ✅ COMPLETE

### 1. Project Setup ✅ DONE
- [x] Initialize Go module with proper structure
- [x] Setup testing framework (testify)
- [x] Add Makefile with common tasks (build, test, lint)
- [x] Setup security audit (govulncheck)
- [x] Add E2E testing framework
- [x] Add golden file testing for TUI snapshots

### 2. GitLab API Client ✅ DONE
- [x] Write test for `GetPipelines()` with mock response
- [x] Implement minimal GitLab client
- [x] Add authentication (token from env/config)
- [x] Add basic error handling (network, auth, not found)
- [x] Test with real GitLab instance (integration test)

### 3. Core Engine Foundation ✅ DONE
- [x] Write test for core pipeline service
- [x] Implement core interface
- [x] Add basic caching (in-memory for now)
- [x] Test error propagation from API to core

### 4. End-to-End TUI ✅ DONE
- [x] Connect GitLab API to core service via adapter
- [x] Implement k9s-style TUI with real data
- [x] Add .env file loading
- [x] Test with real GitLab project
- [x] Pipeline navigation and status indicators

**Test Status**: ✅ All tests pass, working e2e prototype with real GitLab data

## M1: Core Engine (Current Sprint)

### 1. Multi-View Support 🔄 NEXT
- [ ] Add merge requests view (`m` key)
- [ ] Add issues view (`i` key)
- [ ] Implement view state management
- [ ] Add view switching tests

### 2. Enhanced Navigation 🔄 NEXT
- [ ] Pipeline drill-down (ENTER to view jobs)
- [ ] Job details view
- [ ] Log viewing (`l` key)
- [ ] Back navigation (ESC key)

### 3. Error Handling & UX 🔄 NEXT
- [ ] Loading states in TUI
- [ ] Error display in TUI
- [ ] Connection retry logic
- [ ] Graceful API failure handling

## M2: Advanced Features (.soon)
- Search/filtering (`/` key)
- Multi-project support
- Configuration management
- Auto-refresh intervals

## M3: Polish & Performance (.soon)
- Performance optimization
- Memory usage optimization
- Documentation polish
- Release preparation

---

## Current Architecture

```
GitLab API → Adapter → Core Service → TUI
     ↓           ↓          ↓         ↓
  HTTP Client  Convert   Cache    k9s-style
  Auth/Error   Models   30s TTL   Navigation
```

## Decisions Made
- ✅ GitLab API library: custom HTTP client
- ✅ TUI framework: tview (works well)
- ✅ Config format: .env file (simple)
- ✅ Caching: in-memory with TTL

## Notes
- Working prototype connects to real GitLab
- k9s-style navigation implemented
- Ready for feature expansion
