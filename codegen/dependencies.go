package codegen

import (
	"fmt"

	"github.com/alissonbk/goinit-api/constant"
)

func GenereateDependenciesList(dbDriver constant.DatabaseDriver, httpLibrary constant.HttpLibrary,
	dbQueries constant.DatabaseQueries, godotenv bool, logs constant.LoggingOptions) string {

	httpLibDependency := func() string {
		switch httpLibrary {
		case constant.Echo:
			return "github.com/labstack/echo"
		case constant.Fiber:
			return "github.com/gofiber/fiber"
		case constant.Gin:
			return "github.com/gin-gonic/gin"
		default:
			panic("invalid library")
		}
	}()

	dbQueriesDeps := func() string {
		switch dbQueries {
		case constant.GORM:
			return fmt.Sprintf("gorm.io/driver/%s gorm.io/gorm", dbDriver.ToString())
		case constant.Sqlx:
			return "github.com/golang-migrate/migrate/v4 github.com/jmoiron/sqlx "
		default:
			panic("invalid database queries option")
		}
	}()

	dotenvDeps := func() string {
		if godotenv {
			return "github.com/joho/godotenv"
		}
		return ""
	}()

	logsDeps := func() string {
		switch logs {
		case constant.Logrus:
			return `
				github.com/antonfisher/nested-logrus-formatter
				github.com/sirupsen/logrus
			`
		case constant.Zap:
			return `github.com/uber-go/zap`
		default:
			panic("invalid log option")
		}
	}()

	return fmt.Sprintf(` 			
		%s
		%s		
		%s		
		%s		
		%s
		
	`, httpLibDependency, dbQueriesDeps, GetDatabaseDriverDependencies(dbDriver), dotenvDeps, logsDeps)
}
