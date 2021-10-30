package acl

import (
	"context"
	"encoding/json"
	"server/bv5"
	"server/db"
)

type sessionCriteria struct {
	AvailableFields map[string]string
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
	return &sessionCriteria{
		map[string]string{
			bv5.SessionId: context.Value(bv5.SessionId).(string),
		},
	}
}

func (s *sessionCriteria) Evaluate(object *bv5.AclObjectMetadata, acl []db.GetAclBySubjectRow, session *sessionCriteria) *aclEvaluation {
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
