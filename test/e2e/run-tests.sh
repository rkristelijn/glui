#!/bin/bash

# E2E tests for GLUI TUI
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
BINARY="$PROJECT_ROOT/glui"
GOLDEN_DIR="$PROJECT_ROOT/test/golden"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Running GLUI E2E Tests${NC}"

# Build binary
echo "Building GLUI..."
cd "$PROJECT_ROOT"
make build

# Test 1: CLI help output
echo "Testing CLI help..."
if command -v expect >/dev/null 2>&1; then
    expect -c "
        spawn $BINARY --help
        expect {
            \"Usage:\" { 
                puts \"✅ Help command works\"
                exit 0 
            }
            timeout { 
                puts \"❌ Help command timeout\"
                exit 1 
            }
            eof {
                puts \"✅ Help command completed\"
                exit 0
            }
        }
    "
else
    echo -e "${YELLOW}⚠️  expect not installed, skipping interactive tests${NC}"
    echo "Install with: brew install expect"
fi

# Test 2: CLI mode detection
echo "Testing CLI mode detection..."
output=$($BINARY pipelines 2>&1 || true)
if [[ "$output" == *"CLI mode - not implemented yet"* ]]; then
    echo -e "${GREEN}✅ CLI mode detection works${NC}"
else
    echo -e "${RED}❌ CLI mode detection failed${NC}"
    echo "Expected: 'CLI mode - not implemented yet'"
    echo "Got: $output"
    exit 1
fi

# Test 3: TUI mode detection  
echo "Testing TUI mode detection..."
output=$(timeout 2s $BINARY 2>&1 || true)
if [[ "$output" == *"TUI mode - not implemented yet"* ]]; then
    echo -e "${GREEN}✅ TUI mode detection works${NC}"
else
    echo -e "${RED}❌ TUI mode detection failed${NC}"
    echo "Expected: 'TUI mode - not implemented yet'"
    echo "Got: $output"
    exit 1
fi

echo -e "${GREEN}🎉 All E2E tests passed!${NC}"
