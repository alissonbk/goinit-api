package codegen

import (
	"fmt"

	"github.com/alissonbk/goinit-api/constant"
	"github.com/alissonbk/goinit-api/model"
)

func GenerateLogsContent(cfg model.Configuration) string {
	if cfg.Logging.Option == constant.Logrus {
		return generateLogrusLogsContent(cfg.Logging.Loglevel, cfg.Logging.Structured)
	}

	if cfg.Logging.Option == constant.Zap {
		return generateZapLogsContent(cfg.Logging.Loglevel, cfg.Logging.Structured)
	}

	panic("invalid log option")
}

func generateLogrusLogsContent(logLevel constant.LogLevel, structured bool) string {

	formatter := func() string {
		if structured {
			return `
				log.SetFormatter(&log.JSONFormatter{
					TimestampFormat: "2006-01-02 15:04:05",
				})
			`
		}

		return `
			log.SetFormatter(&nested.Formatter{
					HideKeys:        true,
					FieldsOrder:     []string{"component", "category"},
					TimestampFormat: "2006-01-02 15:04:05",
					ShowFullLevel:   true,
					CallerFirst:     true,
				})`
	}()
	return fmt.Sprintf(`
		package config

		import (
			"github.com/sirupsen/logrus"
			log "github.com/sirupsen/logrus"
		)


		func InitLog() {
			lvl, err := logrus.ParseLevel(%s)
			if err != nil {
				panic("failed to parse logrus log level" + err.Error())
			}

			log.SetLevel(lvl)
			log.SetReportCaller(true)
			%s

		}
	`, logLevel.ToString(), formatter)
}

func generateZapLogsContent(logLevel constant.LogLevel, structured bool) string {
	level := func() string {
		switch logLevel {
		case constant.DEBUG:
			return "zap.NewAtomicLevelAt(zap.DebugLevel)"
		case constant.INFO:
			return "zap.NewAtomicLevelAt(zap.InfoLevel)"
		case constant.WARN:
			return "zap.NewAtomicLevelAt(zap.WarnLevel)"
		case constant.ERROR:
			return "zap.NewAtomicLevelAt(zap.ErrorLevel)"
		case constant.FATAL:
			return "zap.NewAtomicLevelAt(zap.FatalLevel)"
		default:
			return "zap.NewAtomicLevelAt(zap.InfoLevel)"
		}
	}()
	encoding := func() string {
		if structured {
			return "json"
		}

		return "console"
	}()
	// zap structured
	return fmt.Sprintf(`
		package config

		import (
			"sync"

			"go.uber.org/zap"
			"go.uber.org/zap/zapcore"
		)

		var (
			logger *zap.Logger
			once   sync.Once
		)

		func InitLog() {
			once.Do(func() {
				config := zap.Config{
					Level:       %s,
					Development: false,
					Encoding:    "%s",
					EncoderConfig: zapcore.EncoderConfig{
						TimeKey:        "timestamp",
						LevelKey:       "level",
						NameKey:        "logger",
						CallerKey:      "caller",
						MessageKey:     "message",
						StacktraceKey:  "stacktrace",
						LineEnding:     zapcore.DefaultLineEnding,
						EncodeLevel:    zapcore.LowercaseLevelEncoder,
						EncodeTime:     zapcore.ISO8601TimeEncoder,
						EncodeDuration: zapcore.SecondsDurationEncoder,
						EncodeCaller:   zapcore.ShortCallerEncoder,
					},
					OutputPaths: []string{
						"stdout",
						//"/var/log/myapp/app.log"
					},
					ErrorOutputPaths: []string{"stderr"},
				}
				logger, err := config.Build()
				if err != nil {
					panic("failed to initialize zap logger, cause: " + err.Error())
				}
				// flushes buffer, if any
				logger.Sync()
			})
		}

		func Logger() *zap.Logger {
			if logger == nil {
				panic("the logger have not been initialized yet!")
			}

			return logger
		}
	`, level, encoding)
}
