package bv5

import (
	"fmt"
	"net/http"
)

type ErrorType int
type DepHttpType string

const (
	ErrTUnknown ErrorType = iota
	ErrTNotFound
	ErrTNotAuthorized
	ErrTFailedToCreate
	ErrTFailedToRead
	ErrTInit

	ObjDnd5eSetting           = "dnd5esetting"
	ObjDnd5eSettingAssignment = "dnd5esettingassignment"
	ObjDnd5eMonster           = "dnd5emonster"
	ObjUser                   = "user"

	DepHttpKratos DepHttpType = "httpkratos"
)

var (
	session401 = &SessionError{
		msg: "401",
	}
	sessionInactiveUser = &SessionError{
		msg: "User not active",
	}
)

type SessionError struct {
	msg string
}

func (s *SessionError) Error() string {
	return s.msg
}

func (s *SessionError) WriteToContext(b *BardView5Http) {
	b.Logger.Err(s).Msg("SessionError")
	status := http.StatusUnauthorized
	b.Context.AbortWithStatus(status)
}

type CrudError struct {
	msg        string
	Object     string
	Id         int64
	IsInternal bool
	ErrorType  ErrorType
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
	switch s.ErrorType {
	case ErrTNotFound:
		status = http.StatusNotFound
		break
	case ErrTNotAuthorized:
		status = http.StatusUnauthorized
		b.Context.AbortWithStatus(status)
		return
	}
	b.Context.AbortWithStatusJSON(status, s.msg)
}

type HttpError struct {
	msg        string
	Dependency DepHttpType
	Relative   string
	ErrorType  ErrorType
}

func (s *HttpError) Error() string {
	return s.msg
}

func (s *HttpError) WriteToContext(b *BardView5Http) {
	b.Logger.Err(s).Msg("HttpError")
	status := http.StatusInternalServerError
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
		ErrorType:  ErrTFailedToCreate,
	}
}
func ErrUnknownStatusCreate(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Creation left in unknown state: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
		ErrorType:  ErrTFailedToCreate,
	}
}

func ErrUnknownStatusRead(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Read failed by unknown state: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
		ErrorType:  ErrTFailedToRead,
	}
}

func ErrFailedRead(obj string, id int64, internal bool) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Failed to read: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: internal,
		ErrorType:  ErrTFailedToRead,
	}
}
func ErrNotFound(obj string, id int64) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Not Found: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: false,
		ErrorType:  ErrTNotFound,
	}
}
func ErrNotAuthorized(obj string, id int64) *CrudError {
	return &CrudError{
		msg:        fmt.Sprintf("Not Authorized: %s, %d", obj, id),
		Object:     obj,
		Id:         id,
		IsInternal: false,
		ErrorType:  ErrTNotAuthorized,
	}
}

func ErrHttpDepUnknown(relative string, httpType DepHttpType) *HttpError {
	return &HttpError{
		msg:        fmt.Sprintf("Failed http call: %s, %s", httpType, relative),
		Dependency: httpType,
		Relative:   relative,
		ErrorType:  ErrTUnknown,
	}
}

func ErrHttpDepInit(relative string, httpType DepHttpType) *HttpError {
	return &HttpError{
		msg:        fmt.Sprintf("Failed to init http call: %s, %s", httpType, relative),
		Dependency: httpType,
		Relative:   relative,
		ErrorType:  ErrTInit,
	}
}
