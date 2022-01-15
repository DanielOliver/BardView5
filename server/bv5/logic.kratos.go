package bv5

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
	"io"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
	"time"
)

type SessionError struct {
	msg string
}

func (s *SessionError) Error() string {
	return s.msg
}

var (
	session401 = &SessionError{
		msg: "401",
	}
)

func (b *BardView5) getUserByUuid(c *gin.Context, userUuid uuid.UUID) {
	logger := bardlog.GetLogger(c)
	users, err := b.Querier().UserFindByUuid(c, userUuid)
	if err != nil {
		logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to return user")
		return
	}
	if len(users) == 0 {
		session, err := b.getKratosSession(c)
		if err != nil {
			logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
			c.AbortWithStatusJSON(http.StatusNotFound, "Failed to return user")
			return
		}

		traits, ok := session.Identity.Traits.(map[string]interface{})
		if !ok {
			logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to understand user")
			return
		}

		newUserId := b.generators.userNode.Generate().Int64()
		newUserUuid := uuid.MustParse(session.Identity.Id)
		changedRows, err := b.Querier().UserInsert(c, db.UserInsertParams{
			UserID:       newUserId,
			Uuid:         newUserUuid,
			Name:         traits["username"].(string),
			Email:        traits["email"].(string),
			UserTags:     []string{},
			SystemTags:   []string{"implicit_registration"},
			CommonAccess: CommonAccessPrivate,
			CreatedBy:    sql.NullInt64{},
			IsActive:     true,
		})
		if err != nil {
			logger.Err(err).Msg("Failed to confirm new user")
			c.AbortWithStatusJSON(http.StatusBadRequest, "Failed to confirm new user")
			return
		}
		if changedRows == 0 {
			c.JSON(http.StatusBadRequest, "Failed to confirm new user")
			return
		}
		c.Header("ETag", "0")
		c.Header("Location", fmt.Sprintf("/v1/users%d/", newUserId))
		c.JSON(http.StatusCreated, api.UserGetOk{
			UserId:  newUserId,
			Version: 0,
		})

	} else {
		user := users[0]
		c.JSON(http.StatusOK, api.UserGet{
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
	}
}

func (b *BardView5) createOrGetUserByUuid(c *gin.Context, session *kratos.Session) (int64, error) {
	logger := bardlog.GetLogger(c)
	userUuid := uuid.MustParse(session.Identity.Id)
	users, err := b.Querier().UserFindByUuid(c, userUuid)
	if err != nil {
		logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0, err
	}
	if len(users) == 0 {
		traits, ok := session.Identity.Traits.(map[string]interface{})
		if !ok {
			logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
			c.AbortWithStatus(http.StatusUnauthorized)
			return 0, err
		}

		newUserId := b.generators.userNode.Generate().Int64()
		changedRows, err := b.Querier().UserInsert(c, db.UserInsertParams{
			UserID:       newUserId,
			Uuid:         userUuid,
			Name:         traits["username"].(string),
			Email:        traits["email"].(string),
			UserTags:     []string{},
			SystemTags:   []string{"implicit_registration"},
			CommonAccess: CommonAccessPrivate,
			CreatedBy:    sql.NullInt64{},
			IsActive:     true,
		})
		if err != nil {
			logger.Err(err).Msg("Failed to confirm new user")
			c.AbortWithStatus(http.StatusUnauthorized)
			return 0, err
		}
		if changedRows == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 0, err
		}
		return newUserId, nil
	} else {
		user := users[0]
		return user.UserID, nil
	}
}

func (b *BardView5) getKratosSession(c *gin.Context) (*kratos.Session, error) {
	logger := bardlog.GetLogger(c)
	req, err := http.NewRequest("GET", b.conf.kratosBaseUrl+"/sessions/whoami", nil)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	req.Header.Add("Cookie", c.Request.Header.Get("Cookie"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, session401
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Err(err)
		return nil, err
	}

	var result kratos.Session
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err)
		return nil, err
	}
	return &result, nil
}

func (b *BardView5) AddSessionToContext(c *gin.Context) {
	oryKratosSession, err := c.Cookie("ory_kratos_session")
	if err == http.ErrNoCookie {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	cachedUserId, foundUserId := b.sessions.GetSessionCache(oryKratosSession)
	if foundUserId {
		c.Set(Session, MakeSession(cachedUserId))
		return
	}

	session, err := b.getKratosSession(c)
	if err != nil {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	userId, err := b.createOrGetUserByUuid(c, session)
	if err != nil {
		return
	}
	b.sessions.SetSessionCache(oryKratosSession, userId)
	c.Set(Session, MakeSession(userId))
}

func (b *BardView5) RequireValidSession(c *gin.Context) {
	s := SessionCriteria(c)
	if s.Anonymous {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
