package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type YamlRepository struct {
	dir string
}

func NewYamlRepository(dir string) *YamlRepository {
	if dir == "" {
		panic("dir cannot be empty")
	}
	return &YamlRepository{dir: dir}
}

func LoadProfile(filename string, dir string) (*models.Profile, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", dir, filename))

	if err != nil {
		return nil, fmt.Errorf("failed to find profile: %w", err)
	}

	var profile models.Profile

	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &profile, nil
}

func SaveProfile(profile *models.Profile, profileName string, dir string) error {

	data, err := yaml.Marshal(profile)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.yaml", dir, profileName), data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
