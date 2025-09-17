package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		// CLI mode
		fmt.Println("CLI mode - not implemented yet")
		os.Exit(1)
	} else {
		// Interactive TUI mode
		fmt.Println("TUI mode - not implemented yet")
		os.Exit(1)
	}
}
