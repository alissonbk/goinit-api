package codegen

import (
	"fmt"

	"github.com/alissonbk/goinit-api/constant"
)

func GenerateDatabaseContent(databaseDriver constant.DatabaseDriver, databaseQueries constant.DatabaseQueries, godotenv bool, logLevel constant.LogLevel) string {

	dsnInfo := func() string {
		if godotenv {
			return `				
				dsn := os.Getenv("DB_DSN")
				maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
				if err != nil {
					logrus.Fatal("ENV DB_MAX_OPEN_CONN should be an integer. Error: ", err)
				}
				maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
				if err != nil {
					logrus.Fatal("ENV DB_MAX_IDLE_CONN should be an integer. Error: ", err)
				}
			`
		}
		return `
			dsn := "host=localhost user=postgres password=1234 dbname=go_rest_template port=5432 sslmode=disable"
			maxOpen := 20
			maxIdle := 10
		`
	}()

	logLevelInfo := func() string {
		if godotenv {
			return `os.Getenv("DB_LOG_LEVEL")`
		}

		return logLevel.ToString()
	}()

	importMigrateDatabaseDriver := func() string {
		if databaseDriver == constant.PGX {
			return "github.com/golang-migrate/migrate/v4/database/pgx/v5"
		}

		if databaseDriver == constant.Mssql {
			return "github.com/golang-migrate/migrate/v4/database/sqlserver"
		}

		return fmt.Sprintf("github.com/golang-migrate/migrate/v4/database/%s", databaseDriver.ToString())
	}()

	if databaseQueries == constant.Sqlx {
		return fmt.Sprintf(`
			package config

			import (
				"fmt"
				"strings"

				"github.com/golang-migrate/migrate/v4"
				"%[3]s"	
				_ "github.com/golang-migrate/migrate/v4/source/file"
				"github.com/jmoiron/sqlx"
				_ %[4]s
				"github.com/sirupsen/logrus"

				"os"
				"strconv"
				"time"
			)

			type DatabaseInformation struct {
				DataSourceName  string
				MaxConnOpenned  int
				MaxConnIddle    int
				MaxConnLifetime time.Duration
			}

			func ConnectDB() *sqlx.DB {
				var err error

				dbInfo := GetDatabaseInformation()

				db, err := sqlx.Open("%[1]s", dbInfo.DataSourceName)
				if err != nil {
					logrus.Error("error while connecting to database, Error: ", err)
				}

				db.SetMaxOpenConns(dbInfo.MaxConnOpenned)
				db.SetMaxIdleConns(dbInfo.MaxConnIddle)
				db.SetConnMaxLifetime(dbInfo.MaxConnLifetime)

				runMigrations(db, dbInfo.DataSourceName)

				return db
			}

			func GetDatabaseInformation() *DatabaseInformation {
				%[2]s

				return &DatabaseInformation{
					DataSourceName:  dsn,
					MaxConnOpenned:  maxOpen,
					MaxConnIddle:    maxIdle,
					MaxConnLifetime: time.Hour,
				}
			}

			func dbNameFromDataSource(dsn string) string {
				return strings.Split(" ", strings.Split("dbname=", dsn)[0])[0]
			}

			func runMigrations(db *sqlx.DB, dsn string) {
				workingDir, err := os.Getwd()
				if err != nil {
					panic("could not get current directory, " + err.Error())
				}

				driver, err := %[1]s.WithInstance(db.DB, &%[1]s.Config{})
				if err != nil {
					panic("failed to load %[1]s driver, " + err.Error())
				}	

				m, err := migrate.NewWithDatabaseInstance(
					fmt.Sprintf("file://%%s/config/migrations", workingDir),
					dbNameFromDataSource(dsn),
					driver,
				)
				if err != nil {
					panic("failed to setup migrations config, " + err.Error())
				}

				err = m.Up()
				if err != nil {
					if err.Error() == "no change" {
						return
					}
					panic("failed to run migrations, " + err.Error())
				}

				logrus.Info("migrations ran sucessfully")
			}

		`, databaseDriver.ToString(), dsnInfo, importMigrateDatabaseDriver, GetDatabaseDriverDependencies(databaseDriver))
	}

	if databaseQueries == constant.GORM {
		return fmt.Sprintf(`
			package config

			import (
				"github.com/sirupsen/logrus"
				"gorm.io/driver/%[1]s"
				"gorm.io/gorm"
				gormlogger "gorm.io/gorm/logger"
				"log"
				"os"
				"strconv"
				"strings"
				"time"
			)

			func ConnectDB() *gorm.DB {
				%[2]s

				db, err := gorm.Open(%[1]s.Open(dsn), &gorm.Config{
					Logger: createDBLogger(),
				})
				if err != nil {
					log.Fatal("Error while connecting to database, Error: ", err)
				}

				sqlDB, err := db.DB()
				if err != nil {
					log.Fatal("Error while acquiring sql.DB from gorm lib, Error: ", err)
				}
				sqlDB.SetMaxOpenConns(maxOpen)
				sqlDB.SetMaxIdleConns(maxIdle)
				sqlDB.SetConnMaxLifetime(time.Hour)

				return db
			}

			func createDBLogger() gormlogger.Interface {
				var logLevel gormlogger.LogLevel
				var ignoreNotFound bool
				var parameterizedQueries bool
				var colorful bool
				switch strings.ToUpper(%[3]s) {
				case "PROD":
					logLevel = gormlogger.Warn
					ignoreNotFound = true
					parameterizedQueries = true
					colorful = false
				case "DEV":
					logLevel = gormlogger.Info
					ignoreNotFound = false
					parameterizedQueries = false
					colorful = true
				case "INFO":
					logLevel = gormlogger.Info
					ignoreNotFound = false
					parameterizedQueries = false
					colorful = true
				case "WARN":
					logLevel = gormlogger.Warn
					ignoreNotFound = false
					parameterizedQueries = false
				case "ERROR":
					logLevel = gormlogger.Error
					ignoreNotFound = true
					parameterizedQueries = false
				case "SILENT":
					logLevel = gormlogger.Silent
					ignoreNotFound = true
					parameterizedQueries = true
				default:
					logLevel = gormlogger.Error
					ignoreNotFound = true
					parameterizedQueries = true
				}

				return gormlogger.New(
					log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
					gormlogger.Config{
						SlowThreshold:             time.Second,
						LogLevel:                  logLevel,
						IgnoreRecordNotFoundError: ignoreNotFound,
						ParameterizedQueries:      parameterizedQueries,
						Colorful:                  colorful,
					},
				)
			}

		`, databaseDriver.ToString(), dsnInfo, logLevelInfo)
	}

	panic("invalid database queries option")
}

func GetDatabaseDriverDependencies(dbDriver constant.DatabaseDriver) string {
	switch dbDriver {
	case constant.Clickhouse:
		return "github.com/ClickHouse/clickhouse-go"
	case constant.PGX:
		return "github.com/jackc/pgx"
	case constant.MySQL:
		return "github.com/go-sql-driver/mysql"
	case constant.Mssql:
		return "github.com/denisenkom/go-mssqldb"
	case constant.Postgres:
		return "github.com/lib/pq"
	case constant.Sqlite3:
		return "github.com/mattn/go-sqlite3"
	default:
		panic("invalid db driver")
	}
}
