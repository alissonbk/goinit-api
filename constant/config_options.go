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
	Oci8
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
	case Oci8:
		return "oci8"
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
	PROD LogLevel = iota
	DEV
	INFO
	WARN
	ERROR
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
	case PROD:
		return "PROD"
	case DEV:
		return "DEV"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case SILENT:
		return "SILENT"
	default:
		return "UNKNOWN"
	}
}
