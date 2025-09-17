# Documentation Orchestration

Rules for organizing documentation between user-facing and developer-facing content.

## The Two README Pattern

### Root README.md - User Interface
**Audience**: End users, QA, new developers, returning developers  
**Purpose**: One-page overview answering the 6 essential questions  
**Length**: Fits on one screen, no scrolling needed

**Must Answer**:
1. **Context** - What is this? Why does it exist?
2. **Prerequisites** - What do I need installed?
3. **Local Setup** - How do I run it locally?
4. **Validation** - How do I test my changes?
5. **Release** - How do I deploy/release?
6. **More Information** - Where can I learn more?

**Content Rules**:
- ✅ Business context and purpose
- ✅ Quick start commands (numbered steps)
- ✅ Links to detailed docs
- ❌ No code implementation details
- ❌ No architecture diagrams
- ❌ No development workflows

### docs/README.md - Developer API
**Audience**: Developers, maintainers, contributors  
**Purpose**: Navigation hub for all technical documentation  
**Length**: Can be longer, organized by concern

**Must Answer** (same 6 questions, developer perspective):
1. **Context** - Architecture, design decisions, tech stack
2. **Prerequisites** - Development tools, Go version, security setup
3. **Local Setup** - Development environment, testing framework
4. **Validation** - Testing strategies, CI/CD, code quality
5. **Release** - Build process, deployment pipeline, versioning
6. **More Information** - All technical docs, principles, guides

**Content Rules**:
- ✅ Links to all technical documentation
- ✅ Architecture and design decisions
- ✅ Development workflows and conventions
- ✅ Testing strategies and tools
- ❌ No business justification
- ❌ No end-user instructions

## Documentation Hierarchy

```
README.md                    # User interface (6 questions)
├── docs/
│   ├── README.md           # Developer API (6 questions, technical)
│   ├── architecture.md     # Technical design
│   ├── developer-guide.md  # How to contribute
│   ├── milestones.md       # Project roadmap
│   └── ...                 # All other technical docs
```

## Content Distribution Rules

### Root README.md Gets:
- **What**: Brief description and business purpose
- **Why**: Problem it solves, business value
- **Quick Start**: Minimal commands to get running
- **Links**: Pointers to detailed documentation

### docs/README.md Gets:
- **How**: Technical implementation details
- **Architecture**: System design and decisions
- **Development**: Contribution guidelines and workflows
- **Navigation**: Links to all technical documentation

## Examples

### Root README.md Structure
```markdown
# GLUI
GitLab Terminal UI - Navigate GitLab from your terminal

## 1. Context
Terminal interface for GitLab operations (pipelines, MRs, issues)

## 2. Prerequisites  
- Go 1.21+
- GitLab token

## 3. Local Setup
```bash
make build
export GITLAB_TOKEN="your-token"
./glui
```

## 4. Validation
```bash
make test-all
```

## 5. Release
```bash
git tag v1.0.0
git push --tags
```

## 6. More Information
- [📖 Documentation](docs/README.md)
- [🚀 Contributing](docs/developer-guide.md)
```

### docs/README.md Structure
```markdown
# GLUI Developer Documentation

## 1. Context (Technical)
| Document | Purpose |
|----------|---------|
| [Architecture](architecture.md) | C4 model, design decisions |

## 2. Prerequisites (Development)
- Go 1.21+, security tools, testing framework

## 3. Local Setup (Development)
- Development environment, testing, debugging

## 4. Validation (Technical)
- Unit tests, integration tests, E2E tests, security audit

## 5. Release (Technical)  
- Build process, CI/CD, deployment pipeline

## 6. More Information (Technical)
- All technical documentation links
```

## Maintenance Rules

### When to Update Root README.md:
- Business purpose changes
- Installation process changes
- Basic usage changes
- New major features that affect user workflow

### When to Update docs/README.md:
- New technical documentation added
- Architecture changes
- Development process changes
- New tools or frameworks added

### Validation Checklist:
- [ ] Root README.md answers all 6 questions for users
- [ ] docs/README.md answers all 6 questions for developers
- [ ] Root README.md fits on one screen
- [ ] No overlap between user and developer content
- [ ] All links work and point to correct documentation
- [ ] Both READMEs stay in sync with actual project state

## Benefits

**For Users**:
- Quick understanding without technical noise
- Clear path from zero to running
- No cognitive overhead from development details

**For Developers**:
- Complete technical context
- All development resources in one place
- Clear separation of concerns

**For Maintenance**:
- Clear ownership of content
- Reduced documentation drift
- Easier to keep both audiences happy
