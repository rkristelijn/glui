# Development Principles

Write code that creates synergy with your team, tools, and future self through clarity, consistency, and shared understanding.

## Core Principles

### RTFM - Respect The Framework's Model
- **Just** use the docs
- **Embrace** idioms  
- **Don't** override conventions
- **Integrate**, don't isolate

*Example: React → use hooks; SpringBoot → follow module structure*

### C4C - Code for Clarity
Write code assuming the next developer is a psychopath who knows your address.

- Clear folder structure (features first, not types)
- Keep code greppable
- Functions fit on screen (no scrolling)
- Document only what's necessary

### KISS - Keep It Simple Stupid
Default to the dumbest version that works.

- If docs can't explain it, it's too clever
- Use minimal abstractions
- Don't build an engine when you need a tricycle
- Hide details, not the why

### YAGNI - You Aren't Gonna Need It
- Don't build until needed
- No speculative abstractions
- Don't generalize without second use case
- Avoid premature optimization

### HIPI - Hide Implementation, Present Interface
Encapsulate behind clear interfaces.

```go
// 🤔 Meh
if user.age > 18 && user.isActive && !user.isBanned { ... }

// ✅ Better  
if user.canAccessDashboard() { ... }
```

### NBI - Naming by Intention
Good names are documentation. Follow NITRO:

- **Name** by intention (not implementation)
- **Imply** types clearly
- **Tell** truth about behavior (query vs command)
- **Reflect** domain logic
- **Optimize** for clarity over brevity

```go
// 🤔 Meh
data := fetchData()
user.accessDashboard()

// ✅ Better
userList := fetchUsers()
canAccess := user.canAccessDashboard() // query
user.openDashboard()                   // command
```

## When to Break Rules

Principles are lighthouses, not laws. Break them only when you:
1. Deeply understand what they protect against
2. Have a deliberate reason
3. Document why you're deviating

**Default to the principle.** Craft lies in knowing when not to follow them, not in ignoring them.
