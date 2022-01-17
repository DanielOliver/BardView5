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
	assert.Nil(t, err)
	assert.NotZero(t, newDnd5eWorldId)
}
