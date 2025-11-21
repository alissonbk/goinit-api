package model

import (
	"github.com/alissonbk/goinit-api/constant"
	"github.com/charmbracelet/bubbles/list"
)

type Configuration struct {
	ModuleName       string
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

func (c *Configuration) setProjectNameByList(v *list.Model) {
	c.ProjectName = v.Title
}

func (c *Configuration) setHttpLibrary(v *list.Model) {
	c.HttpLibrary = constant.HttpLibraryFromString(v.Title)
}
