# GLUI Architecture

## Current Status: Working Prototype ✅

GLUI has a working end-to-end implementation connecting real GitLab APIs to a k9s-style TUI.

## System Overview

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   GitLab    │    │   Adapter   │    │    Core     │    │     TUI     │
│  API Client │───▶│   Layer     │───▶│   Service   │───▶│  Interface  │
│             │    │             │    │             │    │             │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
      │                    │                  │                  │
   HTTP/JSON          Model Convert       Caching           k9s-style
   Auth/Error         gitlab→core         30s TTL           Navigation
```

## Component Details

### 1. GitLab API Client (`internal/gitlab/`)

**Purpose**: HTTP client for GitLab API communication

**Implementation**:
```go
type Client interface {
    GetPipelines(repo string) ([]Pipeline, error)
}

type HTTPClient struct {
    baseURL string
    token   string
    client  *http.Client
}
```

**Features**:
- Bearer token authentication
- URL encoding for project paths
- HTTP error handling (401, 404, etc.)
- JSON response parsing

### 2. Adapter Layer (`internal/adapter/`)

**Purpose**: Convert between GitLab and Core models

**Implementation**:
```go
type GitLabAdapter struct {
    client gitlab.Client
}

func (a *GitLabAdapter) GetPipelines(projectID string) ([]core.Pipeline, error)
```

**Features**:
- Model transformation
- Error propagation
- Interface compliance

### 3. Core Service (`internal/core/`)

**Purpose**: Business logic and caching

**Implementation**:
```go
type PipelineService struct {
    client GitLabClient
    cache  map[string]cacheEntry
    mutex  sync.RWMutex
}
```

**Features**:
- 30-second TTL cache
- Thread-safe operations
- Error wrapping
- Interface abstraction

### 4. TUI Interface (`cmd/`)

**Purpose**: k9s-style terminal interface

**Implementation**:
- **Framework**: tview
- **Navigation**: j/k keys, arrow keys
- **Actions**: r (refresh), q (quit)
- **Display**: Color-coded status indicators

## Data Flow

### Pipeline Loading Sequence

```
User starts app
      │
      ▼
Load .env config
      │
      ▼
Create GitLab client
      │
      ▼
Wrap in adapter
      │
      ▼
Create core service
      │
      ▼
Initialize TUI
      │
      ▼
Call service.ListPipelines()
      │
      ▼
Check cache (30s TTL)
      │
      ▼
[Cache miss] → API call → Parse JSON → Cache result
      │
      ▼
Return pipelines to TUI
      │
      ▼
Render k9s-style table
```

## Configuration

### Environment Variables
```bash
GITLAB_TOKEN=glpat-xxxxxxxxxxxxxxxxxxxx
GITLAB_URL=https://gitlab.com
GITLAB_PROJECT_ID=mygroup/myproject
```

### .env File Support
- Automatic loading on startup
- Environment variable override
- Error handling for missing config

## Error Handling

### API Errors
- **401 Unauthorized**: Invalid token
- **404 Not Found**: Project doesn't exist or no access
- **Network errors**: Connection failures

### TUI Error Display
- Loading states
- Error messages in interface
- Graceful degradation

## Testing Strategy

### Unit Tests
- Core service with mocked GitLab client
- Adapter model conversion
- Cache behavior verification

### Integration Tests
- Real GitLab API calls
- End-to-end pipeline loading
- Configuration loading

## Performance Considerations

### Caching
- 30-second TTL reduces API calls
- Thread-safe cache implementation
- Memory-based storage

### TUI Rendering
- Efficient table updates
- Minimal redraws
- Responsive navigation

## Security

### Token Handling
- Environment variable storage
- No token logging
- Secure HTTP headers

### API Communication
- HTTPS only
- Bearer token authentication
- Input validation

---

## Next Phase: M1 Core Engine

### Planned Additions
- Merge requests API client
- Issues API client
- Multi-view TUI navigation
- Pipeline drill-down functionality
