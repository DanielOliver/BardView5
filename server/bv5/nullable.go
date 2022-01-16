package bv5

import (
	"database/sql"
)

func MaybeInt64(value *int64) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}
	return sql.NullInt64{
		Int64: *value,
		Valid: false,
	}
}

func MaybeString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: *value,
		Valid:  false,
	}
}
