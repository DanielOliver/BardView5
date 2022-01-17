package bv5

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"server/api"
	"testing"
)

var (
	LocalConnectionString = "postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
)

func TestDnd5eWorldCreate(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	bv5, err := ConfigNewBardView5(&BardView5InitConfig{
		ConnectionString: LocalConnectionString,
		KratosBaseUrl:    "",
	})
	assert.Nil(t, err)
	newDnd5eWorldId, err := Dnd5eWorldCreate(&BardView5Http{
		BardView5: bv5,
		Logger:    zerolog.Nop(),
		Session: sessionContext{
			AvailableFields: map[string]string{},
			sessionId:       2,
			Anonymous:       false,
		},
		Context: ctx,
	}, &api.PostApiV1Dnd5eWorldsJSONBody{
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
