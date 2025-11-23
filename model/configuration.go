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

func (c *Configuration) SetProjectName(v *list.Model) {
	c.ProjectName = v.Title
}

func (c *Configuration) SetHttpLibrary(v *list.Model) {
	c.HttpLibrary = constant.HttpLibraryFromString(v.Title)
}

func (c *Configuration) SetProjectStructure(v *list.Model) {
	c.ProjectStructure = constant.ProjectStructureFromString(v.Title)
}

func (c *Configuration) SetDatabaseQueries(v *list.Model) {
	c.DatabaseQueries = constant.DatabaseQueriesFromString(v.Title)
}

func (c *Configuration) SetDatabaseDriver(v *list.Model) {
	c.DatabaseDriver = constant.DatabaseDriverFromString(v.Title)
}

func (c *Configuration) SetLoggingLevel(v *list.Model) {
	c.Logging.Loglevel = constant.LogLevelFromString(v.Title)
}

func (c *Configuration) SetLoggingOption(v *list.Model) {
	c.Logging.Option = constant.LoggingOptionsFromString(v.Title)
}

func (c *Configuration) SetLoggingNested(v *list.Model) {
	c.Logging.Structured = constant.LoggingNestedFromString(v.Title)
}

func (c *Configuration) SetKeycloakServiceAuth(v *list.Model) {
	c.KeycloakServiceAuth = constant.BooleanOptionFromString(v.Title)
}

func (c *Configuration) SetCustomPanicHandler(v *list.Model) {
	c.CustomPanicHandler = constant.BooleanOptionFromString(v.Title)
}

func (c *Configuration) SetGodotEnv(v *list.Model) {
	c.GodotEnv = constant.BooleanOptionFromString(v.Title)
}

func (c *Configuration) SetDockerfile(v *list.Model) {
	c.Dockerfile = constant.BooleanOptionFromString(v.Title)
}
