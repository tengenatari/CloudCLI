package service

import "github.com/stretchr/testify/assert"

func (s *ServiceSuite) TestProfileDeleteCorrect() {

	params := make(map[string]string)
	params["name"] = "name1"

	s.repository.On("ProfileDelete", params["name"]).Return(nil).Once()

	err := s.service.ProfileDelete(params)

	assert.NoError(s.T(), err)
}

func (s *ServiceSuite) TestProfileDeleteFail() {
	params := make(map[string]string)
	params["name"] = "name%q1"

	err := s.service.ProfileDelete(params)

	assert.Error(s.T(), err)
}
