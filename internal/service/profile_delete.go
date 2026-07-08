package service

import "fmt"

func (service *ProfileService) ProfileDelete(args map[string]string) error {
	name, err := getParam("name", args)
	if err != nil {
		return err
	}
	err = service.repository.ProfileDelete(name)
	if err != nil {
		return fmt.Errorf("error deleting profile: no such profile: %v", err)
	}
	return nil
}
