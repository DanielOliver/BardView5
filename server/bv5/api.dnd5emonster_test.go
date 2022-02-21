package bv5

import (
	"github.com/stretchr/testify/assert"
	"server/api"
	"strconv"
	"testing"
)

func TestDnd5eMonsterCreate(t *testing.T) {
	bv5Http := CreateBv5Test(6)
	CreateBv5TestSessionUser(bv5Http)

	newDnd5eSettingId, err := Dnd5eSettingCreate(bv5Http, &api.PostApiV1Dnd5eSettingsJSONBody{
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

	newDnd5eMonsterId, err := Dnd5eMonsterCreate(bv5Http, &api.PostApiV1Dnd5eMonstersJSONBody{
		Alignment:      nil,
		ArmorClass:     nil,
		Description:    nil,
		Dnd5eSettingId: strconv.FormatInt(newDnd5eSettingId, 10),
		Environments:   nil,
		HitPoints:      nil,
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
