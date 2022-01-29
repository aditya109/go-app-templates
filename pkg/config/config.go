package config

import (
	"log"
	"os"

	helper "github.com/aditya109/go-server-template/pkg/helper"
	yaml "gopkg.in/yaml.v2"
)

// GetConfiguration retrieves configuration from config file
func GetConfiguration() *Config {
	var config Config
	var configFilePath string = helper.GetAbsolutePath("/config/config.yaml")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
