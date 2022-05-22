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

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aditya109/go-server-template/internal/constants"
	"github.com/aditya109/go-server-template/internal/models"
	rt "github.com/aditya109/go-server-template/internal/router"
	logger "github.com/sirupsen/logrus"
)

var (
	configuration *models.Config
	httpPort      string
	prefix        string
	endpoint      string
	writeTimeout  time.Duration
	readTimeout   time.Duration
	envs          models.Envs
)

func Start(config *models.Config) {
	configuration = config
	getApplicableEnvironmentVariablesFromConfig() // getting applicable environment variables from config
	setHTTPPortFromConfigObject()                 // getting http port from config
	setEndpointFromConfigObject()                 // getting endpoint from config
	setTimeoutsFromConfigObject()                 // getting timeouts from config

	// configuring router for the server
	router := rt.ConfigureRouter()
	logger.Info("router configuration successful")
	logger.Info(fmt.Sprintf("starting server at %s://%s", prefix, endpoint))
	logger.Info(fmt.Sprintf("swagger docs can be viewed at %s://%s/docs", prefix, endpoint))

	// configuring server
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", httpPort),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	logger.Fatal(srv.ListenAndServe())
}

// getApplicableEnvironmentVariablesFromConfig gets applicable environment variables from configuration
func getApplicableEnvironmentVariablesFromConfig() {
	switch configuration.ServerConfig.EnvironmentType {
	case constants.DEV:
		envs = configuration.ServerConfig.DevEnvs
	case constants.STAGING:
		envs = configuration.ServerConfig.StagEnvs
	case constants.PRODUCTION:
		envs = configuration.ServerConfig.ProdEnvs
	default:
		envs = configuration.ServerConfig.DevEnvs
	}
}

// setHTTPPortFromConfigObject sets httpPort variable to port mentioned in the config object
func setHTTPPortFromConfigObject() {
	httpPort = envs.ServerPort
}

// setEndpointFromConfigObject sets endpoint IP from config object
func setEndpointFromConfigObject() {
	if envs.IsTLSEnabled {
		prefix = "https"
	} else {
		prefix = "http"
	}
	endpoint = fmt.Sprintf("%s:%s", "localhost", httpPort)
}

// setTimeoutsFromConfigObject sets writeTimeout and readTimeout from config object
func setTimeoutsFromConfigObject() {
	writeTimeout = time.Duration(envs.WriteTimeout) * time.Second
	readTimeout = time.Duration(envs.ReadTimeout) * time.Second
}
