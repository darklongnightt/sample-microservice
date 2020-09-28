package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config defines struct for all app configs
type Config struct {
	DB Database `yaml:"database"`
}

// Database defines struct for db config
type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// GetAppConfig parses config from file
func GetAppConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %v", err)
	}

	var config *Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %v", err)
	}

	return config, nil
}
