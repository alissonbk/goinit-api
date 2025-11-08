package codegen

func GenerateInjectionContent() string {
	return `
		package router

		import (
			"com.github.alissonbk/go-rest-template/app/controller"
			"com.github.alissonbk/go-rest-template/app/repository"
			"com.github.alissonbk/go-rest-template/app/service"
			"com.github.alissonbk/go-rest-template/config"
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
`
}
