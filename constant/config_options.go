package constant

import "strings"

type HttpLibrary uint8

const (
	Gin HttpLibrary = iota
	Fiber
	Echo
)

func (hl HttpLibrary) ToString() string {
	switch hl {
	case Gin:
		return "Gin"
	case Fiber:
		return "Fiber"
	case Echo:
		return "Echo"
	default:
		return "invalid type"
	}
}

func HttpLibraryFromString(s string) HttpLibrary {
	switch s {
	case "Gin":
		return Gin
	case "Fiber":
		return Fiber
	case "Echo":
		return Echo
	default:
		panic("invalid http library string " + s)
	}
}

type ProjectStructure uint8

const (
	MVC ProjectStructure = iota
	Hexagonal
)

func (hl ProjectStructure) ToString() string {
	switch hl {
	case MVC:
		return "MVC"
	case Hexagonal:
		return "Hexagonal"
	default:
		return "unknown"
	}
}

func ProjectStructureFromString(s string) ProjectStructure {
	switch s {
	case "MVC":
		return MVC
	case "Hexagonal":
		return Hexagonal
	default:
		panic("invalid string for project structure " + s)
	}
}

type DatabaseQueries uint8

const (
	GORM DatabaseQueries = iota
	Sqlx
)

func (dq DatabaseQueries) ToString() string {
	switch dq {
	case GORM:
		return "GORM"
	case Sqlx:
		return "Sqlx"
	default:
		return "unknown"
	}
}

func DatabaseQueriesFromString(s string) DatabaseQueries {
	switch s {
	case "GORM":
		return GORM
	case "Sqlx":
		return Sqlx
	default:
		panic("invalid database query " + s)
	}
}

type DatabaseDriver uint8

const (
	Postgres DatabaseDriver = iota
	PGX
	MySQL
	Sqlite3
	Mssql
	Clickhouse
	_endDatabaseDriver
)

func (d DatabaseDriver) ToString() string {
	switch d {
	case Postgres:
		return "postgres"
	case PGX:
		return "pgx"
	case MySQL:
		return "mysql"
	case Sqlite3:
		return "sqlite3"
	case Mssql:
		return "mssql"
	case Clickhouse:
		return "clickhouse"
	default:
		return "unknown"
	}
}

func DatabaseDriverFromString(s string) DatabaseDriver {
	switch s {
	case "postgres":
		return Postgres
	case "pgx":
		return PGX
	case "mysql":
		return MySQL
	case "sqlite3":
		return Sqlite3
	case "mssql":
		return Mssql
	case "clickhouse":
		return Clickhouse
	default:
		panic("invalid database driver string " + s)
	}
}

func AllDatabaseDrivers() []string {
	lst := make([]string, _endDatabaseDriver)
	for i := range _endDatabaseDriver {
		lst[i] = DatabaseDriver(i).ToString()
	}

	return lst
}

type LoggingOptions uint8

const (
	Zap LoggingOptions = iota
	Logrus
)

func (l LoggingOptions) ToString() string {
	switch l {
	case Zap:
		return "zap"
	case Logrus:
		return "logrus"
	default:
		return "unknown"
	}
}

func LoggingOptionsFromString(s string) LoggingOptions {
	switch s {
	case "zap":
		return Zap
	case "logrus":
		return Logrus
	default:
		panic("invalid logging options from string " + s)
	}

}

func LoggingNestedFromString(s string) bool {
	if strings.ToLower(s) == "nested" {
		return true
	}
	return false
}

func BooleanOptionFromString(s string) bool {
	if strings.ToLower(s) == "yes" {
		return true
	}

	return false
}

type LogLevel uint8

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	SILENT
	_endLogLevel
)

func AllLogLevels() []string {
	lst := make([]string, _endLogLevel)

	for i := range _endLogLevel {
		lst[i] = LogLevel(i).ToString()
	}

	return lst
}

func (l LogLevel) ToString() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case SILENT:
		return "SILENT"
	default:
		return "UNKNOWN"
	}
}

func LogLevelFromString(s string) LogLevel {
	switch s {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	case "SILENT":
		return SILENT
	default:
		panic("invalid loglevel string")
	}
}
