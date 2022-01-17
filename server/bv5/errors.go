package bv5

import (
	"fmt"
	"net/http"
)

type SessionError struct {
	msg string
}

func (s *SessionError) Error() string {
	return s.msg
}

type CrudError struct {
	msg        string
	Object     string
	Id         int64
	IsInternal bool
}

func (s *CrudError) Error() string {
	return s.msg
}

type ContextWriter interface {
	WriteToContext(b *BardView5Http)
}

func (s *CrudError) WriteToContext(b *BardView5Http) {
	b.Logger.Err(s).Int64("id", s.Id).Str("obj", s.Object).Msg("CrudError")
	status := http.StatusBadRequest
	if s.IsInternal {
		status = http.StatusInternalServerError
	}
	b.Context.AbortWithStatusJSON(status, s.msg)
}

func WriteErrorToContext(b *BardView5Http, err error) {
	switch t := err.(type) {
	case ContextWriter:
		t.WriteToContext(b)
		return
	default:
		b.Logger.Err(err).Msg("")
		b.Context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func ErrFailedCreate(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Failed to create: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
	}
}
func ErrUnknownStatusCreate(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Creation left in unknown state: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
	}
}

func ErrFailedRead(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Failed to read: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
	}
}

var (
	session401 = &SessionError{
		msg: "401",
	}
	sessionInactiveUser = &SessionError{
		msg: "User not active",
	}
	ObjDnd5eWorld           = "dnd5eworld"
	ObjDnd5eWorldAssignment = "dnd5eworldassignment"
	ObjUser                 = "user"
)
