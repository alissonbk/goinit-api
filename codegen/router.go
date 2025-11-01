package codegen

import (
	"fmt"

	"github.com/alissonbk/goinit-api/constant"
)

func generateRouterFileContent(lib constant.HttpLibrary) string {
	switch lib {
	case constant.Gin:
		return fmt.Sprintf(`
		package router

		import (
			"net/http"

			"github.com/gin-gonic/gin"
		)

		func Init() *gin.Engine {
			router := gin.New()
			router.Use(gin.Logger())
			router.Use(gin.Recovery())

			// DI
			injection := NewInjection()
			exampleController := injection.NewExampleController()

			router.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, map[string]string{"message": "hello"})
			})
			api := router.Group("/api/v1")
			{
				example := api.Group("/example")
				example.GET("", exampleController.GetAll)
			}

			return router
		}`)
	case constant.Echo:
		return fmt.Sprintf(`
			package router

			import (
				"net/http"

				"github.com/labstack/echo/v4"
				"github.com/labstack/echo/v4/middleware"
			)


			func Init() *echo.Echo {
				router := echo.New()
				router.Use(middleware.Logger())
				router.Use(middleware.Recover())

				// DI
				injection := NewInjection()
				exampleController := injection.NewExampleController()

				router.GET("", func(c echo.Context) error {
					return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
				})
				api := router.Group("/api/v1")

				api.GET("/ping", func(c echo.Context) error {
					return c.String(http.StatusOK, "pong")
				})

				example := api.Group("/example")

				example.GET("", exampleController.GetAll)

				return router
			}
		`)
	case constant.Fiber:
		return fmt.Sprintf(`
			package router

			import (
				"net/http"

				"github.com/gofiber/fiber/v3"
				"github.com/gofiber/fiber/v3/middleware/logger"
				"github.com/gofiber/fiber/v3/middleware/recover"
			)


			func Init() *fiber.App {
				router := fiber.New()
				router.Use(logger.New())
				router.Use(recover.New())

				// DI
				injection := NewInjection()
				exampleController := injection.NewExampleController()

				router.Get("", func(c fiber.Ctx) error {
					c.Status(http.StatusOK)
					return c.JSON(fiber.Map{"message": "hello"})
				})
				api := router.Group("/api/v1")

				api.Get("/ping", func(c fiber.Ctx) error {
					c.Status(http.StatusOK)
					return c.SendString("pong")
				})

				example := api.Group("/example")

				example.Get("", exampleController.GetAll)

				return router
			}
		`)
	default:
		panic("invalid http library")
	}

}
