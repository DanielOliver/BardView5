package bv5

import (
	"github.com/stretchr/testify/assert"
	"server/api"
	"testing"
)

func TestDnd5eMonsterCreate(t *testing.T) {
	bv5Http := CreateBv5Test(6)
	CreateBv5TestSessionUser(bv5Http)

	newDnd5eSettingId, err := dnd5eSettingCreate(bv5Http, &api.PostApiV1Dnd5eSettingsJSONBody{
		Active:       true,
		CommonAccess: CommonAccessPrivate,
		Description:  "Describe",
		Module:       nil,
		Name:         "Named",
		SystemTags:   []string{},
		UserTags:     []string{},
	})
	assert.NoError(t, err)
	assert.NotZero(t, newDnd5eSettingId)

	var milliChallengeRating int64 = 4000

	newDnd5eMonsterId, err := dnd5eMonsterCreate(bv5Http, &api.PostApiV1Dnd5eSettingsDnd5eSettingIdMonstersJSONBody{
		Alignment:    nil,
		ArmorClass:   nil,
		Description:  nil,
		Environments: nil,
		HitPoints:    nil,
		Languages: &[]string{
			"Goblin",
			"Common",
		},
		Legendary:            nil,
		MilliChallengeRating: &milliChallengeRating,
		MonsterType:          nil,
		Name:                 "Goblin",
		SizeCategory:         nil,
		Sources:              nil,
		Unique:               nil,
		UserTags:             []string{},
	}, newDnd5eSettingId)
	assert.NoError(t, err)
	assert.NotZero(t, newDnd5eMonsterId)
}
