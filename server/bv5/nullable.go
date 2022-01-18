package bv5

import (
	"database/sql"
	"strconv"
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
		Valid: true,
	}
}

func MaybeInt64S(value *string) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}
	val, err := strconv.ParseInt(*value, 10, 64)
	if err != nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}
	return sql.NullInt64{
		Int64: val,
		Valid: true,
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
		Valid:  true,
	}
}
