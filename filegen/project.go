package filegen

import (
	"os"
	"os/exec"

	"github.com/alissonbk/goinit-api/codegen"
	"github.com/alissonbk/goinit-api/model"
)

// Can put the full path as project name if the user desire an alternative folder
func GenereateProject(cfg model.Configuration) {
	os.Mkdir(cfg.ProjectName, os.ModeDir)

	// FIXME
	cfg.ModuleName = "com." + cfg.ProjectName
	exec.Command("go", "mod", "init", cfg.ModuleName)
	exec.Command("go", "get", codegen.GenereateDependenciesList(cfg))
	exec.Command("go", "mod", "tidy")

	os.Exit(0)
}
