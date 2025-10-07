package tui

import (
	"github.com/alissonbk/goinit-api/constant"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Configuration struct {
	ProjectName      string
	HttpLibrary      constant.HttpLibrary
	ProjectStructure constant.ProjectStructure
	DatabaseQueries  constant.DatabaseQueries
	DatabaseDriver   constant.DatabaseDriver
	Logging          struct {
		Option     constant.LoggingOptions
		Structured bool
		Loglevel   constant.LogLevel
	}
	KeycloakServiceAuth bool
	CustomPanicHandler  bool
	GodotEnv            bool
	Dockerfile          bool
}

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#3C3C3C")).
	PaddingTop(2).
	PaddingLeft(4).
	PaddingBottom(2).
	Width(80)

var _ tea.Model = (*TuiModel)(nil)

type TuiModel struct {
	cursor        int
	currentPage   int
	form          form
	configuration Configuration
	selected      map[int]uint8 // uint8 will be any constant type also works for y/n case
}

func NewTuiModel() TuiModel {
	return TuiModel{
		cursor:      0,
		currentPage: 0,
		form: form{
			projectName: initialProjectNameInput(),
		},
		selected: make(map[int]uint8),
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m TuiModel) View() string {
	return style.Render("press ctrl+c or q to quit.")
}
