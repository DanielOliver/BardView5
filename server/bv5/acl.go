package bv5

const (
	UserObject = "user"

	SessionId = "session_id"
)


type AclObjectMetadata struct {
	AvailableFields map[string]string
	ObjectName      string
	Id              int64
}

type AclImplementer interface {
	GetAclMetadata() *AclObjectMetadata
}
