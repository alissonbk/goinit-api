package constant

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
		panic("invalid http library string")
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
