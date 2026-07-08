package service

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
