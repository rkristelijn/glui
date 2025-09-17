package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Client interface for GitLab API operations
type Client interface {
	GetPipelines(repo string) ([]Pipeline, error)
}

// HTTPClient implements the Client interface
type HTTPClient struct {
	baseURL string
	token   string
	client  *http.Client
}

// NewClient creates a new GitLab API client
func NewClient(baseURL, token string) Client {
	return &HTTPClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		token:   token,
		client:  &http.Client{},
	}
}

// GetPipelines fetches pipelines for a given repository
func (c *HTTPClient) GetPipelines(repo string) ([]Pipeline, error) {
	// URL encode the repo path
	encodedRepo := url.PathEscape(repo)
	url := fmt.Sprintf("%s/api/v4/projects/%s/pipelines", c.baseURL, encodedRepo)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: invalid token")
	}
	
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("repository not found: %s", repo)
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	
	var pipelines []Pipeline
	if err := json.NewDecoder(resp.Body).Decode(&pipelines); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	
	return pipelines, nil
}
