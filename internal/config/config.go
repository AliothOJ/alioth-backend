package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Microservice struct {
	Name   string `yaml:name`
	Slaves []struct {
		Host     string `yaml:"host"`
		HTTPPort string `yaml:"HTTPPort"`
		GRPCPort string `yaml:"gRPCPort"`
	} `yaml:slaves`
}
type ServiceConfig struct {
	AppName string `yaml:"appName"`
	Server  struct {
		Host string `yaml:host`
		Port string `yaml:port`
	} `yaml:server`
	Microservices []Microservice `yaml:"microservices"`
}

var globalConfig = new(ServiceConfig)

// CheckExists check if path exists and is a file
func CheckExists(configPath string) error {
	stat, err := os.Stat(configPath)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' is not a yaml file but a directory", configPath)
	}
	return nil
}

// LoadConfig load config from specified path
func LoadConfig(configPath string) (*ServiceConfig, error) {
	conf, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer conf.Close()
	// decode
	d := yaml.NewDecoder(conf)
	if err := d.Decode(globalConfig); err != nil {
		return nil, err
	}
	return globalConfig, nil
}
