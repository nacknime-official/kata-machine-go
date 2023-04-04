package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const DefaultConfigPath = "config.yaml"

type Config struct {
	DSA []string `yaml:"DSA"`
}

func Load(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &config, nil
}
