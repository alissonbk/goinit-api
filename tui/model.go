package tui

import (
	"fmt"

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

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

// var style = lipgloss.NewStyle().
// 	Bold(true).
// 	Foreground(hotPink).
// 	Background(darkGray).
// 	PaddingTop(2).
// 	PaddingLeft(4).
// 	PaddingBottom(2).
// 	Width(80)

var inputStyle = lipgloss.NewStyle().Foreground(hotPink)

var grayStyle = lipgloss.NewStyle().Foreground(darkGray)

var _ tea.Model = (*TuiModel)(nil)

type TuiModel struct {
	currentPage   int
	form          *form
	err           error
	configuration Configuration
	selected      map[int]uint8 // uint8 will be any constant type also works for y/n case
}

func NewTuiModel() TuiModel {
	return TuiModel{
		currentPage: 0,
		form:        newForm(),
		selected:    make(map[int]uint8),
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			// FIXME: do a better checking
			if m.form.projectName.Value() == "" {
				m.err = fmt.Errorf("please inform a project name")
				return m, nil
			}
			m.currentPage += 1
		case tea.KeyCtrlC:
			return m, tea.Quit

		}

		m.form.projectName.Focus()
	case error:
		fmt.Println("reached an error, ", msg)
		m.err = msg
		return m, nil
	}

	m2, cmd := m.form.projectName.Update(msg)
	m.form.projectName = m2
	return m, tea.Batch([]tea.Cmd{cmd}...)
}

func (m TuiModel) View() string {

	if m.currentPage == 1 {
		return fmt.Sprintf(`


	%s



	%s
	`,
			m.form.httpLibrary.View(),
			grayStyle.Render("press ctrl+c to quit."))
	}
	return fmt.Sprintf(`


	%s
	%s



	%s
	`,
		inputStyle.Bold(true).Width(30).Render("Project Name"),
		inputStyle.Render(m.form.projectName.View()),
		grayStyle.Render("press ctrl+c to quit."))
}
