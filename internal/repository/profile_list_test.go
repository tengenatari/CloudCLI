package repository

import (
	"CloudCLI/internal/models"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func (suite *RepositorySuite) TestProfileList() {

	profiles := make(map[string]*models.Profile)

	for x := 0; x < 10; x++ {
		user := fmt.Sprintf("user%d", x)
		name := fmt.Sprintf("name%d", x)
		project := fmt.Sprintf("project%d", x)

		profile := &models.Profile{User: user, Project: project}
		profiles[name] = profile

		err := suite.repository.ProfileCreate(user, name, project)

		assert.NoError(suite.T(), err)
	}

	actualProfiles, err := suite.repository.ProfileList()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), profiles, actualProfiles)

}
