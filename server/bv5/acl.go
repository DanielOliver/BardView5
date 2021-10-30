package bv5

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
