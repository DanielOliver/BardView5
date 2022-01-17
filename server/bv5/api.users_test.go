package bv5

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"server/api"
	"testing"
)

func TestUserCreateWithoutAuthorization(t *testing.T) {
	bv5Http := CreateBv5Test(6)
	testKratosImpl := &TestKratosImpl{EnsureUuid: uuid.New()}
	bv5Http.BardView5.DepKratos = testKratosImpl
	_, err := ensureKratosUserByUuid(bv5Http, testKratosImpl.EnsureUuid)
	assert.Error(t, err)
	//assert.NotNil(t, userGet)
	assert.IsType(t, &CrudError{}, err)
}

func TestUserCreate(t *testing.T) {
	bv5Http := CreateBv5Test(6)
	testKratosImpl := &TestKratosImpl{EnsureUuid: uuid.New()}
	bv5Http.BardView5.DepKratos = testKratosImpl
	_, err := ensureKratosUserByUuid(bv5Http, testKratosImpl.EnsureUuid)
	assert.Error(t, err)
	assert.IsType(t, &CrudError{}, err)

	users, err := bv5Http.BardView5.Querier().UserFindByUuid(bv5Http.Context, testKratosImpl.EnsureUuid)
	user := users[0]
	assert.NoError(t, err)
	bv5Http.Session.SessionId = user.UserID

	userGet, err := ensureKratosUserByUuid(bv5Http, testKratosImpl.EnsureUuid)
	assert.NoError(t, err)
	assert.NotNil(t, userGet)

	ensureUuid := testKratosImpl.EnsureUuid.String()
	assert.Equal(t, ensureUuid, userGet.Uuid)
	assert.Equal(t, api.Email(ensureUuid+"@test.com"), userGet.Email)
}
