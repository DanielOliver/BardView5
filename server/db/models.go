// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CommonAccess struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type Dnd5eInhabitant struct {
	Dnd5eInhabitantID int64         `db:"dnd5e_inhabitant_id"`
	CreatedBy         sql.NullInt64 `db:"created_by"`
	CreatedAt         time.Time     `db:"created_at"`
	Version           int64         `db:"version"`
	UserTags          []string      `db:"user_tags"`
	SystemTags        []string      `db:"system_tags"`
	Dnd5eWorldID      int64         `db:"dnd5e_world_id"`
	Dnd5eMonsterID    int64         `db:"dnd5e_monster_id"`
	OriginalWorld     bool          `db:"original_world"`
}

type Dnd5eLanguage struct {
	Dnd5eLanguageID int64         `db:"dnd5e_language_id"`
	CreatedBy       sql.NullInt64 `db:"created_by"`
	CreatedAt       time.Time     `db:"created_at"`
	Version         int64         `db:"version"`
	Name            string        `db:"name"`
}

type Dnd5eMonster struct {
	Dnd5eMonsterID       int64         `db:"dnd5e_monster_id"`
	CreatedBy            sql.NullInt64 `db:"created_by"`
	CreatedAt            time.Time     `db:"created_at"`
	Version              int64         `db:"version"`
	FirstWorldID         sql.NullInt64 `db:"first_world_id"`
	Name                 string        `db:"name"`
	Tags                 []string      `db:"tags"`
	MonsterType          string        `db:"monster_type"`
	Alignment            string        `db:"alignment"`
	SizeCategory         string        `db:"size_category"`
	MilliChallengeRating int64         `db:"milli_challenge_rating"`
	Languages            []string      `db:"languages"`
	Description          string        `db:"description"`
}

type Dnd5eMonsterType struct {
	CreatedBy sql.NullInt64 `db:"created_by"`
	CreatedAt time.Time     `db:"created_at"`
	Version   int64         `db:"version"`
	Name      string        `db:"name"`
}

type Dnd5eSizeCategory struct {
	CreatedBy sql.NullInt64 `db:"created_by"`
	CreatedAt time.Time     `db:"created_at"`
	Version   int64         `db:"version"`
	Name      string        `db:"name"`
	Space     string        `db:"space"`
}

type Dnd5eWorld struct {
	Dnd5eWorldID     int64         `db:"dnd5e_world_id"`
	CreatedBy        sql.NullInt64 `db:"created_by"`
	CreatedAt        time.Time     `db:"created_at"`
	Version          int64         `db:"version"`
	IsActive         bool          `db:"is_active"`
	CommonAccess     string        `db:"common_access"`
	UserTags         []string      `db:"user_tags"`
	SystemTags       []string      `db:"system_tags"`
	DerivedFromWorld sql.NullInt64 `db:"derived_from_world"`
	Name             string        `db:"name"`
}

type Role struct {
	RoleID        int64         `db:"role_id"`
	CreatedBy     sql.NullInt64 `db:"created_by"`
	CreatedAt     time.Time     `db:"created_at"`
	EffectiveDate time.Time     `db:"effective_date"`
	EndDate       sql.NullTime  `db:"end_date"`
	IsActive      bool          `db:"is_active"`
	Name          string        `db:"name"`
	RoleTypeID    sql.NullInt64 `db:"role_type_id"`
	Tags          []string      `db:"tags"`
}

type RoleAction struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type RoleAssignment struct {
	RoleAssignmentID int64         `db:"role_assignment_id"`
	CreatedBy        sql.NullInt64 `db:"created_by"`
	CreatedAt        time.Time     `db:"created_at"`
	EffectiveDate    time.Time     `db:"effective_date"`
	EndDate          sql.NullTime  `db:"end_date"`
	IsActive         bool          `db:"is_active"`
	RoleID           int64         `db:"role_id"`
	UserID           int64         `db:"user_id"`
	Tags             []string      `db:"tags"`
}

type RolePermission struct {
	RolePermissionID int64           `db:"role_permission_id"`
	CreatedBy        sql.NullInt64   `db:"created_by"`
	CreatedAt        time.Time       `db:"created_at"`
	EffectiveDate    time.Time       `db:"effective_date"`
	EndDate          sql.NullTime    `db:"end_date"`
	IsActive         bool            `db:"is_active"`
	RoleID           int64           `db:"role_id"`
	Action           string          `db:"action"`
	Subject          string          `db:"subject"`
	SubjectID        sql.NullInt64   `db:"subject_id"`
	Conditions       json.RawMessage `db:"conditions"`
	Fields           []string        `db:"fields"`
}

type RoleSubject struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type RoleType struct {
	RoleTypeID                 int64     `db:"role_type_id"`
	CreatedAt                  time.Time `db:"created_at"`
	Name                       string    `db:"name"`
	MultipleAssignmentsAllowed bool      `db:"multiple_assignments_allowed"`
	SystemManaged              bool      `db:"system_managed"`
}

type SchemaMigration struct {
	Version int64 `db:"version"`
	Dirty   bool  `db:"dirty"`
}

type User struct {
	UserID        int64         `db:"user_id"`
	Uuid          uuid.UUID     `db:"uuid"`
	CreatedBy     sql.NullInt64 `db:"created_by"`
	CreatedAt     time.Time     `db:"created_at"`
	Version       int64         `db:"version"`
	EffectiveDate time.Time     `db:"effective_date"`
	EndDate       sql.NullTime  `db:"end_date"`
	IsActive      bool          `db:"is_active"`
	CommonAccess  string        `db:"common_access"`
	Email         string        `db:"email"`
	Name          string        `db:"name"`
	UserTags      []string      `db:"user_tags"`
	SystemTags    []string      `db:"system_tags"`
}
