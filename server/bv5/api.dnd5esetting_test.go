package bv5

import (
	"github.com/stretchr/testify/assert"
	"server/api"
	"testing"
)

func TestDnd5eSettingCreate(t *testing.T) {
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

	_, err = dnd5eSettingById(bv5Http, -100)
	assert.Error(t, err)
	assert.IsType(t, &CrudError{}, err)

	getThatSetting, err := dnd5eSettingById(bv5Http, newDnd5eSettingId)
	assert.NoError(t, err)
	assert.Equal(t, "Named", getThatSetting.Name)
	assert.Equal(t, "Describe", getThatSetting.Description)

	err = dnd5eSettingHasAccess(bv5Http, &getThatSetting)
	assert.NoError(t, err)
}
