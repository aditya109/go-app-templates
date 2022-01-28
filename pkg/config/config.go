package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Env string `yaml:"environment"`
	}
}

func GetConfiguration() *Config {
	var config Config
	configFile, err := os.Open("./config/config.yaml")
	defer configFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
