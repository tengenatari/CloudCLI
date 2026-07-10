package service

import (
	"github.com/stretchr/testify/assert"
)

func (s *ServiceSuite) TestProfileCreateCorrect() {

	params := make(map[string]string)
	params["name"] = "name1"
	params["user"] = "user1"
	params["project"] = "project1"

	s.repository.On("ProfileCreate", params["user"], params["name"], params["project"]).Return(nil).Once()

	err := s.service.ProfileCreate(params)

	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestProfileCreateFail() {
	params := make(map[string]string)
	params["name"] = "name%q1"
	params["user"] = "user1"
	params["project"] = "project1"

	err := s.service.ProfileCreate(params)

	assert.Error(s.T(), err)
}
