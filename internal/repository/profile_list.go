package repository

import (
	"CloudCLI/internal/models"
	"fmt"
	"os"
	"strings"
)

func (y *YamlRepository) ProfileList() (map[string]*models.Profile, error) {
	dir, err := os.ReadDir(".")

	if err != nil {
		return nil, err
	}

	profiles := make(map[string]*models.Profile)
	for _, file := range dir {
		name, profile, err := getProfileFromFile(file)
		if err != nil {
			continue
		}
		profiles[name] = profile
	}

	return profiles, nil
}

func getProfileFromFile(file os.DirEntry) (string, *models.Profile, error) {

	if !strings.HasSuffix(file.Name(), ".yaml") {
		return "", nil, fmt.Errorf("it isnt yaml")
	}
	name := strings.TrimSuffix(file.Name(), ".yaml")

	profile, err := LoadProfile(name)

	if err != nil {
		return "", nil, err
	}
	return name, profile, nil
}
