package db

import (
	"server/bv5"
	"strconv"
)

func (u *User) GetAclMetadata() *bv5.AclObjectMetadata {
	return &bv5.AclObjectMetadata{
		map[string]string{
			"user_id": strconv.FormatInt(u.UserID, 10),
		},
		bv5.UserObject,
		u.UserID,
	}
}
