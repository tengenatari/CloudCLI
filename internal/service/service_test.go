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

func (suite *ServiceSuite) SetupTest() {
	suite.repository = mocks.NewRepositoryInterface(suite.T())
	suite.service = NewProfileService(suite.repository)
	suite.ctx = context.Background()
}

func TestRunServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
