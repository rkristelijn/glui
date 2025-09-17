# GLUI Milestones

Strategic roadmap with baby steps, continuous validation against design, and testable increments.

## M0: Foundation (Week 1-2)
**Goal**: Minimal viable structure with one working feature

- ✅ Documentation foundation
- 🔄 Basic Go project setup with testing
- 🔄 GitLab API client (pipelines only)
- 🔄 CLI mode: `glui pipelines <repo>`
- 🔄 Core architecture validation

**Success Criteria**: Can list pipelines from CLI, tests pass, architecture matches docs

## M1: Core Engine (Week 3-4)
**Goal**: Solid foundation for both CLI and TUI modes

- State management
- Caching layer
- Error handling patterns
- CLI commands for all core features
- Comprehensive test coverage

**Success Criteria**: All CLI commands work, offline mode, error handling tested

## M2: Basic TUI (Week 5-6)
**Goal**: Interactive mode with keyboard navigation

- TUI framework integration (bubbletea)
- Basic list views (pipelines, MRs)
- Keyboard navigation
- Status indicators

**Success Criteria**: Can navigate GitLab data interactively, matches k9s UX patterns

## M3: Advanced Navigation (Week 7-8)
**Goal**: Workflow support (MR → Pipeline → Job → Logs)

- Drill-down navigation
- Context preservation
- Auto-refresh for logs
- Search and filtering

**Success Criteria**: Complete user workflows work smoothly

## M4: Polish & Performance (Week 9-10)
**Goal**: Production-ready tool

- Performance optimization
- Configuration management
- Multi-instance support
- Documentation updates

**Success Criteria**: Fast, configurable, well-documented

## M5: Advanced Features (Future)
**Goal**: Power user features

- Pipeline creation with variables
- MR creation from templates
- Notifications
- Artifact downloads

## Validation Checkpoints

**Every milestone**:
- [ ] Architecture still matches docs/architecture.md
- [ ] Follows principles in docs/this-code-principles.md
- [ ] All tests pass
- [ ] Documentation updated
- [ ] Demo works end-to-end

## Risk Mitigation

**Technical Risks**:
- TUI framework limitations → Prototype early in M2
- GitLab API complexity → Start simple in M0
- Performance issues → Profile in M4

**Scope Risks**:
- Feature creep → Stick to milestone goals
- Over-engineering → Follow YAGNI principle
- Documentation drift → Update docs with each PR
