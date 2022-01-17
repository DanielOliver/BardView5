package bv5

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
	"github.com/rs/zerolog"
	"net/http/httptest"
	"os"
	"server/db"
)

var (
	LocalConnectionString = "postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
	DefaultUserId         = int64(5)
)

type TestKratosImpl struct {
	EnsureUuid uuid.UUID
}

func (t *TestKratosImpl) GetKratosSessionM(c *gin.Context) (*kratos.Session, error) {
	everything := t.EnsureUuid.String()

	return kratos.NewSession(uuid.NewString(), kratos.Identity{
		CreatedAt:         nil,
		Credentials:       nil,
		Id:                everything,
		RecoveryAddresses: nil,
		SchemaId:          "",
		SchemaUrl:         "",
		State:             nil,
		StateChangedAt:    nil,
		Traits: map[string]interface{}{
			"username": everything,
			"email":    fmt.Sprintf("%s@test.com", everything),
		},
	}), nil
}

func (t *TestKratosImpl) GetKratosSession(b *BardView5Http) (*kratos.Session, error) {
	return t.GetKratosSessionM(b.Context)
}

func CreateBv5Test(sessionId int64) *BardView5Http {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	bardview5, err := ConfigNewBardView5(&BardView5InitConfig{
		ConnectionString: LocalConnectionString,
		KratosBaseUrl:    "",
	})
	bardview5.DepKratos = &TestKratosImpl{
		EnsureUuid: uuid.New(),
	}
	if err != nil {
		panic(err)
	}
	return &BardView5Http{
		BardView5: bardview5,
		Logger:    zerolog.New(os.Stdout),
		Session:   *MakeSession(sessionId),
		Context:   ctx,
	}
}

func CreateBv5TestSessionUser(b *BardView5Http) {
	_, err := b.Querier().UserInsert(b.Context, db.UserInsertParams{
		UserID:       b.Session.SessionId,
		Uuid:         uuid.New(),
		Name:         "Daniel",
		Email:        fmt.Sprintf("test.%d@test.com", b.Session.SessionId),
		UserTags:     []string{},
		SystemTags:   []string{},
		CreatedBy:    sql.NullInt64{Valid: false},
		CommonAccess: "public",
		IsActive:     true,
	})
	if err != nil {
		panic(err)
	}
}
