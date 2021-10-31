package bv5

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
)

func (b *BardView5) PostUsersCreate(c *gin.Context) {
	session := NewSessionCriteria(c)
	logger := bardlog.GetLogger(c)

	var body api.PostUsersJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingCheckParams := db.UserFindByIdOrEmailOrUuidParams{
		UserID: 0,
		Email:  string(body.Email),
		Uuid:   uuid.UUID{},
	}
	if body.UserId != nil {
		existingCheckParams.UserID = *body.UserId
	}

	usersFound, err := b.Querier().UserFindByIdOrEmailOrUuid(c, existingCheckParams)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(usersFound) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Doesn't yet support updating users")
		return
	}

	newUserId := b.generators.userNode.Generate().Int64()
	newUserUuid := uuid.New()
	changedRows, err := b.Querier().UserInsert(c, db.UserInsertParams{
		UserID:       newUserId,
		Uuid:         newUserUuid,
		Name:         body.Name,
		Email:        string(body.Email),
		UserTags:     body.UserTags,
		SystemTags:   body.SystemTags,
		CommonAccess: body.CommonAccess,
		CreatedBy: sql.NullInt64{
			session.SessionId(),
			false,
		},
	})
	if err != nil {
		logger.Err(err).Msg("Failed to create new user")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if changedRows == 0 {
		c.JSON(http.StatusBadRequest, "Failed to create new user")
	}
	c.JSON(http.StatusOK, newUserId)
}
