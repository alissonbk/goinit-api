package tui

import (
	"github.com/alissonbk/goinit-api/constant"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

var _ list.Item = (*listItem)(nil)

type listItem struct {
	title, desc string
	evalue      uint8
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return i.desc }
func (i listItem) FilterValue() string { return i.title }

type form struct {
	projectName textinput.Model
	httpLibrary list.Model
}

func newForm() *form {
	return &form{
		projectName: initialProjectNameInput(),
		httpLibrary: list.New([]list.Item{
			listItem{
				title:  "Gin",
				desc:   "Lightweight simple usage http library...",
				evalue: uint8(constant.Gin),
			},
			listItem{
				title:  "Fiber",
				desc:   "Fiber ...",
				evalue: uint8(constant.Gin),
			},
			listItem{
				title:  "Echo",
				desc:   "Lightweight simple usage http library...",
				evalue: uint8(constant.Gin),
			},
		}, list.NewDefaultDelegate(), 0, 0),
	}
}

func initialProjectNameInput() textinput.Model {
	input := textinput.New()
	input.Focus()
	input.CharLimit = 100
	input.Width = 30
	input.Prompt = ""
	// should create a function that returns error to validate the project name later
	input.Validate = nil

	return input
}
