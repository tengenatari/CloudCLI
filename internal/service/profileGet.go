package service

import "errors"

func (service *ProfileService) ProfileGet(args map[string]string) error {

	name, err := getParam("name", args)

	if err != nil {
		return err
	}
	profile, err := service.repository.ProfileGet(name)

	if err != nil {
		return errors.New("profile not found")
	}

	printProfile(name, profile)
	return nil
}
