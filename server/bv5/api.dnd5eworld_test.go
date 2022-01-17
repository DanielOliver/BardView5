package bv5

import (
	"github.com/stretchr/testify/assert"
	"server/api"
	"testing"
)

func TestDnd5eWorldCreate(t *testing.T) {
	bv5Http := CreateBv5Test(6)
	CreateBv5TestSessionUser(bv5Http)

	newDnd5eWorldId, err := Dnd5eWorldCreate(bv5Http, &api.PostApiV1Dnd5eWorldsJSONBody{
		Active:           true,
		CommonAccess:     CommonAccessPrivate,
		DerivedFromWorld: nil,
		Description:      "Describe",
		Module:           nil,
		Name:             "Named",
		SystemTags:       []string{},
		UserTags:         []string{},
	})
	assert.NoError(t, err)
	assert.NotZero(t, newDnd5eWorldId)

	_, err = Dnd5eWorldById(bv5Http, -100)
	assert.Error(t, err)
	assert.IsType(t, &CrudError{}, err)

	getThatWorld, err := Dnd5eWorldById(bv5Http, newDnd5eWorldId)
	assert.NoError(t, err)
	assert.Equal(t, "Named", getThatWorld.Name)
	assert.Equal(t, "Describe", getThatWorld.Description)

	err = Dnd5eWorldHasAccess(bv5Http, &getThatWorld)
	assert.NoError(t, err)
}
