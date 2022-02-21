// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package api

import (
	"time"
)

// Defines values for PatchDocumentOp.
const (
	PatchDocumentOpAdd PatchDocumentOp = "add"

	PatchDocumentOpCopy PatchDocumentOp = "copy"

	PatchDocumentOpMove PatchDocumentOp = "move"

	PatchDocumentOpRemove PatchDocumentOp = "remove"

	PatchDocumentOpReplace PatchDocumentOp = "replace"

	PatchDocumentOpTest PatchDocumentOp = "test"
)

// The created time of this record
type Created time.Time

// Dnd5eMonster defines model for Dnd5eMonster.
type Dnd5eMonster struct {
	Alignment       *string   `binding:"max=40" json:"alignment,omitempty"`
	ArmorClass      *int      `json:"armorClass,omitempty"`
	ChallengeRating *string   `binding:"max=40" json:"challengeRating,omitempty"`
	Description     *string   `binding:"max=1024" json:"description,omitempty"`
	Dnd5eSettingId  string    `json:"dnd5eSettingId"`
	Environments    *[]string `binding:"max=16,dive,max=256" json:"environments,omitempty"`
	HitPoints       *int      `json:"hitPoints,omitempty"`
	Languages       *[]string `binding:"max=16,dive,max=256" json:"languages,omitempty"`
	Legendary       *bool     `json:"legendary,omitempty"`
	MonsterType     *string   `binding:"max=80" json:"monsterType,omitempty"`
	Name            string    `binding:"required,min=1,max=512" json:"name"`
	SizeCategory    *string   `binding:"max=80" json:"sizeCategory,omitempty"`
	Sources         *[]string `binding:"max=16,dive,max=256" json:"sources,omitempty"`
	Unique          *bool     `json:"unique,omitempty"`
	UserTags        UserTags  `binding:"required,max=64,dive,max=256" json:"userTags"`
}

// Dnd5eMonsterGet defines model for Dnd5eMonsterGet.
type Dnd5eMonsterGet struct {
	// Embedded struct due to allOf(#/components/schemas/Dnd5eMonster)
	Dnd5eMonster `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// The created time of this record
	Created        Created `json:"created"`
	Dnd5eMonsterId string  `json:"dnd5eMonsterId"`
	Version        int64   `json:"version"`
}

// Dnd5eMonsterPostOk defines model for Dnd5eMonsterPostOk.
type Dnd5eMonsterPostOk struct {
	Dnd5eMonsterId string `json:"dnd5eMonsterId"`
	Version        int64  `json:"version"`
}

// Dnd5eSetting defines model for Dnd5eSetting.
type Dnd5eSetting struct {
	Active       bool       `json:"active"`
	CommonAccess string     `binding:"required,oneof=private anyuser public" json:"commonAccess"`
	Description  string     `binding:"required,min=1,max=1024" json:"description"`
	Module       *string    `binding:"max=512" json:"module,omitempty"`
	Name         string     `binding:"required,min=1,max=512" json:"name"`
	SystemTags   SystemTags `binding:"required,max=64,dive,max=256" json:"systemTags"`
	UserTags     UserTags   `binding:"required,max=64,dive,max=256" json:"userTags"`
}

// Dnd5eSettingGet defines model for Dnd5eSettingGet.
type Dnd5eSettingGet struct {
	// Embedded struct due to allOf(#/components/schemas/Dnd5eSetting)
	Dnd5eSetting `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// The created time of this record
	Created        Created `json:"created"`
	Dnd5eSettingId string  `json:"dnd5eSettingId"`
	Version        int64   `json:"version"`
}

// Dnd5eSettingPostOk defines model for Dnd5eSettingPostOk.
type Dnd5eSettingPostOk struct {
	Dnd5eSettingId string `json:"dnd5eSettingId"`
	Version        int64  `json:"version"`
}

// Email defines model for Email.
type Email string

// The last modified time of this record
type LastModified time.Time

// A JSONPatch document as defined by RFC 6902
type PatchDocument []struct {
	// A string containing a JSON Pointer value.
	From *string `json:"from,omitempty"`

	// The operation to be performed
	Op PatchDocumentOp `json:"op"`

	// A JSON-Pointer
	Path string `json:"path"`

	// The value to be used within the operations.
	Value *map[string]interface{} `json:"value,omitempty"`
}

// The operation to be performed
type PatchDocumentOp string

// A 27 character string representing an unique id
type StringId string

// SystemTags defines model for SystemTags.
type SystemTags []string

// User defines model for User.
type User struct {
	Active       bool       `json:"active"`
	CommonAccess string     `binding:"required,oneof=private anyuser public" json:"commonAccess"`
	Email        Email      `binding:"required,email,min=1,max=512" json:"email"`
	Name         string     `binding:"required,min=1,max=512" json:"name"`
	SystemTags   SystemTags `binding:"required,max=64,dive,max=256" json:"systemTags"`
	UserTags     UserTags   `binding:"required,max=64,dive,max=256" json:"userTags"`
}

// UserGet defines model for UserGet.
type UserGet struct {
	// Embedded struct due to allOf(#/components/schemas/User)
	User `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// The created time of this record
	Created Created `json:"created"`
	UserId  string  `json:"userId"`
	Uuid    string  `binding:"required,uuid" json:"uuid"`
	Version int64   `json:"version"`
}

// UserTags defines model for UserTags.
type UserTags []string

// Dnd5eSettingId defines model for Dnd5eSettingId.
type Dnd5eSettingId int64

// Dnd5eSettingName defines model for Dnd5eSettingName.
type Dnd5eSettingName string

// UserId defines model for UserId.
type UserId string

// UserUUID defines model for UserUUID.
type UserUUID string

// Dnd5eMonsterArrayGetOk defines model for Dnd5eMonsterArrayGetOk.
type Dnd5eMonsterArrayGetOk []Dnd5eMonsterGet

// Dnd5eMonsterGetOk defines model for Dnd5eMonsterGetOk.
type Dnd5eMonsterGetOk Dnd5eMonsterGet

// Dnd5eSettingArrayGetOk defines model for Dnd5eSettingArrayGetOk.
type Dnd5eSettingArrayGetOk []Dnd5eSettingGet

// Dnd5eSettingGetOk defines model for Dnd5eSettingGetOk.
type Dnd5eSettingGetOk Dnd5eSettingGet

// UserGetOk defines model for UserGetOk.
type UserGetOk UserGet

// UserPostOk defines model for UserPostOk.
type UserPostOk struct {
	UserId  string `json:"userId"`
	Version int64  `json:"version"`
}

// A JSONPatch document as defined by RFC 6902
type Patch PatchDocument

// GetApiV1Dnd5eSettingsParams defines parameters for GetApiV1Dnd5eSettings.
type GetApiV1Dnd5eSettingsParams struct {
	// The beginning of the name
	Name *Dnd5eSettingName `json:"name,omitempty"`
}

// PostApiV1Dnd5eSettingsJSONBody defines parameters for PostApiV1Dnd5eSettings.
type PostApiV1Dnd5eSettingsJSONBody Dnd5eSetting

// PostApiV1Dnd5eSettingsDnd5eSettingIdJSONBody defines parameters for PostApiV1Dnd5eSettingsDnd5eSettingId.
type PostApiV1Dnd5eSettingsDnd5eSettingIdJSONBody Dnd5eSetting

// PostApiV1UsersJSONBody defines parameters for PostApiV1Users.
type PostApiV1UsersJSONBody User

// PostApiV1Dnd5eSettingsJSONRequestBody defines body for PostApiV1Dnd5eSettings for application/json ContentType.
type PostApiV1Dnd5eSettingsJSONRequestBody PostApiV1Dnd5eSettingsJSONBody

// PostApiV1Dnd5eSettingsDnd5eSettingIdJSONRequestBody defines body for PostApiV1Dnd5eSettingsDnd5eSettingId for application/json ContentType.
type PostApiV1Dnd5eSettingsDnd5eSettingIdJSONRequestBody PostApiV1Dnd5eSettingsDnd5eSettingIdJSONBody

// PostApiV1UsersJSONRequestBody defines body for PostApiV1Users for application/json ContentType.
type PostApiV1UsersJSONRequestBody PostApiV1UsersJSONBody

// PatchApiV1UsersUserIdJSONRequestBody defines body for PatchApiV1UsersUserId for application/json ContentType.
type PatchApiV1UsersUserIdJSONRequestBody Patch
