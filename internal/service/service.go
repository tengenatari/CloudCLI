package service

import "CloudCLI/internal/models"

type RepositoryInterface interface {
	ProfileGet(name string) (*map[string]models.Profile, error)
	ProfileDelete(name string) error
	ProfileList() error
	ProfileCreate(user, name, project string) error
}

type ProfileService struct {
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (service *ProfileService) ProfileCreate(args map[string]string) error {
	return nil
}

func (service *ProfileService) ProfileDelete(args map[string]string) error {
	return nil
}
func (service *ProfileService) ProfileList(args map[string]string) error {
	return nil
}
func (service *ProfileService) ProfileGet(args map[string]string) error {
	return nil
}

func (service *ProfileService) Help(args map[string]string) error {
	return nil
}
