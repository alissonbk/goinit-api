package filegen

import (
	"os"
	"os/exec"

	"github.com/alissonbk/goinit-api/codegen"
	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/model"
)

func createProjectFiles(cfg model.Configuration) {
	if cfg.ProjectStructure == constant.Hexagonal {
		// TODO
		panic("hexagonal structure not implemented")
	}

	if cfg.ProjectStructure == constant.MVC {
		err := os.Chdir(cfg.ProjectName)
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// main.go
		exec.Command("echo", codegen.GenerateMainContent(cfg), ">", "main.go")

		// dockerfile and compose
		if cfg.Dockerfile {
			exec.Command("echo", codegen.GenerateDockerfileContent(), ">", "Dockerfile")
			exec.Command("echo", codegen.GenerateDockerComposeContent(), ">", "docker-compose.yaml")
		}

		// config dir
		os.Mkdir("config", 0755)
		err = os.Chdir("config")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}
		// database.go
		exec.Command("echo", codegen.GenerateDatabaseContent(cfg), ">", "database.go")
		// logs.go
		exec.Command("echo", codegen.GenerateLogsContent(cfg), ">", "logs.go")
		// migrations dir
		os.Mkdir("migrations", 0755)
		err = os.Chdir("migrations")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}
		// migration up
		exec.Command("echo", codegen.GenerateExampleMigrationUpContent(), ">", "000001_create_example_table.up.sql")
		// migration down
		exec.Command("echo", codegen.GenerateExampleMigrationDownContent(), ">", "000001_create_example_table.down.sql")

		// back to main directory
		err = os.Chdir("..")
		err = os.Chdir("..")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// app dir
		os.Mkdir("app", 0755)
		err = os.Chdir("app")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		exec.Command("echo", codegen.GenerateBaseEntity())
	}
}

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

	// depsCmd := exec.Command("go", "get", codegen.GenereateDependenciesList(cfg))
	// depsCmd.Dir = targetDir
	// dpsOutput, err := depsCmd.CombinedOutput()
	// if err != nil {
	// 	panic(string(dpsOutput))
	// }

	createProjectFiles(cfg)

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = targetDir
	tidyOutput, err := tidyCmd.CombinedOutput()
	if err != nil {
		panic("failed to run go mod tidy command : " + string(tidyOutput))
	}

	os.Exit(0)
}
