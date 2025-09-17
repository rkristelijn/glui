package main

import (
	"fmt"
	"os"

	"github.com/nlrxk0145/glui/cmd"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("CLI mode not implemented yet")
		return
	}

	tui := cmd.NewTUI()
	if err := tui.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
