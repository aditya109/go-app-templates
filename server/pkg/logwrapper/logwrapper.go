package logwrapper

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*log.Logger
}

// NewLogger returns a configured logger object
func NewLogger(env string) *StandardLogger {
	var baseLogger = log.New()
	var standardLogger = &StandardLogger{baseLogger}

	switch env {
	case "dev":
		standardLogger.Formatter = &log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		}
		standardLogger.Out = os.Stdout
	case "prod":
		standardLogger.Formatter = &log.JSONFormatter{}
		// standardLogger.Out = os.Stdout
	case "stag":
		standardLogger.Formatter = &log.TextFormatter{}
		standardLogger.Out = os.Stdout
	default:
		standardLogger.Formatter = &log.TextFormatter{}
		standardLogger.Out = os.Stdout
	}

	return standardLogger
}
