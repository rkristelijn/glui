package gitlab

import "time"

// Pipeline represents a GitLab CI/CD pipeline
type Pipeline struct {
	ID     int       `json:"id"`
	Status string    `json:"status"`
	Ref    string    `json:"ref"`
	WebURL string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
}
