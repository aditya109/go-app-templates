package logwrapper

import (
	"fmt"
	"io"
	"os"

	cfg "github.com/aditya109/go-server-template/pkg/config"
	log "github.com/sirupsen/logrus"
)

var (
	config *cfg.Config
)

// InitializeLogging returns a configured logger object
func InitializeLogging() {
	config = cfg.GetConfiguration()
	env := config.Server.Env
	switch env {
	case "dev":
		f, err := os.OpenFile("baremetrix.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		// don't forget to close it
		// defer f.Close()

		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	case "prod":
		log.SetFormatter(&log.JSONFormatter{})
	case "stag":
		log.SetFormatter(&log.TextFormatter{})
		log.SetOutput(os.Stdout)
	default:
		log.SetFormatter(&log.TextFormatter{})
		log.SetOutput(os.Stdout)
	}
}
