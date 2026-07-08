package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type YamlRepository struct {
}

func NewYamlRepository() *YamlRepository {
	return &YamlRepository{}
}

func LoadProfile(filename string) (*models.Profile, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s.yaml", filename))

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

func SaveProfile(profile *models.Profile, profileName string) error {

	data, err := yaml.Marshal(profile)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("%s.yaml", profileName), data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
