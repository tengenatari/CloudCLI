package repository

import "CloudCLI/internal/models"

func (y *YamlRepository) ProfileCreate(user, name, project string) error {
	profile := models.Profile{User: user, Project: project}
	return SaveProfile(&profile, name)
}
