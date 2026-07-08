package service

import "fmt"

func (service *ProfileService) ProfileList(args map[string]string) error {
	profiles, err := service.repository.ProfileList()
	if err != nil {
		return fmt.Errorf("error listing profiles: %v", err)
	}
	for profileName, profile := range profiles {
		printProfile(profileName, profile)
	}
	return nil
}
