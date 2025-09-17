package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/nlrxk0145/glui/internal/adapter"
	"github.com/nlrxk0145/glui/internal/core"
	"github.com/nlrxk0145/glui/internal/gitlab"
)

type TUI struct {
	app             *tview.Application
	main            *tview.Flex
	view            string
	cursor          int
	pipelines       []core.Pipeline
	pipelineService *core.PipelineService
	projectID       string
}

func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		return // .env file doesn't exist, use environment variables
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if os.Getenv(key) == "" { // Don't override existing env vars
				os.Setenv(key, value)
			}
		}
	}
}

func NewTUI() *TUI {
	// Load .env file if it exists
	loadEnv()

	// Get GitLab config from environment
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITLAB_TOKEN environment variable not set")
		fmt.Println("Set it with: export GITLAB_TOKEN=\"your-token\"")
		fmt.Println("Or add it to .env file")
		os.Exit(1)
	}

	baseURL := os.Getenv("GITLAB_URL")
	if baseURL == "" {
		baseURL = "https://gitlab.com"
	}

	projectID := os.Getenv("GITLAB_PROJECT_ID")
	if projectID == "" {
		// Try DEFAULT_GROUP as fallback
		defaultGroup := os.Getenv("DEFAULT_GROUP")
		if defaultGroup != "" {
			fmt.Printf("Error: GITLAB_PROJECT_ID not set. Try:\n")
			fmt.Printf("export GITLAB_PROJECT_ID=\"%s/your-project\"\n", defaultGroup)
		} else {
			fmt.Println("Error: GITLAB_PROJECT_ID environment variable not set")
			fmt.Println("Example: export GITLAB_PROJECT_ID=\"mygroup/myproject\"")
		}
		os.Exit(1)
	}

	// Setup GitLab client and core service
	gitlabClient := gitlab.NewClient(baseURL, token)
	adapter := adapter.NewGitLabAdapter(gitlabClient)
	pipelineService := core.NewPipelineService(adapter)

	return &TUI{
		app:             tview.NewApplication(),
		view:            "pipelines",
		pipelineService: pipelineService,
		projectID:       projectID,
	}
}

func (t *TUI) Run() error {
	t.loadData()
	t.setupUI()
	t.app.SetInputCapture(t.handleInput)
	return t.app.SetRoot(t.main, true).Run()
}

func (t *TUI) setupUI() {
	t.main = tview.NewFlex().SetDirection(tview.FlexRow)

	// Header
	header := tview.NewTextView().
		SetText(fmt.Sprintf(" Context: %s <%s>%s<glui> ", t.projectID, t.view, strings.Repeat(" ", 50))).
		SetBackgroundColor(tcell.ColorDefault)

	// Content area
	content := tview.NewTextView()
	content.SetDynamicColors(true)
	content.SetBackgroundColor(tcell.ColorDefault)

	t.renderContent(content)

	// Footer
	footer := tview.NewTextView().
		SetText(fmt.Sprintf("<%s> %s(%d)%s%s",
			t.view[:1], t.getViewName(), len(t.pipelines),
			strings.Repeat(" ", 60), time.Now().Format("15:04"))).
		SetBackgroundColor(tcell.ColorDefault)

	t.main.AddItem(header, 1, 0, false).
		AddItem(content, 0, 1, true).
		AddItem(footer, 1, 0, false)
}

func (t *TUI) renderContent(content *tview.TextView) {
	var output strings.Builder

	if len(t.pipelines) == 0 {
		output.WriteString("┌────────────────────────────────────────────────────────────────────────────────┐\n")
		output.WriteString("│ Loading pipelines...                                                          │\n")
		output.WriteString("└────────────────────────────────────────────────────────────────────────────────┘")
		content.SetText(output.String())
		return
	}

	// Table header
	output.WriteString("┌────────────────────────────────────────────────────────────────────────────────┐\n")
	output.WriteString("│ ID       STATUS    BRANCH      CREATED                                          │\n")
	output.WriteString("├────────────────────────────────────────────────────────────────────────────────┤\n")

	// Table rows
	for i, pipeline := range t.pipelines {
		prefix := "│ "
		if i == t.cursor {
			prefix = "│[black:blue] "
		}

		line := prefix + t.formatPipelineRow(pipeline)
		if i == t.cursor {
			line += "[white:default]"
		}
		line += " │\n"
		output.WriteString(line)
	}

	// Fill remaining space
	for i := len(t.pipelines); i < 15; i++ {
		output.WriteString("│                                                                                │\n")
	}

	output.WriteString("└────────────────────────────────────────────────────────────────────────────────┘")

	content.SetText(output.String())
}

func (t *TUI) formatPipelineRow(pipeline core.Pipeline) string {
	status := t.formatStatus(pipeline.Status)
	age := t.formatAge(pipeline.CreatedAt)

	return fmt.Sprintf("%-8s %-9s %-11s %-20s",
		strconv.Itoa(pipeline.ID), status, pipeline.Ref, age)
}

func (t *TUI) formatStatus(status string) string {
	switch status {
	case "success":
		return "[green]● success[white]"
	case "failed":
		return "[red]✗ failed[white]"
	case "running":
		return "[yellow]◐ running[white]"
	case "pending":
		return "[yellow]⏸ pending[white]"
	case "canceled":
		return "- canceled"
	case "skipped":
		return "- skipped"
	default:
		return status
	}
}

func (t *TUI) formatAge(createdAt time.Time) string {
	duration := time.Since(createdAt)

	if duration < time.Hour {
		return fmt.Sprintf("%dm ago", int(duration.Minutes()))
	} else if duration < 24*time.Hour {
		return fmt.Sprintf("%dh ago", int(duration.Hours()))
	} else {
		return fmt.Sprintf("%dd ago", int(duration.Hours()/24))
	}
}

func (t *TUI) loadData() {
	if t.view == "pipelines" {
		pipelines, err := t.pipelineService.ListPipelines(t.projectID)
		if err != nil {
			// Show error in UI
			t.pipelines = []core.Pipeline{}
			return
		}
		t.pipelines = pipelines
	}
}

func (t *TUI) handleInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyUp:
		if t.cursor > 0 {
			t.cursor--
			t.refresh()
		}
		return nil
	case tcell.KeyDown:
		if t.cursor < len(t.pipelines)-1 {
			t.cursor++
			t.refresh()
		}
		return nil
	}

	switch event.Rune() {
	case 'q':
		t.app.Stop()
		return nil
	case 'r':
		t.loadData()
		t.refresh()
	case 'j':
		if t.cursor < len(t.pipelines)-1 {
			t.cursor++
			t.refresh()
		}
	case 'k':
		if t.cursor > 0 {
			t.cursor--
			t.refresh()
		}
	}
	return event
}

func (t *TUI) refresh() {
	t.setupUI()
	t.app.SetRoot(t.main, true)
}

func (t *TUI) getViewName() string {
	return "Pipelines"
}
