package service

import (
	"CloudCLI/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctx        context.Context
	repository *mocks.RepositoryInterface
	service    *ProfileService
}

func (s *ServiceSuite) SetupTest() {
	s.repository = mocks.NewRepositoryInterface(s.T())
	s.service = NewProfileService(s.repository)
	s.ctx = context.Background()
}

func TestRunServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
