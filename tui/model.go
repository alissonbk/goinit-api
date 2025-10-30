package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/utils"
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
	hotPink                 = lipgloss.Color("#FF06B7")
	darkGray                = lipgloss.Color("#767676")
	blue                    = lipgloss.Color("#0090FF")
	red                     = lipgloss.Color("#CC3000")
	projectNamePage         = 0
	httpLibraryPage         = 1
	projectStructurePage    = 2
	databaseQueriesPage     = 3
	databaseDriverPage      = 4
	loggingPage             = 5
	loggingDefaultPage      = 6
	loggingNestedPage       = 7
	loggingLevelPage        = 8
	keycloakServiceAuthPage = 9
	customPanicHandlerPage  = 10
	godotEnvPage            = 11
	dockerfilePage          = 12
	_endPage                = 13
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// var style = lipgloss.NewStyle().
// 	Bold(true).
// 	Foreground(hotPink).
// 	Background(darkGray).
// 	PaddingTop(2).
// 	PaddingLeft(4).
// 	PaddingBottom(2).
// 	Width(80)

var inputStyle = lipgloss.NewStyle().Foreground(hotPink)

var listTitleStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Background(blue)

var grayStyle = lipgloss.NewStyle().Foreground(darkGray)

var errorStyle = lipgloss.NewStyle().Foreground(red)

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

func (m TuiModel) updateListModel(msg tea.Msg, attributeIndex int) (TuiModel, tea.Cmd) {
	reflectedAttribute := m.form.getAttributeByReflectionIndex(attributeIndex)
	m2, cmd := reflectedAttribute.Update(msg)
	m.form.setAttributeByReflectionIndex(attributeIndex, &m2)
	(reflectedAttribute).ResetSelected()
	return m, tea.Batch([]tea.Cmd{cmd}...)
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			// FIXME: do a better checking
			if strings.Trim(m.form.projectName.Value(), " ") == "" {
				m.err = fmt.Errorf("please inform a project name")
				return m, nil
			}
			m.currentPage += 1
			// FIXME: remove this page rounding
			if m.currentPage > _endPage {
				m.currentPage = 0
			}

		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		m.form.projectName.Focus()

		switch msg.String() {
		case "y":
			if m.currentPage == _endPage {
				panic("program ended sucessfully")
			}
		case "n":
			if m.currentPage == _endPage {
				m.currentPage = 0
			}
		}

	case error:
		fmt.Println("reached an error, ", msg)
		m.err = msg
		return m, nil

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()

		for i := 1; i < _endPage; i++ {
			m.form.updateListModelSizeByReflectionIndex(i, msg.Width-h*3, msg.Height-v*3)
		}

	}

	if m.currentPage == projectNamePage {
		m2, cmd := m.form.projectName.Update(msg)
		m.form.projectName = m2
		return m, tea.Batch([]tea.Cmd{cmd}...)
	}

	if m.currentPage == httpLibraryPage {
		return m.updateListModel(msg, httpLibraryPage)
	}

	if m.currentPage == projectStructurePage {
		return m.updateListModel(msg, projectStructurePage)
	}

	if m.currentPage == databaseQueriesPage {
		return m.updateListModel(msg, databaseQueriesPage)
	}

	if m.currentPage == databaseDriverPage {
		return m.updateListModel(msg, databaseDriverPage)
	}

	if m.currentPage == loggingPage {
		return m.updateListModel(msg, loggingPage)
	}

	if m.currentPage == loggingDefaultPage {
		return m.updateListModel(msg, loggingDefaultPage)
	}

	if m.currentPage == loggingNestedPage {
		return m.updateListModel(msg, loggingNestedPage)
	}

	if m.currentPage == loggingLevelPage {
		return m.updateListModel(msg, loggingLevelPage)
	}

	if m.currentPage == keycloakServiceAuthPage {
		return m.updateListModel(msg, keycloakServiceAuthPage)
	}

	if m.currentPage == customPanicHandlerPage {
		return m.updateListModel(msg, customPanicHandlerPage)
	}

	if m.currentPage == godotEnvPage {
		return m.updateListModel(msg, godotEnvPage)
	}

	if m.currentPage == dockerfilePage {
		return m.updateListModel(msg, dockerfilePage)
	}

	if m.currentPage == _endPage {
		return m, nil
	}

	panic("didn't match any page")
}

func (m TuiModel) renderListViewByIndex(idx int) string {
	attribute := m.form.getAttributeByReflectionIndex(idx)
	return fmt.Sprintf(`


		%s


		%s
		`,
		attribute.View(),
		grayStyle.Render("press ctrl+c to quit."))
}

func (m TuiModel) renderError() string {
	if m.err != nil {
		return errorStyle.Render(m.err.Error())
	}

	return ""
}

func (m TuiModel) View() string {
	if m.currentPage > _endPage {
		panic("current page bigger than _endPage " + strconv.Itoa(_endPage))
	}
	if m.currentPage == 0 {
		return fmt.Sprintf(`
		%s
		%s


		%s

		%s
		`,
			inputStyle.Bold(true).Width(30).Render("Project Name"),
			inputStyle.Render(m.form.projectName.View()),
			m.renderError(),
			grayStyle.Render("press ctrl+c to quit."))
	}

	if m.currentPage < _endPage {
		return m.renderListViewByIndex(m.currentPage)
	}

	utils.ClearScreen()
	return fmt.Sprintf(`
	
		Project name: %s
		Http library: %s
		Project structure: %s
		Database queries: %s
		Database driver: %s
		Logging: %s	
		LoggingDefault: %s	
		LoggingNested: %s	
		LoggingLevel: %s	
		KeycloakSA: %s	
		CustomPanicHandler: %s	
		Godotenv: %s	
		Dockerfile: %s	

		%s
		`,
		inputStyle.Width(30).Render(m.form.projectName.Value()),
		inputStyle.Width(30).Render(m.form.HttpLibrary.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.ProjectStructure.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.DatabaseQueries.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.DatabaseDriver.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.Logging.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.LoggingDefault.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.LoggingNested.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.LoggingLevel.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.KeycloakSA.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.CustomPanicHandler.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.Godotenv.SelectedItem().FilterValue()),
		inputStyle.Width(30).Render(m.form.Dockerfile.SelectedItem().FilterValue()),
		grayStyle.Render("Confirm? (y/n)"),
	)

}
