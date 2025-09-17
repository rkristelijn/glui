package core

import (
	"fmt"
	"sync"
	"time"
)

type Pipeline struct {
	ID        int       `json:"id"`
	Status    string    `json:"status"`
	Ref       string    `json:"ref"`
	WebURL    string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
}

type GitLabClient interface {
	GetPipelines(projectID string) ([]Pipeline, error)
}

type Core interface {
	ListPipelines(projectID string) ([]Pipeline, error)
}

type PipelineService struct {
	client GitLabClient
	cache  map[string]cacheEntry
	mutex  sync.RWMutex
}

type cacheEntry struct {
	pipelines []Pipeline
	timestamp time.Time
}

const (
	cacheTimeout = 30 * time.Second
	maxCacheSize = 100 // Prevent unbounded growth
)

func NewPipelineService(client GitLabClient) *PipelineService {
	return &PipelineService{
		client: client,
		cache:  make(map[string]cacheEntry),
	}
}

func (s *PipelineService) ListPipelines(projectID string) ([]Pipeline, error) {
	// Check cache first
	s.mutex.Lock()
	if entry, exists := s.cache[projectID]; exists {
		if time.Since(entry.timestamp) < cacheTimeout {
			pipelines := entry.pipelines
			s.mutex.Unlock()
			return pipelines, nil
		}
		// Remove expired entry
		delete(s.cache, projectID)
	}

	// Clean cache if too large
	if len(s.cache) >= maxCacheSize {
		s.cleanCache()
	}
	s.mutex.Unlock()

	// Fetch from API
	pipelines, err := s.client.GetPipelines(projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pipelines: %w", err)
	}

	// Update cache
	s.mutex.Lock()
	s.cache[projectID] = cacheEntry{
		pipelines: pipelines,
		timestamp: time.Now(),
	}
	s.mutex.Unlock()

	return pipelines, nil
}

// cleanCache removes expired entries (must be called with mutex held)
func (s *PipelineService) cleanCache() {
	now := time.Now()
	for key, entry := range s.cache {
		if now.Sub(entry.timestamp) >= cacheTimeout {
			delete(s.cache, key)
		}
	}
}
