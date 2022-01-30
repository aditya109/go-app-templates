package logwrapper

import (
	"os"

	cfg "github.com/aditya109/go-server-template/pkg/config"
	log "github.com/sirupsen/logrus"
)

var (
	config *cfg.Config
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*log.Logger
}

// NewLogger returns a configured logger object
func NewLogger(env ...string) *StandardLogger {
	if len(env) == 0 {
		config = cfg.GetConfiguration()
		env = []string{config.Server.Env}
	}
	var baseLogger = log.New()
	var standardLogger = &StandardLogger{baseLogger}
	switch env[0] {
		case "dev":
			standardLogger.Formatter = &log.TextFormatter{
				DisableColors: false,
				FullTimestamp: true,
			}
			standardLogger.Out = os.Stdout
		case "prod":
			standardLogger.Formatter = &log.JSONFormatter{}
		case "stag":
			standardLogger.Formatter = &log.TextFormatter{}
			standardLogger.Out = os.Stdout
		default:
			standardLogger.Formatter = &log.TextFormatter{}
			standardLogger.Out = os.Stdout
	}
	return standardLogger
}
