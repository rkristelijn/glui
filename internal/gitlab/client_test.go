// Package gitlab provides GitLab API client tests
//
// Documentation:
// - Developer Guide: docs/developer-guide.md (Testing, TDD, Table-driven tests)
// - Code Principles: docs/this-code-principles.md (Fail Fast, Fail Clear)
// - TODO: TODO.md (M0 Foundation - GitLab API Client TDD)
package gitlab

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetPipelines(t *testing.T) {
	// Mock GitLab API response
	mockPipelines := []Pipeline{
		{
			ID:        123,
			Status:    "success",
			Ref:       "main",
			WebURL:    "https://gitlab.com/repo/-/pipelines/123",
			CreatedAt: time.Now(),
		},
		{
			ID:        124,
			Status:    "running",
			Ref:       "feature-branch",
			WebURL:    "https://gitlab.com/repo/-/pipelines/124",
			CreatedAt: time.Now(),
		},
	}

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v4/projects/test/repo/pipelines", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockPipelines)
	}))
	defer server.Close()

	// Create client
	client := NewClient(server.URL, "test-token")
	
	// Test GetPipelines
	pipelines, err := client.GetPipelines("test/repo")
	
	require.NoError(t, err)
	assert.Len(t, pipelines, 2)
	assert.Equal(t, 123, pipelines[0].ID)
	assert.Equal(t, "success", pipelines[0].Status)
	assert.Equal(t, "main", pipelines[0].Ref)
}

func TestClient_GetPipelines_AuthError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "401 Unauthorized"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "invalid-token")
	
	_, err := client.GetPipelines("test/repo")
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unauthorized")
}

func TestClient_GetPipelines_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "404 Project Not Found"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	
	_, err := client.GetPipelines("nonexistent/repo")
	
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
