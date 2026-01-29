package model

import (
	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/utils"
	"github.com/charmbracelet/bubbles/list"
)

type Configuration struct {
	ModuleName       string
	ProjectName      string
	ModulePath       string
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

func (c *Configuration) SetProjectName(name string) {
	c.ProjectName = name
}

func (c *Configuration) SetHttpLibrary(v *list.Model) {
	c.HttpLibrary = constant.HttpLibraryFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetProjectStructure(v *list.Model) {
	c.ProjectStructure = constant.ProjectStructureFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetDatabaseQueries(v *list.Model) {
	c.DatabaseQueries = constant.DatabaseQueriesFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetDatabaseDriver(v *list.Model) {
	c.DatabaseDriver = constant.DatabaseDriverFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetLoggingLevel(v *list.Model) {
	c.Logging.Loglevel = constant.LogLevelFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetLoggingOption(v *list.Model) {
	c.Logging.Option = constant.LoggingOptionsFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetLoggingNested(v *list.Model) {
	c.Logging.Structured = constant.LoggingNestedFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetKeycloakServiceAuth(v *list.Model) {
	c.KeycloakServiceAuth = constant.BooleanOptionFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetCustomPanicHandler(v *list.Model) {
	c.CustomPanicHandler = constant.BooleanOptionFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetGodotEnv(v *list.Model) {
	c.GodotEnv = constant.BooleanOptionFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}

func (c *Configuration) SetDockerfile(v *list.Model) {
	c.Dockerfile = constant.BooleanOptionFromString(utils.ExtractStringFromListItem(v.SelectedItem()))
}
