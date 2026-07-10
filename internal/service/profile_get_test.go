package service

import (
	"CloudCLI/internal/models"

	"github.com/stretchr/testify/assert"
)

func (s *ServiceSuite) TestProfileGetCorrect() {

	params := make(map[string]string)
	params["name"] = "name1"

	profile := models.Profile{User: "user1", Project: "project1"}
	s.repository.On("ProfileGet", params["name"]).Return(&profile, nil).Once()

	err := s.service.ProfileGet(params)

	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestProfileGetFail() {
	params := make(map[string]string)
	params["name"] = "name%q1"

	err := s.service.ProfileGet(params)

	assert.Error(s.T(), err)
}
