package repository

import "CloudCLI/internal/models"

func (y *YamlRepository) ProfileGet(name string) (*models.Profile, error) {
	return LoadProfile(name, y.dir)
}
