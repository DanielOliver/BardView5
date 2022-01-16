package bv5

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api"
	"server/db"
	"time"
)

func mapDnd5eWorldToJsonBody(dnd5eWorld *db.Dnd5eWorld) *api.Dnd5eWorldGet {
	ret := &api.Dnd5eWorldGet{
		Dnd5eWorld: api.Dnd5eWorld{
			Active:           dnd5eWorld.IsActive,
			CommonAccess:     dnd5eWorld.CommonAccess,
			DerivedFromWorld: nil,
			Module:           nil,
			Name:             dnd5eWorld.Name,
			Description:      dnd5eWorld.Description,
			SystemTags:       dnd5eWorld.SystemTags,
			UserTags:         dnd5eWorld.UserTags,
		},
		Created:      api.Created(dnd5eWorld.CreatedAt.Format(time.RFC3339)),
		Dnd5eWorldId: dnd5eWorld.Dnd5eWorldID,
		Version:      dnd5eWorld.Version,
	}
	if dnd5eWorld.DerivedFromWorld.Valid {
		ret.Dnd5eWorld.DerivedFromWorld = &dnd5eWorld.DerivedFromWorld.Int64
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

	dnd5eWorlds, err := b.Querier().Dnd5eWorldFindById(b.Context, params.Dnd5eWorldId)

	if err != nil {
		b.Logger.Err(err).Int64("id", params.Dnd5eWorldId).Msg("Failed to get dnd5eworld")
		b.Context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(dnd5eWorlds) == 0 {
		b.Context.AbortWithStatus(http.StatusNotFound)
		return
	}

	dnd5eWorld := dnd5eWorlds[0]
	createReturn := func() {
		b.Context.JSON(http.StatusOK, mapDnd5eWorldToJsonBody(&dnd5eWorld))
	}

	if dnd5eWorld.CommonAccess == CommonAccessPublic {
		createReturn()
		return
	}
	if b.Session.Anonymous {
		b.Context.AbortWithStatus(http.StatusNotFound)
		return
	}

	if dnd5eWorld.CommonAccess == CommonAccessAnyUser && !b.Session.Anonymous {
		createReturn()
		return
	}

	if dnd5eWorld.CommonAccess != CommonAccessPrivate {
		b.Context.AbortWithStatus(http.StatusNotFound)
		return
	}

	worldAssignments, err := b.Querier().Dnd5eWorldFindAssignment(b.Context, db.Dnd5eWorldFindAssignmentParams{
		UserID:       b.Session.sessionId,
		Dnd5eWorldID: dnd5eWorld.Dnd5eWorldID,
	})
	if len(worldAssignments) == 0 || err != nil {
		b.Context.AbortWithStatus(http.StatusNotFound)
		return
	}

	createReturn()
}

func GetMyDnd5eWorlds(b *BardView5Http) {
	dnd5eWorlds, err := b.Querier().Dnd5eWorldFindByAssignment(b.Context, b.Session.sessionId)

	if err != nil {
		b.Logger.Err(err).Int64("id", b.Session.sessionId).Msg("Failed to get mine dnd5eworld")
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
	newDnd5eWorldId := b.GenDnd5eWorld().Generate().Int64()
	changedRows, err := b.Querier().Dnd5eWorldInsert(b.Context, db.Dnd5eWorldInsertParams{
		Dnd5eWorldID:     newDnd5eWorldId,
		DerivedFromWorld: MaybeInt64(body.DerivedFromWorld),
		CommonAccess:     body.CommonAccess,
		CreatedBy:        MaybeInt64(&b.Session.sessionId),
		IsActive:         body.Active,
		SystemTags:       body.SystemTags,
		UserTags:         body.UserTags,
		Name:             body.Name,
		Module:           MaybeString(body.Module),
		Description:      body.Description,
	})
	if err != nil {
		b.Logger.Err(err).Msg("Failed to create new dnd5eworld")
		b.Context.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to create new dnd5eworld")
		return
	}
	if changedRows == 0 {
		b.Context.JSON(http.StatusBadRequest, "Failed to create new dnd5eworld")
		return
	}

	_, err = b.Querier().Dnd5eWorldUpsertAssignment(b.Context, db.Dnd5eWorldUpsertAssignmentParams{
		CreatedBy:    MaybeInt64(&b.Session.sessionId),
		UserID:       b.Session.sessionId,
		Dnd5eWorldID: newDnd5eWorldId,
		RoleAction:   RoleActionOwner,
	})
	if err != nil {
		b.Logger.Err(err).Msg("Failed to assign new dnd5eworld")
		b.Context.AbortWithStatusJSON(http.StatusInternalServerError, "Troubled creating new dnd5eworld")
		return
	}

	b.Context.Header("ETag", "0")
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/world/%d/", newDnd5eWorldId))
	b.Context.JSON(http.StatusCreated, api.UserPostOk{
		UserId:  newDnd5eWorldId,
		Version: 0,
	})
}
