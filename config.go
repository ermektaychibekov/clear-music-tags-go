// config.go
package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RemoveStrings []string `yaml:"remove_strings"`
	Paths         []string `yaml:"paths"`
}

func defaultConfig() *Config {
	return &Config{
		RemoveStrings: []string{
			"https://djsoundtop.com",
			"https://electronicfresh.com",
			"djsoundtop.com",
			"electronicfresh.com",
		},
		Paths: []string{"."},
	}
}

func loadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
