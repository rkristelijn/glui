#!/bin/bash
set -e

# Check if CLI help matches documented commands
echo "Checking CLI documentation sync..."

# Generate current help output
./glui --help > /tmp/current-help.txt 2>/dev/null || echo "glui not built yet"

# Check if docs/cli-reference.md exists and is current
if [ -f "docs/cli-reference.md" ]; then
    if ! diff -q /tmp/current-help.txt docs/cli-reference.md > /dev/null 2>&1; then
        echo "❌ CLI documentation is out of sync. Run 'make docs-cli' to update."
        exit 1
    fi
fi

# Check if version in README matches code
VERSION_README=$(grep -o 'v[0-9]\+\.[0-9]\+\.[0-9]\+' README.md | head -1)
VERSION_CODE=$(grep -o 'version.*=.*"[^"]*"' main.go | cut -d'"' -f2 2>/dev/null || echo "")

if [ -n "$VERSION_README" ] && [ -n "$VERSION_CODE" ] && [ "$VERSION_README" != "v$VERSION_CODE" ]; then
    echo "❌ Version mismatch: README has $VERSION_README, code has v$VERSION_CODE"
    exit 1
fi

echo "✅ Documentation sync check passed"
