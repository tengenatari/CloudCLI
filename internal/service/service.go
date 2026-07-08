package service

import (
	"CloudCLI/internal/models"
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

func printProfile(name string, profile *models.Profile) {
	fmt.Printf("Profile: %s\n\tUser:    %s\n\tProject: %s\n", name, profile.User, profile.Project)
}
