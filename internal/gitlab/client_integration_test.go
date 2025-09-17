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

	"github.com/nlrxk0145/glui/test/golden"
	"github.com/stretchr/testify/require"
)

// Integration test - only runs with real GitLab token
func TestClient_GetPipelines_Integration(t *testing.T) {
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		t.Skip("GITLAB_TOKEN not set, skipping integration test")
	}
	
	url := os.Getenv("GITLAB_URL")
	if url == "" {
		url = "https://gitlab.com"
	}
	
	client := NewClient(url, token)
	
	// Test with a known public repo (GitLab's own repo)
	pipelines, err := client.GetPipelines("gitlab-org/gitlab")
	require.NoError(t, err)
	
	// Basic validation
	require.NotEmpty(t, pipelines, "Should have at least one pipeline")
	
	// Snapshot test - capture structure (not exact data since it changes)
	if len(pipelines) > 0 {
		p := pipelines[0]
		output := formatPipelineForSnapshot(p)
		golden.AssertTUISnapshot(t, output, "pipeline-structure")
	}
}

func formatPipelineForSnapshot(p Pipeline) string {
	// Format pipeline for consistent snapshot testing
	// Replace dynamic values with placeholders
	return "Pipeline Structure:\n" +
		"ID: <number>\n" +
		"Status: " + p.Status + "\n" +
		"Ref: <branch-name>\n" +
		"WebURL: <url>\n" +
		"CreatedAt: <timestamp>\n"
}
