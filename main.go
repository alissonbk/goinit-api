package main

import (
	"fmt"
	"os"

	"github.com/alissonbk/goinit-api/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.NewTuiModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to run teabubble program %v", err)
		os.Exit(1)
	}
}
