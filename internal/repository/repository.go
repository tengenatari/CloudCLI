package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type YamlRepository struct {
}

func (y *YamlRepository) ProfileCreate(user, name, project string) error {
	profileInner := models.Profile{User: user, Project: project}
	profile := make(map[string]models.Profile)
	profile[name] = profileInner
	return SaveProfile(&profile, name)
}

func (y *YamlRepository) ProfileDelete(name string) error {
	return nil
}

func (y *YamlRepository) ProfileList() error {
	return nil
}

func (y *YamlRepository) ProfileGet(name string) (*map[string]models.Profile, error) {
	return LoadProfile(name)
}

func LoadProfile(filename string) (*map[string]models.Profile, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s.yaml", filename))

	if err != nil {
		return nil, fmt.Errorf("failed to find profile: %w", err)
	}

	var profile map[string]models.Profile

	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &profile, nil
}

func SaveProfile(profile *map[string]models.Profile, profileName string) error {

	data, err := yaml.Marshal(profile)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// 2. Записываем в файл
	err = os.WriteFile(profileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
