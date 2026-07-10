package repository

import "github.com/stretchr/testify/assert"

func (suite *RepositorySuite) TestProfileCreateCorrect() {
	name := "name"
	project := "project"
	user := "user"

	err := suite.repository.ProfileCreate(user, name, project)

	assert.NoError(suite.T(), err)
}

func (suite *RepositorySuite) TestDoubleProfileCreateError() {
	name := "name2"
	project := "project2"
	user := "user2"

	name2 := "name2"
	project2 := "project2"
	user2 := "user2"

	err := suite.repository.ProfileCreate(user, name, project)
	err2 := suite.repository.ProfileCreate(user2, name2, project2)

	assert.NoError(suite.T(), err)
	assert.Error(suite.T(), err2)
}
