// Package main provides the entry point for GLUI - GitLab Terminal UI
//
// Documentation:
// - Architecture: docs/architecture.md (Container Diagram - GLUI Application)
// - Developer Guide: docs/developer-guide.md (Project Setup, Environment Setup)
// - TODO: TODO.md (M0 Foundation progress)
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (ignore errors - it's optional)
	_ = godotenv.Load()
	
	if len(os.Args) > 1 {
		// CLI mode
		fmt.Println("CLI mode - not implemented yet")
		fmt.Printf("Would run command: %s\n", os.Args[1])
		if token := os.Getenv("GITLAB_TOKEN"); token != "" {
			fmt.Printf("GitLab token loaded: %s...\n", token[:10])
		} else {
			fmt.Println("No GITLAB_TOKEN found")
		}
		os.Exit(1)
	} else {
		// Interactive TUI mode
		fmt.Println("TUI mode - not implemented yet")
		if token := os.Getenv("GITLAB_TOKEN"); token != "" {
			fmt.Printf("GitLab token loaded: %s...\n", token[:10])
		} else {
			fmt.Println("No GITLAB_TOKEN found")
		}
		os.Exit(1)
	}
}
