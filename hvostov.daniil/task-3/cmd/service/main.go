package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	if path == "" {
		return config, fmt.Errorf("Use (--config or -config) to indicate configuration path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("parse config %q: %w", path, err)
	}

	if config.InputFile == "" {
		return config, fmt.Errorf("config %q: input-file is required", path)
	}

	if config.OutputFile == "" {
		return config, fmt.Errorf("config %q: output-file is required", path)
	}

	return config, nil
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
