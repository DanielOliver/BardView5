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

func mapDnd5eWorldToJsonBody(dnd5eWorld *db.Dnd5eWorld) *api.Dnd5eWorldGet {
	ret := &api.Dnd5eWorldGet{
		Dnd5eWorld: api.Dnd5eWorld{
			Active:       dnd5eWorld.IsActive,
			CommonAccess: dnd5eWorld.CommonAccess,
			Module:       nil,
			Name:         dnd5eWorld.Name,
			Description:  dnd5eWorld.Description,
			SystemTags:   dnd5eWorld.SystemTags,
			UserTags:     dnd5eWorld.UserTags,
		},
		Created:      api.Created(dnd5eWorld.CreatedAt.Format(time.RFC3339)),
		Dnd5eWorldId: strconv.FormatInt(dnd5eWorld.Dnd5eWorldID, 10),
		Version:      dnd5eWorld.Version,
	}
	if dnd5eWorld.Module.Valid {
		ret.Dnd5eWorld.Module = &dnd5eWorld.Module.String
	}
	return ret
}

type GetDnd5eWorldByIdParams struct {
	Dnd5eWorldId int64 `uri:"dnd5eWorldId" binding:"required"`
}

func GetDnd5eWorldById(b *BardView5Http) {
	var params GetDnd5eWorldByIdParams
	if err := b.Context.ShouldBindUri(&params); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dnd5eWorld, err := Dnd5eWorldById(b, params.Dnd5eWorldId)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	err = Dnd5eWorldHasAccess(b, &dnd5eWorld)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", strconv.FormatInt(dnd5eWorld.Version, 10))
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/world/%d/", dnd5eWorld.Dnd5eWorldID))
	b.Context.JSON(http.StatusOK, mapDnd5eWorldToJsonBody(&dnd5eWorld))
}

func GetMyDnd5eWorlds(b *BardView5Http) {
	dnd5eWorlds, err := b.Querier().Dnd5eWorldFindByAssignment(b.Context, b.Session.SessionId)

	if err != nil {
		b.Logger.Err(err).Int64("id", b.Session.SessionId).Msg("Failed to get mine dnd5eworld")
		b.Context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	results := make([]api.Dnd5eWorldGet, len(dnd5eWorlds))
	for i, world := range dnd5eWorlds {
		results[i] = *mapDnd5eWorldToJsonBody(&world)
	}

	b.Context.JSON(http.StatusOK, results)
}

func PostDnd5eWorldsCreate(b *BardView5Http) {
	var body api.PostApiV1Dnd5eWorldsJSONBody
	if err := b.Context.ShouldBindJSON(&body); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newDnd5eWorldId, err := Dnd5eWorldCreate(b, &body)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", "0")
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/world/%d/", newDnd5eWorldId))
	b.Context.JSON(http.StatusCreated, api.Dnd5eWorldPostOk{
		Dnd5eWorldId: strconv.FormatInt(newDnd5eWorldId, 10),
		Version:      0,
	})
}

func Dnd5eWorldById(b *BardView5Http, dnd5eWorldId int64) (db.Dnd5eWorld, error) {
	dnd5eWorlds, err := b.Querier().Dnd5eWorldFindById(b.Context, dnd5eWorldId)

	empty := db.Dnd5eWorld{}
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eWorld)
		return empty, ErrFailedRead(ObjDnd5eWorld, dnd5eWorldId, true)
	}
	if len(dnd5eWorlds) == 0 {
		return empty, ErrNotFound(ObjDnd5eWorld, dnd5eWorldId)
	}
	return dnd5eWorlds[0], nil
}

func Dnd5eWorldHasAccess(b *BardView5Http, dnd5eWorld *db.Dnd5eWorld) error {
	switch dnd5eWorld.CommonAccess {
	case CommonAccessPublic:
		return nil
	case CommonAccessAnyUser:
		if b.Session.Anonymous {
			return ErrNotAuthorized(ObjDnd5eWorld, dnd5eWorld.Dnd5eWorldID)
		}
		return nil
	case CommonAccessPrivate:
		worldAssignments, err := b.Querier().Dnd5eWorldFindAssignment(b.Context, db.Dnd5eWorldFindAssignmentParams{
			UserID:       b.Session.SessionId,
			Dnd5eWorldID: dnd5eWorld.Dnd5eWorldID,
		})
		if err != nil {
			fmt.Println(err.Error())
			b.Logger.Err(err).Msg(ObjDnd5eWorldAssignment)
			return ErrFailedRead(ObjDnd5eWorldAssignment, dnd5eWorld.Dnd5eWorldID, true)
		}
		if len(worldAssignments) == 0 {
			return ErrNotAuthorized(ObjDnd5eWorld, dnd5eWorld.Dnd5eWorldID)
		}
		return nil
	default:
		return ErrNotAuthorized(ObjDnd5eWorld, dnd5eWorld.Dnd5eWorldID)
	}
}

func Dnd5eWorldCreate(b *BardView5Http, body *api.PostApiV1Dnd5eWorldsJSONBody) (int64, error) {
	newDnd5eWorldId := b.GenDnd5eWorld().Generate().Int64()
	changedRows, err := b.Querier().Dnd5eWorldInsert(b.Context, db.Dnd5eWorldInsertParams{
		Dnd5eWorldID: newDnd5eWorldId,
		CommonAccess: body.CommonAccess,
		CreatedBy:    MaybeInt64(&b.Session.SessionId),
		IsActive:     body.Active,
		SystemTags:   body.SystemTags,
		UserTags:     body.UserTags,
		Name:         body.Name,
		Module:       MaybeString(body.Module),
		Description:  body.Description,
	})
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eWorld)
		return 0, ErrFailedCreate(ObjDnd5eWorld, newDnd5eWorldId, true)
	}
	if changedRows == 0 {
		return 0, ErrUnknownStatusCreate(ObjDnd5eWorld, newDnd5eWorldId, true)
	}

	_, err = b.Querier().Dnd5eWorldUpsertAssignment(b.Context, db.Dnd5eWorldUpsertAssignmentParams{
		CreatedBy:    MaybeInt64(&b.Session.SessionId),
		UserID:       b.Session.SessionId,
		Dnd5eWorldID: newDnd5eWorldId,
		RoleAction:   RoleActionOwner,
	})
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eWorldAssignment)
		return 0, ErrFailedCreate(ObjDnd5eWorldAssignment, newDnd5eWorldId, true)
	}

	return newDnd5eWorldId, nil
}
