package golden

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update golden files")

// Assert compares output with golden file, creating/updating if needed
func Assert(t *testing.T, actual string, goldenFile string) {
	t.Helper()
	
	goldenPath := filepath.Join("test", "golden", goldenFile)
	
	if *update {
		// Update golden file
		err := os.MkdirAll(filepath.Dir(goldenPath), 0755)
		require.NoError(t, err)
		
		err = os.WriteFile(goldenPath, []byte(actual), 0644)
		require.NoError(t, err)
		
		t.Logf("Updated golden file: %s", goldenPath)
		return
	}
	
	// Compare with existing golden file
	expected, err := os.ReadFile(goldenPath)
	if os.IsNotExist(err) {
		t.Fatalf("Golden file %s does not exist. Run with -update to create it.", goldenPath)
	}
	require.NoError(t, err)
	
	assert.Equal(t, string(expected), actual, "Output differs from golden file %s", goldenPath)
}

// AssertTUISnapshot captures TUI output for snapshot testing
func AssertTUISnapshot(t *testing.T, output string, step string) {
	t.Helper()
	goldenFile := step + ".txt"
	Assert(t, output, goldenFile)
}
