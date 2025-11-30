package filegen

import (
	"os"
	"os/exec"

	"github.com/alissonbk/goinit-api/codegen"
	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/model"
)

// TODO return error message instead of panicking
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
		if mainoutput, err := exec.Command("echo", codegen.GenerateMainContent(cfg), ">", "main.go").Output(); err != nil {
			panic("failed to write main.go, output: " + string(mainoutput))
		}

		// dockerfile and compose
		if cfg.Dockerfile {
			if dockerfileoutput, err := exec.Command("echo", codegen.GenerateDockerfileContent(), ">", "Dockerfile").Output(); err != nil {
				panic("failed to write dockerfile content, output: " + string(dockerfileoutput))
			}
			if composeoutput, err := exec.Command("echo", codegen.GenerateDockerComposeContent(), ">", "docker-compose.yaml").Output(); err != nil {
				panic("failed to write docker-compose content, output: " + string(composeoutput))
			}
		}

		if cfg.GodotEnv {
			echocmd := exec.Command("echo", codegen.GenerateEnvContent())
			teecmd := exec.Command("tee", ".env", ".env.example")

			stdoutpipe, err := echocmd.StdoutPipe()
			if err != nil {
				panic("failed to get stdout pipe from echo, cause: " + err.Error())
			}
			defer stdoutpipe.Close()

			teecmd.Stdin = stdoutpipe

			if err = echocmd.Start(); err != nil {
				panic("failed to write .env echo cmd: " + err.Error())
			}

			if output, err := teecmd.CombinedOutput(); err != nil {
				panic("failed to run tee cmd creating .env output:" + string(output))
			}

			// maybe should call echocmd.Wait() here or before

		}

		// config dir
		os.Mkdir("config", 0755)
		if err = os.Chdir("config"); err != nil {
			panic("failed to change directory: " + err.Error())
		}
		if databaseContentOutput, err := exec.Command("echo", codegen.GenerateDatabaseContent(cfg), ">", "database.go").Output(); err != nil {
			panic("failed to create database content output:" + string(databaseContentOutput))
		}
		if logsOutput, err := exec.Command("echo", codegen.GenerateLogsContent(cfg), ">", "logs.go").Output(); err != nil {
			panic("failed to create logs content output:" + string(logsOutput))
		}

		// migrations
		os.Mkdir("migrations", 0755)
		if migrationUpOutput, err := exec.Command("echo", codegen.GenerateExampleMigrationUpContent(), ">", "migrations/000001_create_example_table.up.sql").Output(); err != nil {
			panic("failed to create migration up content output:" + string(migrationUpOutput))
		}
		if migrationDownOutput, err := exec.Command("echo", codegen.GenerateExampleMigrationDownContent(), ">", "migrations/000001_create_example_table.down.sql").Output(); err != nil {
			panic("failed to create migration down content output:" + string(migrationDownOutput))
		}

		// back to main directory
		err = os.Chdir("../../")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// app dir
		os.Mkdir("app", 0755)
		err = os.Chdir("app")
		if err != nil {
			panic("failed to change directory: " + err.Error())
		}

		// constant
		os.Mkdir("constant", 0755)
		if constantOutput, err := exec.Command("echo", codegen.GenerateConstantContent(), ">", "constant/status.go").Output(); err != nil {
			panic("failed to create constant content output:" + string(constantOutput))
		}

		// controller
		os.Mkdir("controller", 0755)
		if controllerOutput, err := exec.Command("echo", codegen.GenerateControllerContent(), ">", "controller/example.go").Output(); err != nil {
			panic("failed to create controller content output:" + string(controllerOutput))
		}

		// exeption
		os.Mkdir("exception", 0755)
		if panicOutput, err := exec.Command("echo", codegen.GeneratePanicContent(cfg.Logging.Option), ">", "exception/panic.go").Output(); err != nil {
			panic("failed to create panic content output:" + string(panicOutput))
		}

		// model
		os.Mkdir("model", 0755)
		err = os.Chdir("model")
		os.Mkdir("entity", 0755)
		os.Mkdir("dto", 0755)
		if baseEntityOutput, err := exec.Command("echo", codegen.GenerateBaseEntity(), ">", "entity/base.go").Output(); err != nil {
			panic("failed to create base entity content output:" + string(baseEntityOutput))
		}
		if exampleEntityOutput, err := exec.Command("echo", codegen.GenerateExampleEntity(), ">", "entity/example.go").Output(); err != nil {
			panic("failed to create entity example content output:" + string(exampleEntityOutput))
		}
		err = os.Chdir("../")

		// repository
		os.Mkdir("repository", 0755)
		if repositoryOutput, err := exec.Command("echo", codegen.GenerateRepositoryContent(cfg.DatabaseQueries), ">", "repository/example.go").Output(); err != nil {
			panic("failed to create repository content output:" + string(repositoryOutput))
		}

		// router
		os.Mkdir("router", 0755)
		if routerOutput, err := exec.Command("echo", codegen.GenerateRouterContent(cfg.HttpLibrary), ">", "router/routes.go").Output(); err != nil {
			panic("failed to create router content output:" + string(routerOutput))
		}
		if injectionOutput, err := exec.Command("echo", codegen.GenerateInjectionContent(), ">", "router/injection.go").Output(); err != nil {
			panic("failed to create injection content output:" + string(injectionOutput))
		}

		// service
		os.Mkdir("service", 0755)
		if serviceOutput, err := exec.Command("echo", codegen.GenerateServiceContent(), ">", "service/example.go").Output(); err != nil {
			panic("failed to create service content output:" + string(serviceOutput))
		}

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
	tidyOutput, err := tidyCmd.Output()
	if err != nil {
		panic("failed to run go mod tidy command : " + string(tidyOutput))
	}

	os.Exit(0)
}
