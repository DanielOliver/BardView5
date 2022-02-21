package bv5

import (
	"database/sql"
	"strconv"
)

func SMaybeInt32(v sql.NullInt32) *int {
	if v.Valid {
		value := int(v.Int32)
		return &value
	}
	return nil
}

func SMaybeInt64(v sql.NullInt64) *int64 {
	if v.Valid {
		value := v.Int64
		return &value
	}
	return nil
}

func SMaybeString(v sql.NullString) *string {
	if v.Valid {
		value := v.String
		return &value
	}
	return nil
}

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

func MaybeInt(value *int) sql.NullInt32 {
	if value == nil {
		return sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
	}
	return sql.NullInt32{
		Int32: int32(*value),
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

func DefaultBool(value *bool, d bool) bool {
	if value == nil {
		return d
	}
	return *value
}

func DefaultStringArr(value *[]string) []string {
	if value == nil {
		return []string{}
	}
	return *value
}
