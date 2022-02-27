package bv5

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
	"github.com/rs/zerolog"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
	"strconv"
	"time"
)

func mapUserToJsonBody(user *db.User) *api.UserGet {
	return &api.UserGet{
		User: api.User{
			CommonAccess: user.CommonAccess,
			Email:        api.Email(user.Email),
			Name:         user.Name,
			SystemTags:   user.SystemTags,
			UserTags:     user.UserTags,
			Active:       user.IsActive,
		},
		Created: api.Created(user.CreatedAt.Format(time.RFC3339)),
		UserId:  strconv.FormatInt(user.UserID, 10),
		Uuid:    user.Uuid.String(),
		Version: user.Version,
	}
}

func createUserBySession(c *gin.Context, b *BardView5, logger zerolog.Logger, session *kratos.Session) (int64, error) {
	userUuid := uuid.MustParse(session.Identity.Id)

	traits, ok := session.Identity.Traits.(map[string]interface{})
	if !ok {
		logger.Error().Str("uuid", userUuid.String()).Msg("Failed to understand user session")
		return 0, ErrUnknownStatusRead(ObjUser, 0, true)
	}

	newUserId := b.generators.userNode.Generate().Int64()
	logger.Info().Str("uuid", userUuid.String()).Int64("id", newUserId).Msg("Creating user in BardView5")
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
		logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
		return 0, ErrUnknownStatusRead(ObjUser, newUserId, true)
	}
	if changedRows == 0 {
		logger.Err(err).Str("uuid", userUuid.String()).Int64("id", newUserId).Msg("No rows changed for user. Some conflict exists.")
		return 0, ErrNotFound(ObjUser, newUserId)
	}
	return newUserId, nil
}

func ensureKratosUserByUuid(b *BardView5Http, userUuid uuid.UUID) (*api.UserGet, error) {
	users, err := b.BardView5.Querier().UserFindByUuid(b.Context, userUuid)
	if err != nil {
		b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
		return nil, ErrUnknownStatusRead(ObjUser, 0, true)
	}
	if len(users) == 0 {
		session, err := b.BardView5.DepKratos.GetKratosSession(b)
		if err != nil {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to get user")
			return nil, ErrNotFound(ObjUser, 0)
		}

		newUserId, err := createUserBySession(b.Context, b.BardView5, b.Logger, session)

		//Prime user array for next fetch.
		users, err = b.BardView5.Querier().UserFindByUuid(b.Context, userUuid)
		if err != nil {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Int64("id", newUserId).Msg("Failed to get user")
			return nil, ErrFailedRead(ObjUser, newUserId, true)
		}
		if len(users) == 0 {
			return nil, ErrNotFound(ObjUser, newUserId)
		}
	}

	user := users[0]
	if err = userHasAccess(b, &user); err != nil {
		return mapUserToJsonBody(&user), err
	}

	return mapUserToJsonBody(&user), nil
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
		return createUserBySession(c, b, logger, session)
	} else {
		user := users[0]
		if !user.IsActive {
			return 0, sessionInactiveUser
		}
		return user.UserID, nil
	}
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

	session, err := b.DepKratos.GetKratosSessionM(c)
	if err != nil {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	userId, err := b.createOrGetUserByUuid(c, session)
	if err != nil || userId == 0 {
		c.Set(Session, MakeAnonymousSession())
		return
	}
	b.sessions.SetSessionCache(oryKratosSession, userId)
	c.Set(Session, MakeSession(userId))
}

func (b *BardView5) ApiRequireValidSession(c *gin.Context) {
	s := SessionCriteria(c)
	if s.Anonymous {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
