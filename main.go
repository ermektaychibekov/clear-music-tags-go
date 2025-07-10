// main.go
package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Starting tag cleaner")

	config := loadConfig()
	processor := NewFileProcessor(config)
	processor.ProcessPaths()

	log.Println("Processing completed")
}

func loadConfig() *Config {
	args := os.Args[1:]
	configPath := "config.yaml"

	if len(args) > 0 && args[0] == "-c" && len(args) > 1 {
		configPath = args[1]
	}

	config, err := loadConfig(configPath)
	if err != nil {
		log.Println("Using default configuration")
		return defaultConfig()
	}
	return config
}
