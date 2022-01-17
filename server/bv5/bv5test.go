package bv5

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http/httptest"
	"server/db"
)

var (
	LocalConnectionString = "postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
	DefaultUserId         = int64(5)
)

func CreateBv5Test(sessionId int64) *BardView5Http {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	bardview5, err := ConfigNewBardView5(&BardView5InitConfig{
		ConnectionString: LocalConnectionString,
		KratosBaseUrl:    "",
	})
	if err != nil {
		panic(err)
	}
	return &BardView5Http{
		BardView5: bardview5,
		Logger:    zerolog.Nop(),
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
