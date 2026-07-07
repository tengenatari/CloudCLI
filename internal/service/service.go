package service

import (
	"CloudCLI/internal/models"
	"errors"
	"fmt"
	"regexp"
)

type RepositoryInterface interface {
	ProfileGet(name string) (*models.Profile, error)
	ProfileDelete(name string) error
	ProfileList() (map[string]*models.Profile, error)
	ProfileCreate(user string, name string, project string) error
}

type ProfileService struct {
	repository RepositoryInterface
}

func NewProfileService(repositoryInterface RepositoryInterface) *ProfileService {
	return &ProfileService{repositoryInterface}
}
func validateParam(param string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9\s\-_]+$`, param)
	return matched
}

func getParam(paramName string, args map[string]string) (string, error) {
	param, ok := args[paramName]
	if !ok || !validateParam(param) {
		return "", fmt.Errorf("%s is required and may contain only letters, numbers, and hyphens", param)
	}
	return param, nil
}

func (service *ProfileService) ProfileCreate(args map[string]string) error {
	var name, project, user string

	name, err := getParam("name", args)
	if err != nil {
		return err
	}
	user, err = getParam("user", args)
	if err != nil {
		return err
	}
	project, err = getParam("project", args)
	if err != nil {
		return err
	}

	return service.repository.ProfileCreate(user, name, project)

}

func (service *ProfileService) ProfileDelete(args map[string]string) error {
	name, err := getParam("name", args)
	if err != nil {
		return err
	}
	err = service.repository.ProfileDelete(name)
	if err != nil {
		return fmt.Errorf("error deleting profile: no such profile")
	}
	return nil
}

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

func printProfile(name string, profile *models.Profile) {
	fmt.Printf("Profile: %s\n\tUser:    %s\n\tProject: %s\n", name, profile.User, profile.Project)
}
