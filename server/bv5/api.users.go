package bv5

import (
	"database/sql"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
	"strconv"
	"time"
)

func (b *BardView5) PostUsersCreate(c *gin.Context) {
	session := SessionCriteria(c)
	logger := bardlog.GetLogger(c)

	var body api.PostApiV1UsersJSONBody
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
			IsActive:     body.Active,
		})
		if err != nil {
			logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Error updating user")
			c.AbortWithStatusJSON(http.StatusBadRequest, "Error updating")
			return
		}
		if len(updatedUserRows) == 0 {
			logger.Warn().Str(bardlog.KeySubjectType, ObjectUser).Msg("Version out of date")
			c.AbortWithStatus(http.StatusNotModified)
			return
		}
		updatedUser := updatedUserRows[0]
		c.Header("ETag", strconv.FormatInt(updatedUser.Version, 10))
		c.Header("Location", fmt.Sprintf("/v1/users%d/", updatedUser.UserID))
		c.JSON(http.StatusOK, api.UserPostOk{
			UserId:  updatedUser.UserID,
			Version: updatedUser.Version,
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
				session.SessionId,
				true,
			},
			IsActive: body.Active,
		})
		if err != nil {
			logger.Err(err).Msg("Failed to create new user")
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to create new user")
			return
		}
		if changedRows == 0 {
			c.JSON(http.StatusBadRequest, "Failed to create new user")
			return
		}
		c.Header("ETag", "0")
		c.Header("Location", fmt.Sprintf("/v1/users/%d/", newUserId))
		c.JSON(http.StatusCreated, api.UserPostOk{
			UserId:  newUserId,
			Version: 0,
		})
	}
}

func GetUserThatIsMe(b *BardView5Http) {
	session, err := b.BardView5.DepKratos.GetKratosSession(b)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	userUuid := uuid.MustParse(session.Identity.Id)
	userGet, err := ensureKratosUserByUuid(b, userUuid)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", strconv.FormatInt(userGet.Version, 10))
	b.Context.Header("Location", fmt.Sprintf("/v1/users/%d/", userGet.UserId))
	b.Context.JSON(http.StatusOK, userGet)
}

type GetUserByIdParams struct {
	UserID int64 `uri:"userId" binding:"required"`
}

type GetUserByUuidParams struct {
	UserID string `uri:"userId" binding:"required,uuid"`
}

func GetUsersById(b *BardView5Http) {
	var params GetUserByIdParams
	if err := b.Context.ShouldBindUri(&params); err != nil {

		var uuidParams GetUserByUuidParams
		if err := b.Context.ShouldBindUri(&uuidParams); err != nil {
			b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userGet, err := ensureKratosUserByUuid(b, uuid.MustParse(uuidParams.UserID))
		if err != nil {
			WriteErrorToContext(b, err)
			return
		}
		b.Context.Header("ETag", strconv.FormatInt(userGet.Version, 10))
		b.Context.Header("Location", fmt.Sprintf("/v1/users/%d/", userGet.UserId))
		b.Context.JSON(http.StatusOK, userGet)
		return
	}
	userGet, err := getUserById(b, params.UserID)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}
	b.Context.Header("ETag", strconv.FormatInt(userGet.Version, 10))
	b.Context.Header("Location", fmt.Sprintf("/v1/users/%d/", userGet.UserId))
	b.Context.JSON(http.StatusOK, userGet)
}

func getUserById(b *BardView5Http, userId int64) (*api.UserGet, error) {
	users, err := b.BardView5.Querier().UserFindById(b.Context, userId)
	if err != nil {
		b.Logger.Err(err).Int64("id", userId).Msg("Failed to get user")
		return nil, ErrUnknownStatusRead(ObjUser, userId, true)
	}
	if len(users) == 0 {
		return nil, ErrNotFound(ObjUser, userId)
	}
	user := users[0]
	if err = UserHasAccess(b, &user); err != nil {
		return nil, err
	}
	return mapUserToJsonBody(&user), nil
}

func (b *BardView5) PatchUserById(c *gin.Context) {
	var params GetUserByIdParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger := bardlog.GetLogger(c)
	users, err := b.Querier().UserFindById(c, params.UserID)
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to find")
		c.AbortWithStatusJSON(http.StatusNotFound, "Failed to find user")
		return
	}
	if len(users) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, "Failed to find user")
		return
	}
	user := users[0]
	originalUserJson, err := json.Marshal(api.UserGet{
		User: api.User{
			CommonAccess: user.CommonAccess,
			Email:        api.Email(user.Email),
			Name:         user.Name,
			SystemTags:   user.SystemTags,
			UserTags:     user.UserTags,
			Active:       user.IsActive,
		},
		Created: api.Created(user.CreatedAt.Format(time.RFC3339)),
		UserId:  user.UserID,
		Version: user.Version,
	})
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to patch")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to patch user")
		return
	}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to patch")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to patch user")
		return
	}
	patch, err := jsonpatch.DecodePatch(jsonData)
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to patch")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to patch user")
		return
	}
	modifiedUserJson, err := patch.Apply(originalUserJson)
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to patch")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to patch user")
		return
	}
	var modifiedUser api.UserGet
	if err := json.Unmarshal(modifiedUserJson, &modifiedUser); err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Failed to patch")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to patch user")
		return
	}

	updatedUserRows, err := b.Querier().UserUpdate(c, db.UserUpdateParams{
		Name:         modifiedUser.Name,
		UserTags:     modifiedUser.UserTags,
		SystemTags:   modifiedUser.SystemTags,
		CommonAccess: modifiedUser.CommonAccess,
		IsActive:     modifiedUser.Active,
		UserID:       user.UserID,
		Version:      user.Version,
	})
	if err != nil {
		logger.Err(err).Str(bardlog.KeySubjectType, ObjectUser).Msg("Error updating user")
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error updating")
		return
	}
	if len(updatedUserRows) == 0 {
		logger.Warn().Str(bardlog.KeySubjectType, ObjectUser).Msg("Version out of date")
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	updatedUser := updatedUserRows[0]
	c.Header("ETag", strconv.FormatInt(updatedUser.Version, 10))
	c.Header("Location", fmt.Sprintf("/v1/users%d/", updatedUser.UserID))
	c.JSON(http.StatusOK, api.UserGet{
		User: api.User{
			CommonAccess: updatedUser.CommonAccess,
			Email:        api.Email(updatedUser.Email),
			Name:         updatedUser.Name,
			SystemTags:   updatedUser.SystemTags,
			UserTags:     updatedUser.UserTags,
			Active:       updatedUser.IsActive,
		},
		Created: api.Created(updatedUser.CreatedAt.Format(time.RFC3339)),
		UserId:  updatedUser.UserID,
		Version: updatedUser.Version,
	})
}

func UserHasAccess(b *BardView5Http, user *db.User) error {
	switch user.CommonAccess {
	case CommonAccessPublic:
		return nil
	case CommonAccessAnyUser:
		if b.Session.Anonymous {
			return ErrNotAuthorized(ObjUser, user.UserID)
		}
		return nil
	case CommonAccessPrivate:
		if b.Session.SessionId != user.UserID {
			return ErrNotAuthorized(ObjUser, user.UserID)
		}
		return nil
	default:
		return ErrNotAuthorized(ObjUser, user.UserID)
	}
}
