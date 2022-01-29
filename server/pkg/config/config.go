package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// GetConfiguration retrieves configuration from config file
func GetConfiguration() *Config {
	var config Config
	configFile, err := os.Open("./config/config.yaml")
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
