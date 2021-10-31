package db

func (u *User) UserId_() int64 {
	return u.UserID
}

func (u *User) IsActive_() bool {
	return u.IsActive
}

func (u *User) CommonAccess_() string {
	return u.CommonAccess
}
