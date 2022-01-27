package bv5

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api"
	"server/db"
	"strconv"
	"time"
)

func mapDnd5eSettingToJsonBody(dnd5eSetting *db.Dnd5eSetting) *api.Dnd5eSettingGet {
	ret := &api.Dnd5eSettingGet{
		Dnd5eSetting: api.Dnd5eSetting{
			Active:       dnd5eSetting.IsActive,
			CommonAccess: dnd5eSetting.CommonAccess,
			Module:       nil,
			Name:         dnd5eSetting.Name,
			Description:  dnd5eSetting.Description,
			SystemTags:   dnd5eSetting.SystemTags,
			UserTags:     dnd5eSetting.UserTags,
		},
		Created:        api.Created(dnd5eSetting.CreatedAt.Format(time.RFC3339)),
		Dnd5eSettingId: strconv.FormatInt(dnd5eSetting.Dnd5eSettingID, 10),
		Version:        dnd5eSetting.Version,
	}
	if dnd5eSetting.Module.Valid {
		ret.Dnd5eSetting.Module = &dnd5eSetting.Module.String
	}
	return ret
}

type GetDnd5eSettingByIdParams struct {
	Dnd5eSettingId int64 `uri:"dnd5eSettingId" binding:"required"`
}

func GetDnd5eSettingById(b *BardView5Http) {
	var params GetDnd5eSettingByIdParams
	if err := b.Context.ShouldBindUri(&params); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dnd5eSetting, err := Dnd5eSettingById(b, params.Dnd5eSettingId)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	err = Dnd5eSettingHasAccess(b, &dnd5eSetting)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", strconv.FormatInt(dnd5eSetting.Version, 10))
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/setting/%d/", dnd5eSetting.Dnd5eSettingID))
	b.Context.JSON(http.StatusOK, mapDnd5eSettingToJsonBody(&dnd5eSetting))
}

func GetMyDnd5eSettings(b *BardView5Http) {
	dnd5eSettings, err := b.Querier().Dnd5eSettingFindByAssignment(b.Context, b.Session.SessionId)

	if err != nil {
		b.Logger.Err(err).Int64("id", b.Session.SessionId).Msg("Failed to get mine dnd5esetting")
		b.Context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	results := make([]api.Dnd5eSettingGet, len(dnd5eSettings))
	for i, setting := range dnd5eSettings {
		results[i] = *mapDnd5eSettingToJsonBody(&setting)
	}

	b.Context.JSON(http.StatusOK, results)
}

func PostDnd5eSettingsCreate(b *BardView5Http) {
	var body api.PostApiV1Dnd5eSettingsJSONBody
	if err := b.Context.ShouldBindJSON(&body); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newDnd5eSettingId, err := Dnd5eSettingCreate(b, &body)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", "0")
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/setting/%d/", newDnd5eSettingId))
	b.Context.JSON(http.StatusCreated, api.Dnd5eSettingPostOk{
		Dnd5eSettingId: strconv.FormatInt(newDnd5eSettingId, 10),
		Version:        0,
	})
}

func Dnd5eSettingById(b *BardView5Http, dnd5eSettingId int64) (db.Dnd5eSetting, error) {
	dnd5eSettings, err := b.Querier().Dnd5eSettingFindById(b.Context, dnd5eSettingId)

	empty := db.Dnd5eSetting{}
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eSetting)
		return empty, ErrFailedRead(ObjDnd5eSetting, dnd5eSettingId, true)
	}
	if len(dnd5eSettings) == 0 {
		return empty, ErrNotFound(ObjDnd5eSetting, dnd5eSettingId)
	}
	return dnd5eSettings[0], nil
}

func Dnd5eSettingHasAccess(b *BardView5Http, dnd5eSetting *db.Dnd5eSetting) error {
	switch dnd5eSetting.CommonAccess {
	case CommonAccessPublic:
		return nil
	case CommonAccessAnyUser:
		if b.Session.Anonymous {
			return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
		}
		return nil
	case CommonAccessPrivate:
		settingAssignments, err := b.Querier().Dnd5eSettingFindAssignment(b.Context, db.Dnd5eSettingFindAssignmentParams{
			UserID:         b.Session.SessionId,
			Dnd5eSettingID: dnd5eSetting.Dnd5eSettingID,
		})
		if err != nil {
			fmt.Println(err.Error())
			b.Logger.Err(err).Msg(ObjDnd5eSettingAssignment)
			return ErrFailedRead(ObjDnd5eSettingAssignment, dnd5eSetting.Dnd5eSettingID, true)
		}
		if len(settingAssignments) == 0 {
			return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
		}
		return nil
	default:
		return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
	}
}

func Dnd5eSettingCreate(b *BardView5Http, body *api.PostApiV1Dnd5eSettingsJSONBody) (int64, error) {
	newDnd5eSettingId := b.GenDnd5eSetting().Generate().Int64()
	changedRows, err := b.Querier().Dnd5eSettingInsert(b.Context, db.Dnd5eSettingInsertParams{
		Dnd5eSettingID: newDnd5eSettingId,
		CommonAccess:   body.CommonAccess,
		CreatedBy:      MaybeInt64(&b.Session.SessionId),
		IsActive:       body.Active,
		SystemTags:     body.SystemTags,
		UserTags:       body.UserTags,
		Name:           body.Name,
		Module:         MaybeString(body.Module),
		Description:    body.Description,
	})
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eSetting)
		return 0, ErrFailedCreate(ObjDnd5eSetting, newDnd5eSettingId, true)
	}
	if changedRows == 0 {
		return 0, ErrUnknownStatusCreate(ObjDnd5eSetting, newDnd5eSettingId, true)
	}

	_, err = b.Querier().Dnd5eSettingUpsertAssignment(b.Context, db.Dnd5eSettingUpsertAssignmentParams{
		CreatedBy:      MaybeInt64(&b.Session.SessionId),
		UserID:         b.Session.SessionId,
		Dnd5eSettingID: newDnd5eSettingId,
		RoleAction:     RoleActionOwner,
	})
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eSettingAssignment)
		return 0, ErrFailedCreate(ObjDnd5eSettingAssignment, newDnd5eSettingId, true)
	}

	return newDnd5eSettingId, nil
}