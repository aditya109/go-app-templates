package main

import (
	"fmt"
	"net/http"
	"time"

	cfg "github.com/aditya109/go-server-template/pkg/config"
	log "github.com/aditya109/go-server-template/pkg/logwrapper"
	"github.com/gorilla/mux"
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

func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "welcome to server ⚡ !")
}

func main() {
	// retrieving configuration
	config = cfg.GetConfiguration()
	// initializing logger
	logger = log.NewLogger(config.Server.Env)

	// getting environment from config
	environment = config.Server.Env

	// getting http port from config
	for _, v := range config.Server.Ports {
		if v.Name == "http-port" {
			httpPort = v.Port
		}
	}

	// getting endpoint from config
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

	// getting timeouts from config
	for _, v := range config.Server.Timeouts {
		if v.Name == "write-timeout" {
			writeTimeout = time.Duration(v.Value) * time.Second
		}
		if v.Name == "read-timeout" {
			readTimeout = time.Duration(v.Value) * time.Second
		}
	}

	// configuring router for the server
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler)
	logger.Info("router configuration successful")

	logger.Info(fmt.Sprintf("starting server at %s://%s", prefix, endpoint))

	// configuring server
	srv := &http.Server{
		Handler:      router,
		Addr:         endpoint,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	logger.Fatal(srv.ListenAndServe())
}
