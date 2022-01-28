package main

import (
	"os"

	li "github.com/aditya109/go-server-template/pkg/logwrapper"
	log "github.com/sirupsen/logrus"
)

func main() {
	standardLogger := li.NewLogger()
	standardLogger.Info("Hello from logger")
	pwd, _ := os.Getwd()
	log.Info(pwd)
	log.Info("Hello world !")

	standardLogger.Info()

}
