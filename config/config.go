package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Commands map[string]Command `yaml:"commands"`
	Args     map[string]Args    `yaml:"args"`
}

type Command struct {
	Name        string   `yaml:"name"`
	ArgNames    []string `yaml:"args"`
	Description string   `yaml:"description"`
}

type Args struct {
	Description string `yaml:"description"`
	ShortName   string `yaml:"shortName"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}
