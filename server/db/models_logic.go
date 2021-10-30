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
		bv5.ObjectUser,
		u.UserID,
	}
}

func (u *User) SystemEvaluate(session bv5.SessionContext) []string {
	if u.UserID == session.SessionId() {
		if u.IsActive {
			return []string{
				bv5.ActionOwner,
				bv5.ActionManage,
				bv5.ActionView,
			}
		} else {
			return []string{
				bv5.ActionView,
			}
		}
	}
	if u.CommonAccess == bv5.CommonAccessPublic {
		return []string{
			bv5.ActionPublicView,
		}
	}
	return []string{}
}
