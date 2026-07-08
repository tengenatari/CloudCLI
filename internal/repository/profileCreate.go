package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"
)

func (y *YamlRepository) ProfileCreate(user, name, project string) error {

	profile, _ := LoadProfile(name)

	if profile != nil {
		return fmt.Errorf("profile %s already exists", name)
	}
	_, err := os.ReadFile(fmt.Sprintf("%s.yaml", name))

	if err == nil {
		return fmt.Errorf("yaml file in dir with name %s already exists, please replace or remove it", name)
	}

	profile = &models.Profile{User: user, Project: project}
	return SaveProfile(profile, name)
}
