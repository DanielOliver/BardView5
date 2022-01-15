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
	"strconv"
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
	sessionInactiveUser = &SessionError{
		msg: "User not active",
	}
)

func getUserByUuid(b *BardView5Http, userUuid uuid.UUID) {
	users, err := b.BardView5.Querier().UserFindByUuid(b.Context, userUuid)
	if err != nil {
		b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
		b.Context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		session, err := b.BardView5.getKratosSession(b.Context)
		if err != nil {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
			b.Context.AbortWithStatus(http.StatusNotFound)
			return
		}

		traits, ok := session.Identity.Traits.(map[string]interface{})
		if !ok {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
			b.Context.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		newUserId := b.BardView5.generators.userNode.Generate().Int64()
		newUserUuid := uuid.MustParse(session.Identity.Id)
		b.Logger.Info().Str("uuid", userUuid.String()).Int64("userid", newUserId).Msg("Creating user in BardView5")
		changedRows, err := b.BardView5.Querier().UserInsert(b.Context, db.UserInsertParams{
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
			b.Logger.Err(err).Msg("Failed to confirm new user")
			b.Context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if changedRows == 0 {
			b.Logger.Err(err).Msg("No rows changed for user")
			b.Context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		b.Context.Header("ETag", "0")
		b.Context.Header("Location", fmt.Sprintf("/v1/users/%d/", newUserId))
		b.Context.JSON(http.StatusCreated, api.UserGetOk{
			UserId:  newUserId,
			Version: 0,
		})

		//Prime user array for next fetch.
		users, err = b.BardView5.Querier().UserFindByUuid(b.Context, userUuid)
	}

	if len(users) == 0 {
		b.Context.AbortWithStatus(http.StatusNotFound)
		return
	}

	user := users[0]
	if user.CommonAccess == CommonAccessPublic || b.Session.sessionId == user.UserID {
		b.Context.Header("ETag", strconv.FormatInt(user.Version, 10))
		b.Context.Header("Location", fmt.Sprintf("/v1/users/%d/", user.UserID))
		b.Context.JSON(http.StatusOK, api.UserGet{
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
		return
	}

	b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("NotAuthorized")
	b.Context.AbortWithStatus(http.StatusNotFound)
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
		if !user.IsActive {
			return 0, sessionInactiveUser
		}

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
		c.Set(Session, MakeAnonymousSession())
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
