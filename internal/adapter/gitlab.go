package adapter

import (
	"github.com/nlrxk0145/glui/internal/core"
	"github.com/nlrxk0145/glui/internal/gitlab"
)

// GitLabAdapter adapts gitlab.Client to core.GitLabClient interface
type GitLabAdapter struct {
	client gitlab.Client
}

func NewGitLabAdapter(client gitlab.Client) *GitLabAdapter {
	return &GitLabAdapter{client: client}
}

func (a *GitLabAdapter) GetPipelines(projectID string) ([]core.Pipeline, error) {
	gitlabPipelines, err := a.client.GetPipelines(projectID)
	if err != nil {
		return nil, err
	}

	// Convert gitlab.Pipeline to core.Pipeline
	corePipelines := make([]core.Pipeline, len(gitlabPipelines))
	for i, p := range gitlabPipelines {
		corePipelines[i] = core.Pipeline{
			ID:        p.ID,
			Status:    p.Status,
			Ref:       p.Ref,
			WebURL:    p.WebURL,
			CreatedAt: p.CreatedAt,
		}
	}

	return corePipelines, nil
}
