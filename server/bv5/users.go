package bv5

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
	"strconv"
)

type userPostResponse struct {
	UserId  int64 `json:"user_id"`
	Version int64 `json:"version"`
}

func (b *BardView5) PostUsersCreate(c *gin.Context) {
	session := NewSessionCriteria(c)
	logger := bardlog.GetLogger(c)

	var body api.PostUsersJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersFound, err := b.Querier().UserFindByEmail(c, string(body.Email))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(usersFound) > 0 {
		userToUpdate := usersFound[0]
		updatedUserRows, err := b.Querier().UserUpdate(c, db.UserUpdateParams{
			Name:         body.Name,
			UserTags:     body.UserTags,
			SystemTags:   body.SystemTags,
			CommonAccess: body.CommonAccess,
			UserID:       userToUpdate.UserID,
			Version:      userToUpdate.Version,
		})
		if err != nil {
			logger.Err(err).Msg("Error updating user")
			c.AbortWithStatusJSON(http.StatusBadRequest, "Error updating user")
			return
		}
		if len(updatedUserRows) == 0 {
			logger.Error().Msg("Error updating user")
			c.AbortWithStatusJSON(http.StatusBadRequest, "Error updating user")
			return
		}
		updatedUser := updatedUserRows[0]
		c.Header("ETag", strconv.FormatInt(updatedUser.Version, 10))
		c.Header("Location", fmt.Sprintf("/users%d/", updatedUser.UserID))
		c.JSON(http.StatusOK, userPostResponse{
			UserId:  updatedUser.UserID,
			Version: userToUpdate.Version,
		})
	} else {
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
			c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to create new user")
			return
		}
		if changedRows == 0 {
			c.JSON(http.StatusBadRequest, "Failed to create new user")
			return
		}
		c.Header("ETag", "0")
		c.Header("Location", fmt.Sprintf("/users%d/", newUserId))
		c.JSON(http.StatusCreated, userPostResponse{
			UserId:  newUserId,
			Version: 0,
		})
	}
}
