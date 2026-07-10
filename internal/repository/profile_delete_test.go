package repository

import (
	"github.com/stretchr/testify/assert"
)

func (suite *RepositorySuite) TestProfileDeleteCorrect() {
	name := "nameToDelete"
	project := "projectToDelete"
	user := "userToDelete"

	errCreate := suite.repository.ProfileCreate(user, name, project)
	errDelete := suite.repository.ProfileDelete(name)
	assert.NoError(suite.T(), errCreate)
	assert.Error(suite.T(), errDelete)

}

func (suite *RepositorySuite) TestProfileDeleteFail() {
	name := "nameToDelete"

	errDelete := suite.repository.ProfileDelete(name)
	assert.Error(suite.T(), errDelete)
}
