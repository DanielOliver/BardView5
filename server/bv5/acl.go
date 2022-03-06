package bv5

import (
	"context"
	"strconv"
)

const (
	ObjectUser = "user"

	SessionId = "session_id"
	Session   = "session"

	CommonAccessPrivate = "private"
	CommonAccessAnyUser = "anyuser"
	CommonAccessPublic  = "public"

	RoleActionOwner = "owner"

	RoleSubDnd5eSetting = "dnd5e_setting"
)

type SessionContext struct {
	AvailableFields map[string]string
	SessionId       int64
	Anonymous       bool
}

func SessionCriteria(context context.Context) *SessionContext {
	return context.Value(Session).(*SessionContext)
}

func MakeAnonymousSession() *SessionContext {
	return &SessionContext{
		map[string]string{
			SessionId: "",
		},
		0,
		true,
	}
}

func MakeSession(sessionId int64) *SessionContext {
	return &SessionContext{
		map[string]string{
			SessionId: strconv.FormatInt(sessionId, 10),
		},
		sessionId,
		sessionId == 0,
	}
}
