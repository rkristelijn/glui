// Package gitlab provides GitLab API integration tests
//
// Documentation:
// - Developer Guide: docs/developer-guide.md (Testing - Integration Testing)
// - Code Principles: docs/this-code-principles.md (Fail Fast, Fail Clear)
// - TODO: TODO.md (M0 Foundation - Test with real GitLab instance)
package gitlab

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Integration test - uses public GitLab repo (no token needed)
func TestClient_GetPipelines_Integration(t *testing.T) {
	url := os.Getenv("GITLAB_URL")
	if url == "" {
		url = "https://gitlab.com"
	}

	// Use empty token for public repo access
	client := NewClient(url, "")

	// Test with GitLab's own public repo
	pipelines, err := client.GetPipelines("gitlab-org/gitlab")
	require.NoError(t, err)

	// Basic validation
	require.NotEmpty(t, pipelines, "Should have at least one pipeline")

	// Validate structure
	p := pipelines[0]
	require.NotZero(t, p.ID, "Pipeline should have ID")
	require.NotEmpty(t, p.Status, "Pipeline should have status")
	require.NotEmpty(t, p.Ref, "Pipeline should have ref")
	require.NotEmpty(t, p.WebURL, "Pipeline should have web URL")
	require.False(t, p.CreatedAt.IsZero(), "Pipeline should have created timestamp")
}
