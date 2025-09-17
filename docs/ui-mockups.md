# GLUI UI Mockups - k9s Style

## Pipelines View

```
 Context: myproject/main <pipelines>                                    <glui> 
┌────────────────────────────────────────────────────────────────────────────────┐
│ NAME     STATUS  BRANCH      COMMIT    AUTHOR       DURATION  AGE              │
├────────────────────────────────────────────────────────────────────────────────┤
│ 156    ● passed  main        a1b2c3d4  john.doe     3m42s     2h               │
│ 155      passed  main        e5f6g7h8  jane.smith   4m12s     1d               │
│ 154    ✗ failed  feature/x   i9j0k1l2  bob.wilson   2m18s     2d               │
│ 153    ◐ running develop     m3n4o5p6  alice.brown  1m33s     2d               │
│ 152    - skipped hotfix      q7r8s9t0  john.doe     -         3d               │
│ 151      passed  main        u1v2w3x4  jane.smith   5m01s     3d               │
│ 150    ✗ failed  feature/y   y5z6a7b8  bob.wilson   1m45s     4d               │
│ 149      passed  main        c9d0e1f2  alice.brown  3m28s     5d               │
│ 148    ⚠ warning develop     g3h4i5j6  john.doe     4m55s     6d               │
│ 147      passed  main        k7l8m9n0  jane.smith   2m12s     1w               │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
└────────────────────────────────────────────────────────────────────────────────┘
<pipeline> Pipelines(10)                                                    0:00 
```

## Merge Requests View

```
 Context: myproject/main <mrs>                                          <glui> 
┌────────────────────────────────────────────────────────────────────────────────┐
│ IID    STATUS    TITLE                           AUTHOR       BRANCH     AGE   │
├────────────────────────────────────────────────────────────────────────────────┤
│ !23  ● opened    fix: resolve login issue       jane.smith   fix/login  4h    │
│ !22    merged    feat: add user dashboard       john.doe     feature    1d    │
│ !21  ◐ draft     wip: refactor auth system      bob.wilson   auth-v2    2d    │
│ !20    closed    docs: update API documentation alice.brown  docs       3d    │
│ !19    merged    chore: bump dependencies       jane.smith   deps       5d    │
│ !18  ✗ conflicts fix: navbar responsive issue   john.doe     ui-fix     1w    │
│ !17    merged    test: add integration tests    bob.wilson   tests      1w    │
│ !16    closed    feat: dark mode support        alice.brown  dark-mode  2w    │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
│                                                                                │
└────────────────────────────────────────────────────────────────────────────────┘
<mr> Merge Requests(8)                                                      0:00 
```

## Key k9s Features to Match

- Clean table layout with minimal borders
- Status indicators: `●` (active), `✗` (failed), `◐` (running), `-` (inactive), `⚠` (warning)
- Context in header: `Context: project/branch <resource>`
- Resource count in footer: `<resource> ResourceName(count)`
- Time display in bottom right
- No extra decorative elements
- Consistent column alignment
- Selected row highlighting (cursor)
