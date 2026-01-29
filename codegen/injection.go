package codegen

import (
	"fmt"
	"github.com/alissonbk/goinit-api/model"
)

func GenerateInjectionContent(cfg model.Configuration) string {
	return fmt.Sprintf(`
		package router

		import (
			"%s/app/controller"
			"%s/app/repository"
			"%s/app/service"
			"%s/config"
			"github.com/jmoiron/sqlx"
		)

		// Injection is responsible for dependency injection for each route by returning a "Controller Object" ready to be used by the router
		// If you prefer you can use uber fx (https://github.com/uber-go/fx)

		type Injection struct {
			db *sqlx.DB
		}

		func NewInjection() *Injection {
			return &Injection{db: config.ConnectDB()}
		}

		func (i *Injection) NewExampleController() *controller.ExampleController {
			r := repository.NewExampleRepository(i.db)
			s := service.NewExampleService(r)
			return controller.NewExampleController(s)
		}
`, cfg.ModulePath, cfg.ModulePath, cfg.ModulePath, cfg.ModulePath)
}
