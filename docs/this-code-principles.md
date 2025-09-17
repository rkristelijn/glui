# GLUI Code Principles

Project-specific principles for building a GitLab terminal UI that's fast, reliable, and user-friendly.

## Core Principles

### Fail Fast, Fail Clear
API calls will fail. Handle it gracefully.

- Show specific error messages (auth failed, network timeout, repo not found)
- Don't retry silently - let users know what's happening
- Provide actionable next steps in error messages
- Log errors for debugging but keep UI clean

```go
// ✅ Good
if err := gitlab.GetPipelines(); err != nil {
    return fmt.Errorf("failed to fetch pipelines for %s: %w", repo, err)
}
```

### Responsive by Default
Terminal UIs must feel instant.

- Show loading indicators for API calls >200ms
- Use async operations for network requests
- Cache frequently accessed data (pipelines, MRs)
- Set reasonable timeouts (5s for API calls)
- Allow cancellation with Ctrl+C

### CLI-First Design
Optimize for keyboard warriors.

- Every action has a keyboard shortcut
- Provide comprehensive help text (`?` key)
- Use consistent navigation patterns (vim-like: j/k, h/l)
- Show current context in status bar
- Default to most common workflows

### Context Awareness
Remember where users are and what they're doing.

- Maintain navigation history (back/forward)
- Remember last selected items when returning to views
- Auto-refresh data that changes frequently (pipeline status)
- Preserve filters and search terms
- Show breadcrumbs for deep navigation

### Offline Graceful
Work when GitLab is unreachable.

- Cache last known state
- Show cached data with timestamps
- Allow offline browsing of previously loaded data
- Queue actions when offline, sync when back online
- Clear offline indicators

## Implementation Guidelines

- **Config**: Store in `~/.config/glui/` following XDG spec
- **Logging**: Use structured logging, separate debug/user logs  
- **Testing**: Mock GitLab API responses for reliable tests
- **Performance**: Profile API calls, optimize hot paths
- **Security**: Never log tokens, use secure credential storage

## UI Patterns

- **Lists**: Show most important info first (status, title, time)
- **Details**: Progressive disclosure - summary first, details on demand
- **Actions**: Confirm destructive operations
- **Search**: Real-time filtering, fuzzy matching
- **Status**: Use color coding consistently (red=failed, green=success, yellow=pending)
