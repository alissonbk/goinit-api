package filegen

import (
	"os"
	"os/exec"

	"github.com/alissonbk/goinit-api/codegen"
	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/model"
)

const (
	DIR_PERM  = 0755
	FILE_PERM = 0644
)

// TODO return error message instead of panicking
func writeFile(filename, content string) {
	if err := os.WriteFile(filename, []byte(content), FILE_PERM); err != nil {
		panic("failed to write " + filename + ": " + err.Error())
	}
}

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
		writeFile("main.go", codegen.GenerateMainContent(cfg))

		// dockerfile and compose
		if cfg.Dockerfile {
			writeFile("Dockerfile", codegen.GenerateDockerfileContent())
			writeFile("docker-compose.yaml", codegen.GenerateDockerComposeContent())
		}

		if cfg.GodotEnv {
			envContent := codegen.GenerateEnvContent()
			writeFile(".env", envContent)
			writeFile(".env.example", envContent)
		}

		// config dir
		os.Mkdir("config", DIR_PERM)
		if err = os.Chdir("config"); err != nil {
			panic("failed to change directory: " + err.Error())
		}
		writeFile("database.go", codegen.GenerateDatabaseContent(cfg))
		writeFile("logs.go", codegen.GenerateLogsContent(cfg))

		// migrations
		os.Mkdir("migrations", DIR_PERM)
		writeFile("migrations/000001_create_example_table.up.sql", codegen.GenerateExampleMigrationUpContent())
		writeFile("migrations/000001_create_example_table.down.sql", codegen.GenerateExampleMigrationDownContent())

		// back to main directory
		err = os.Chdir("../")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// app dir
		os.Mkdir("app", DIR_PERM)
		err = os.Chdir("app")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// constant
		os.Mkdir("constant", DIR_PERM)
		writeFile("constant/status.go", codegen.GenerateConstantContent())

		// controller
		os.Mkdir("controller", DIR_PERM)
		writeFile("controller/example.go", codegen.GenerateControllerContent())

		// exeption
		os.Mkdir("exception", DIR_PERM)
		writeFile("exception/panic.go", codegen.GeneratePanicContent(cfg.Logging.Option))

		// model
		os.Mkdir("model", DIR_PERM)
		err = os.Chdir("model")
		os.Mkdir("entity", DIR_PERM)
		os.Mkdir("dto", DIR_PERM)
		writeFile("entity/base.go", codegen.GenerateBaseEntity())
		writeFile("entity/example.go", codegen.GenerateExampleEntity())
		err = os.Chdir("../")

		// repository
		os.Mkdir("repository", DIR_PERM)
		writeFile("repository/example.go", codegen.GenerateRepositoryContent(cfg.DatabaseQueries))

		// router
		os.Mkdir("router", DIR_PERM)
		writeFile("router/routes.go", codegen.GenerateRouterContent(cfg.HttpLibrary))
		writeFile("router/injection.go", codegen.GenerateInjectionContent())

		// service
		os.Mkdir("service", DIR_PERM)
		writeFile("service/example.go", codegen.GenerateServiceContent())

		err = os.Chdir("../")
	}
}

// Can put the full path as project name if the user desire an alternative folder
func GenereateProject(cfg model.Configuration) {
	os.Mkdir(cfg.ProjectName, DIR_PERM)

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

	_, err = exec.Command("go", "mod", "tidy").Output()
	if err != nil {
		panic("failed to run go mod tidy command: " + err.Error())
	}

	os.Exit(0)
}
