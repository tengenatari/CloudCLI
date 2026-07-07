package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v4"
)

type YamlRepository struct {
}

func NewYamlRepository() *YamlRepository {
	return &YamlRepository{}
}

func (y *YamlRepository) ProfileCreate(user, name, project string) error {
	profile := models.Profile{User: user, Project: project}
	return SaveProfile(&profile, name)
}

func (y *YamlRepository) ProfileDelete(name string) error {
	err := os.Remove(fmt.Sprintf("%s.yaml", name))
	if err != nil {
		return err
	}
	return nil
}

func (y *YamlRepository) ProfileList() (map[string]*models.Profile, error) {
	dir, err := os.ReadDir(".")

	if err != nil {
		return nil, err
	}

	profiles := make(map[string]*models.Profile)
	for _, file := range dir {
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}
		name := strings.TrimSuffix(file.Name(), ".yaml")
		fmt.Printf("Loading profile %s\n", name)
		profile, err := LoadProfile(name)
		if err != nil {
			return nil, err
		}
		profiles[name] = profile
	}

	return profiles, nil
}

func (y *YamlRepository) ProfileGet(name string) (*models.Profile, error) {
	return LoadProfile(name)
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
