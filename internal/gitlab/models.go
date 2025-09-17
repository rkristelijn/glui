// Package gitlab provides GitLab API models and data structures
//
// Documentation:
// - Architecture: docs/architecture.md (Data Flow sequence diagram)
// - Developer Guide: docs/developer-guide.md (Go Best Practices - Code Organization)
// - TODO: TODO.md (M0 Foundation - GitLab API Client)
package gitlab

import "time"

// Pipeline represents a GitLab CI/CD pipeline
type Pipeline struct {
	ID        int       `json:"id"`
	Status    string    `json:"status"`
	Ref       string    `json:"ref"`
	WebURL    string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
}
