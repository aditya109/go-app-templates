package main

import (
	"github.com/aditya109/go-server-template/internal/server"
	logCfg "github.com/aditya109/go-server-template/pkg/logger"
)

func main() {
	logCfg.InitializeLogging() // initializing logger
	server.Start()
}
