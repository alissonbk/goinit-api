package filegen

import (
	"os"
	"os/exec"

	"github.com/alissonbk/goinit-api/codegen"
	"github.com/alissonbk/goinit-api/model"
)

// Can put the full path as project name if the user desire an alternative folder
func GenereateProject(cfg model.Configuration) {
	os.Mkdir(cfg.ProjectName, 0755)

	// FIXME: user should be able to specify the module name
	cfg.ModuleName = "com." + cfg.ProjectName
	targetDir := "./" + cfg.ProjectName

	modinitCmd := exec.Command("go", "mod", "init", cfg.ModuleName)
	modinitCmd.Dir = targetDir
	_, err := modinitCmd.Output()
	if err != nil {
		panic("failed to run mod init command : " + err.Error())
	}

	depsCmd := exec.Command("go", "get", codegen.GenereateDependenciesList(cfg))
	depsCmd.Dir = targetDir
	_, err = depsCmd.Output()
	if err != nil {
		// panic(fmt.Sprintf("dependencies list: %v", codegen.GenereateDependenciesList(cfg)))
		panic("failed to run dependencies command : " + err.Error())
	}

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = targetDir
	_, err = tidyCmd.Output()
	if err != nil {
		panic("failed to run go mod tidy command : " + err.Error())
	}

	os.Exit(0)
}
