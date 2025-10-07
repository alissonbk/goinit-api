package tui

import "github.com/charmbracelet/bubbles/textinput"

type form struct {
	projectName textinput.Model
}

func initialProjectNameInput() textinput.Model {
	input := textinput.New()
	input.Focus()
	input.CharLimit = 20
	input.Width = 30
	input.Prompt = ""
	// should create a function that returns error to validate the project name later
	input.Validate = nil

	return input
}
