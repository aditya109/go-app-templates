// Package sample http API
//
// Documentation of crud-template API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"net/http"
	"time"

	rt "github.com/aditya109/go-server-template/internal/router"
	cfg "github.com/aditya109/go-server-template/pkg/config"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
)

var (
	config       *cfg.Config
	logger       *log.StandardLogger
	environment  string
	httpPort     int64
	prefix       string
	endpoint     string
	writeTimeout time.Duration
	readTimeout  time.Duration
)

func main() {
	config = cfg.GetConfiguration()           // retrieving configuration
	logger = log.NewLogger(config.Server.Env) // initializing logger
	environment = config.Server.Env           // getting environment from config
	setHttpPortFromConfigObject()             // getting http port from config
	setEndpointFromConfigObject()             // getting endpoint from config
	setTimeoutsFromConfigObject()             // getting timeouts from config
	// configuring router for the server
	router := rt.ConfigureRouter()
	logger.Info("router configuration successful")
	logger.Info(fmt.Sprintf("starting server at %s://%s", prefix, endpoint))
	logger.Info(fmt.Sprintf("swagger docs can be viewed at %s://%s/docs", prefix, endpoint))
	// configuring server
	srv := &http.Server{
		Handler:      router,
		Addr:         endpoint,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	logger.Fatal(srv.ListenAndServe())
}

// setEndpointFromConfigObject sets endpoint IP from config object
func setEndpointFromConfigObject() {
	for _, v := range config.Server.Endpoints {
		if v.EnvType == environment {
			if v.TlsEnabled {
				prefix = "https"
			} else {
				prefix = "http"
			}
			endpoint = fmt.Sprintf("%s:%d", v.Name, httpPort)
		}
	}
}

// setHttpPortFromConfigObject sets httpPort variable to port mentioned in the config object
func setHttpPortFromConfigObject() {
	for _, v := range config.Server.Ports {
		if v.Name == "http-port" {
			httpPort = v.Port
		}
	}
}

// setTimeoutsFromConfigObject sets writeTimeout and readTimeout from config object
func setTimeoutsFromConfigObject() {
	for _, v := range config.Server.Timeouts {
		if v.Name == "write-timeout" {
			writeTimeout = time.Duration(v.Value) * time.Second
		}
		if v.Name == "read-timeout" {
			readTimeout = time.Duration(v.Value) * time.Second
		}
	}
}
