package main

import (
	cfg "github.com/aditya109/go-server-template/pkg/config"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

func main() {
	var config = cfg.GetConfiguration()
	var logger = log.NewLogger(config.Server.Env)

	logger.Info("Hello there ! Logger working ok !")
}
