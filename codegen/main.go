package codegen

import (
	"fmt"

	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/model"
)

func GenerateMainContent(cfg model.Configuration) string {
	logImport := func() string {
		if cfg.Logging.Option == constant.Logrus {
			return `log "github.com/sirupsen/logrus"`
		}
		return ""
	}()

	// 0 - import
	// 1 - load env
	// 2 - port
	godotEnvCodes := func() []string {
		if cfg.GodotEnv {
			return []string{
				`"github.com/joho/godotenv"`,
				`err := godotenv.Load(".env")
				if err != nil {
					log.Fatal("Failed to load godotenv. Error: ", err)
					return err
				}`,
				`port := os.Getenv("PORT")`,
			}
		}
		return []string{"", "", "port := 5000"}
	}()
	return fmt.Sprintf(`
		package main

		import (
			"com.github.alissonbk/go-rest-template/app/router"
			"com.github.alissonbk/go-rest-template/config"
			%s
			%s
			"os"
		)

		func load() error {
			%s
			config.InitLog()
			return nil
		}

		func main() {
			err := load()
			if err != nil {
				log.Fatal("Failed to load application files. Error: ", err)
				return
			}
			
			%s
			app := router.Init()
			err = app.Run(":" + port)
			if err != nil {
				log.Fatal("Failed to startup application. Error: ", err)
				return
			}

			log.Info("app listening on port " + port)
		}

	`, logImport, godotEnvCodes[0], godotEnvCodes[1], godotEnvCodes[2])
}
