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
	CommonAccessPublic  = "public"
)

type sessionContext struct {
	AvailableFields map[string]string
	sessionId       int64
	Anonymous       bool
}

func (s *sessionContext) SessionId() int64 {
	return s.sessionId
}

func SessionCriteria(context context.Context) *sessionContext {
	return context.Value(Session).(*sessionContext)
}

func MakeAnonymousSession() *sessionContext {
	return &sessionContext{
		map[string]string{
			SessionId: "",
		},
		0,
		true,
	}
}

func MakeSession(sessionId int64) *sessionContext {
	return &sessionContext{
		map[string]string{
			SessionId: strconv.FormatInt(sessionId, 10),
		},
		sessionId,
		sessionId == 0,
	}
}
