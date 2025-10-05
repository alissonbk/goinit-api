package main

import (
	"fmt"
	"os"

	"github.com/alissonbk/goinit-api/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.NewTuiModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to run teabubble program %v", err)
		os.Exit(1)
	}
}
