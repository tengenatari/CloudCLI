package repository

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	ctx        context.Context
	repository *YamlRepository
}

func (suite *RepositorySuite) SetupTest() {
	suite.ctx = context.Background()
	testDir := "./testConfigCli"

	err := os.Mkdir(testDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	suite.repository = NewYamlRepository(testDir)
}

func TestRunRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (suite *RepositorySuite) TearDownTest() {
	err := os.RemoveAll("./testConfigCli")
	if err != nil {
		return
	}

}
