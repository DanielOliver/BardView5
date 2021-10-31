package bv5

import (
	"context"
	"encoding/json"
	"server/db"
	"strconv"
)

const (
	ObjectUser = "user"

	SessionId = "session_id"

	ActionManage = "manage"
	ActionOwner = "owner"
	ActionView = "view"
	ActionPublicView = "public_view"

	CommonAccessPrivate = "private"
	CommonAccessPublic = "public"
)

type AclObjectMetadata struct {
	AvailableFields map[string]string
	ObjectName      string
	Id              int64
}

type AclImplementer interface {
	GetAclMetadata() *AclObjectMetadata
}

type SessionContext interface {
	SessionId() int64
}

type AclEvaluator interface {
	SystemEvaluate(session SessionContext) []string
}

type ModelCommons interface {
	UserId_() int64
	IsActive_() bool
	CommonAccess_() string
}

type sessionCriteria struct {
	AvailableFields map[string]string
	sessionId       int64
}

func (s *sessionCriteria) SessionId() int64 {
	return s.sessionId
}

type aclEvaluation struct {
	Object  string
	Id      int64
	Actions []string
}

type aclCondition struct {
	Op    string  `json:"op"`
	Field *string `json:"field"`
}

func NewSessionCriteria(context context.Context) *sessionCriteria {
	sessionIdStr := context.Value(SessionId).(string)
	sessionId, _ := strconv.ParseInt(sessionIdStr, 10, 64)
	return &sessionCriteria{
		map[string]string{
			SessionId: sessionIdStr,
		},
		sessionId,
	}
}

func (s *sessionCriteria) Evaluate(object *AclObjectMetadata, acl []db.GetAclBySubjectRow, session *sessionCriteria) *aclEvaluation {
	result := &aclEvaluation{
		Object:  object.ObjectName,
		Id:      object.Id,
		Actions: nil,
	}
	actions := make(map[string]bool)
	for _, row := range acl {

		if row.Subject != object.ObjectName {
			continue
		}
		if row.SubjectID.Valid {
			if row.SubjectID.Int64 == object.Id {
				actions[row.Action] = true
			} else {
				continue
			}
		}
		if row.Conditions != nil {
			var conditions map[string]aclCondition
			if err := json.Unmarshal(row.Conditions, &conditions); err == nil {
				for field, condition := range conditions {
					if condition.Op == "eq" {
						if condition.Field != nil && object.AvailableFields[field] == session.AvailableFields[*condition.Field] {
							actions[row.Action] = true
						}
					}
				}
			}
		}
	}

	for action, _ := range actions {
		result.Actions = append(result.Actions, action)
	}
	return result
}

