package tui

import (
	"github.com/alissonbk/goinit-api/constant"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

var _ list.Item = (*listItem)(nil)

type listItem struct {
	title, desc string
	// for boolean options will be 0 - false 1 - true
	evalue uint8
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return i.desc }
func (i listItem) FilterValue() string { return i.title }

type form struct {
	projectName        textinput.Model
	httpLibrary        list.Model
	projectStructure   list.Model
	databaseQueries    list.Model
	databaseDriver     list.Model
	logging            list.Model
	loggingDefault     list.Model
	loggingNested      list.Model
	loggingLevel       list.Model
	keycloakSA         list.Model
	customPanicHandler list.Model
	godotenv           list.Model
	dockerfile         list.Model
}

func defaultBoolList(trueTitle string, falseTitle string) list.Model {
	return list.New(
		[]list.Item{
			listItem{
				title:  trueTitle,
				desc:   "",
				evalue: uint8(1),
			},
			listItem{
				title:  falseTitle,
				desc:   "",
				evalue: uint8(0),
			},
		}, list.NewDefaultDelegate(), 0, 0,
	)
}

func defaultFromList(lst []string) list.Model {
	return list.New(
		func() []list.Item {
			list := make([]list.Item, len(lst))
			for i, s := range lst {
				list[i] = listItem{
					title:  s,
					desc:   s,
					evalue: uint8(i),
				}
			}

			return list
		}(), list.NewDefaultDelegate(), 0, 0,
	)
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
		projectStructure: list.New([]list.Item{
			listItem{
				title:  "MVC",
				desc:   "MVC project structure",
				evalue: uint8(constant.MVC),
			},
			listItem{
				title:  "Hexagonal",
				desc:   "Simplified hexagonal project structure",
				evalue: uint8(constant.MVC),
			},
		}, list.NewDefaultDelegate(), 0, 0),
		databaseQueries: list.New([]list.Item{
			listItem{
				title:  "GORM",
				desc:   "A ORM library for Golang.",
				evalue: uint8(constant.MVC),
			},
			listItem{
				title:  "sqlx",
				desc:   "A set of extensions on go's standard database/sql",
				evalue: uint8(constant.MVC),
			},
		}, list.NewDefaultDelegate(), 0, 0),
		databaseDriver: defaultFromList(constant.AllDatabaseDrivers()),
		logging: list.New(
			[]list.Item{
				listItem{
					title:  "Logrus",
					desc:   "Structured logger for Go, completely API compatible with the standard library logger.",
					evalue: uint8(constant.Logrus),
				},
				listItem{
					title:  "uber/zap",
					desc:   "Blazing fast, structured, leveled logging in Go.",
					evalue: uint8(constant.Zap),
				},
			}, list.NewDefaultDelegate(), 0, 0,
		),

		loggingDefault:     defaultBoolList("Yes", "No"),
		loggingNested:      defaultBoolList("Nested", "Structured"),
		loggingLevel:       defaultFromList(constant.AllLogLevels()),
		keycloakSA:         defaultBoolList("Yes", "No"),
		customPanicHandler: defaultBoolList("Yes", "No"),
		godotenv:           defaultBoolList("Yes", "No"),
		dockerfile:         defaultBoolList("Yes", "No"),
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
