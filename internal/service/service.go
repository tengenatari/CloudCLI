package service

type ProfileService struct {
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (service *ProfileService) ProfileCreate(name string, user string, profile string) error {
	return nil
}

func (service *ProfileService) ProfileDelete(name string) error {
	return nil
}
func (service *ProfileService) ProfileList() error {
	return nil
}
func (service *ProfileService) ProfileGet(name string) error {
	return nil
}
func (service *ProfileService) Help() error {
	return nil
}
