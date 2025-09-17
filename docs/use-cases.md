# GLUI Use Cases

## Primary Use Cases (k9s-style GitLab workflow)

### 1. Pipeline Monitoring
**As a developer, I want to monitor CI/CD pipelines without leaving my terminal**

- Launch `glui` and see current pipeline status
- Navigate with `j/k` or arrow keys through pipeline list
- See status indicators: ● (running), ✓ (passed), ✗ (failed), ⚠ (warning)
- Press `ENTER` to drill down into pipeline jobs
- Press `l` to view pipeline logs
- Press `r` to refresh pipeline status

### 2. Merge Request Review
**As a reviewer, I want to quickly check MR status and details**

- Press `m` to switch to merge requests view
- See MR status: ● (open), ✓ (merged), ✗ (conflicts), ◐ (draft)
- Navigate through MRs and press `ENTER` for details
- View MR description, changes, and pipeline status
- Quick approve/comment workflow

### 3. Issue Triage
**As a maintainer, I want to triage issues efficiently**

- Press `i` to switch to issues view
- Filter by labels, assignee, or status
- Quick navigation through issue list
- Press `ENTER` to view issue details
- Assign issues or add labels

### 4. Job Debugging
**As a developer, I want to debug failed CI jobs quickly**

- From pipelines view, drill down to failed pipeline
- See job breakdown with status indicators
- Press `ENTER` on failed job to view logs
- Navigate between job stages
- Retry failed jobs if permissions allow

### 5. Repository Overview
**As a team lead, I want a quick project health overview**

- Default dashboard view showing:
  - Recent pipeline status
  - Open MRs requiring attention
  - Critical issues
  - Repository activity
- Quick navigation to any resource type

## Secondary Use Cases

### 6. Branch Management
- Press `b` to view branches
- See branch status, last commit, pipeline status
- Switch between branches for different contexts

### 7. Multi-Project Monitoring
- Switch between different GitLab projects
- Maintain context per project
- Quick project switching with saved configurations

### 8. Search and Filtering
- Press `/` to enter search mode
- Filter pipelines by branch, author, or status
- Filter MRs by reviewer, label, or status
- Save common filter patterns

## Interaction Patterns (k9s-inspired)

### Navigation
- `↑/↓` or `j/k` - Navigate lists
- `ENTER` - Drill down/view details
- `ESC` - Go back to previous view
- `q` - Quit application

### View Switching
- `p` - Pipelines view
- `m` - Merge requests view
- `i` - Issues view
- `j` - Jobs view (when in pipeline context)
- `b` - Branches view
- `0-9` - Quick switch to numbered views

### Actions
- `r` - Refresh current view
- `d` - Show details panel
- `l` - View logs (context-dependent)
- `/` - Search/filter
- `?` - Show help/shortcuts
- `c` - Cancel/retry (context-dependent)

### Context Management
- Views remember cursor position
- Drill-down preserves navigation stack
- Auto-refresh for real-time updates
- Project context switching

## User Personas

### 1. DevOps Engineer
- Monitors multiple pipelines across projects
- Needs quick access to logs and job details
- Wants real-time status updates
- Focuses on pipeline health and deployment status

### 2. Developer
- Checks own MR and pipeline status
- Reviews code and approves MRs
- Debugs failed CI jobs
- Tracks issue assignments

### 3. Team Lead/Maintainer
- Triages issues and reviews MRs
- Monitors overall project health
- Manages releases and deployments
- Oversees team productivity

### 4. QA Engineer
- Tracks testing pipeline results
- Reviews MRs for quality gates
- Monitors deployment to staging environments
- Reports and tracks bugs

## Success Metrics

- **Speed**: Navigate to any resource in <3 keystrokes
- **Efficiency**: Complete common tasks without mouse
- **Visibility**: See critical information at a glance
- **Productivity**: Reduce context switching between tools
