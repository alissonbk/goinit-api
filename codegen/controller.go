package codegen

import (
	"fmt"
	"github.com/alissonbk/goinit-api/model"
)

func GenerateControllerContent(cfg model.Configuration) string {
	return fmt.Sprintf(`
		package controller

		
		import (
			"%s/app/exception"
			"%s/app/service"
			"github.com/gin-gonic/gin"
		)

		type ExampleController struct {
			service *service.ExampleService
		}

		func NewExampleController(service *service.ExampleService) *ExampleController {
			return &ExampleController{service: service}
		}

		func (uc *ExampleController) GetAll(ctx *gin.Context) {
			defer exception.PanicHandler(ctx)
			ctx.JSON(200, uc.service.GetAll())
		}
	`, cfg.ModulePath, cfg.ModulePath)
}
