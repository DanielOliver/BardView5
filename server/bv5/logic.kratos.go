package bv5

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	kratos "github.com/ory/kratos-client-go"
	"net/http"
	"server/api"
	"server/bardlog"
	"server/db"
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
		UserId:  user.UserID,
		Uuid:    user.Uuid.String(),
		Version: user.Version,
	}
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

		traits, ok := session.Identity.Traits.(map[string]interface{})
		if !ok {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
			return nil, ErrUnknownStatusRead(ObjUser, 0, true)
		}

		newUserId := b.BardView5.generators.userNode.Generate().Int64()
		newUserUuid := uuid.MustParse(session.Identity.Id)
		b.Logger.Info().Str("uuid", userUuid.String()).Int64("id", newUserId).Msg("Creating user in BardView5")
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
			b.Logger.Err(err).Str("uuid", userUuid.String()).Msg("Failed to understand user")
			return nil, ErrUnknownStatusRead(ObjUser, newUserId, true)
		}
		if changedRows == 0 {
			b.Logger.Err(err).Str("uuid", userUuid.String()).Int64("id", newUserId).Msg("No rows changed for user. Some conflict exists.")
			return nil, ErrNotFound(ObjUser, newUserId)
		}

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
	if err = UserHasAccess(b, &user); err != nil {
		return nil, err
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
