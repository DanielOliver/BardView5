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

func mapDnd5eMonsterToJsonBody(m *db.Dnd5eMonster) *api.Dnd5eMonsterGet {
	ret := &api.Dnd5eMonsterGet{
		Dnd5eMonster: api.Dnd5eMonster{
			Alignment:            SMaybeString(m.Alignment),
			ArmorClass:           SMaybeInt32(m.ArmorClass),
			MilliChallengeRating: SMaybeInt64(m.MilliChallengeRating),
			Description:          SMaybeString(m.Description),
			Environments:         &m.Environments,
			HitPoints:            SMaybeInt32(m.HitPoints),
			Languages:            &m.Environments,
			Legendary:            &m.IsLegendary,
			MonsterType:          SMaybeString(m.MonsterType),
			Name:                 m.Name,
			SizeCategory:         SMaybeString(m.SizeCategory),
			Sources:              &m.Sources,
			Unique:               &m.IsUnique,
			UserTags:             m.UserTags,
		},
		Created:        api.Created(m.CreatedAt.Format(time.RFC3339)),
		Dnd5eMonsterId: strconv.FormatInt(m.Dnd5eMonsterID, 10),
		Dnd5eSettingId: strconv.FormatInt(m.Dnd5eSettingID, 10),
		Version:        m.Version,
	}
	return ret
}

func ApiGetDnd5eMonstersBySettingId(b *BardView5Http) {
	var params GetDnd5eSettingByIdParams
	if err := b.Context.ShouldBindUri(&params); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dnd5eSetting, err := dnd5eSettingById(b, params.Dnd5eSettingId)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	err = dnd5eSettingHasAccess(b, &dnd5eSetting)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	monsters, _ := b.Querier().Dnd5eMonstersFindBySetting(b.Context, db.Dnd5eMonstersFindBySettingParams{
		Dnd5eSettingID: params.Dnd5eSettingId,
		RowOffset:      0,
		RowLimit:       1000,
	})

	results := make([]api.Dnd5eMonsterGet, len(monsters))
	for i, monster := range monsters {
		results[i] = *mapDnd5eMonsterToJsonBody(&monster)
	}

	b.Context.Header("ETag", strconv.FormatInt(dnd5eSetting.Version, 10))
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/setting/%d/monsters", dnd5eSetting.Dnd5eSettingID))
	b.Context.JSON(http.StatusOK, api.Dnd5eMonsterArrayGetOk(results))
}

func ApiPostDnd5eMonstersCreate(b *BardView5Http) {
	var params GetDnd5eSettingByIdParams
	if err := b.Context.ShouldBindUri(&params); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body api.PostApiV1Dnd5eSettingsDnd5eSettingIdMonstersJSONBody
	if err := b.Context.ShouldBindJSON(&body); err != nil {
		b.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dnd5eSettingId := params.Dnd5eSettingId
	dnd5eSetting, err := dnd5eSettingById(b, dnd5eSettingId)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	err = dnd5eSettingHasAccess(b, &dnd5eSetting)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	newDnd5eSettingId, err := dnd5eMonsterCreate(b, &body, dnd5eSettingId)
	if err != nil {
		WriteErrorToContext(b, err)
		return
	}

	b.Context.Header("ETag", "0")
	b.Context.Header("Location", fmt.Sprintf("/v1/dnd5e/monsters/%d/", newDnd5eSettingId))
	b.Context.JSON(http.StatusCreated, api.Dnd5eSettingPostOk{
		Dnd5eSettingId: strconv.FormatInt(newDnd5eSettingId, 10),
		Version:        0,
	})
}

func dnd5eMonsterCreate(b *BardView5Http, body *api.PostApiV1Dnd5eSettingsDnd5eSettingIdMonstersJSONBody, dnd5eSettingId int64) (int64, error) {
	newDnd5eMonsterId := b.GenDnd5eMonster().Generate().Int64()
	changedRows, err := b.Querier().Dnd5eMonsterInsert(b.Context, db.Dnd5eMonsterInsertParams{
		Dnd5eMonsterID:       newDnd5eMonsterId,
		CreatedBy:            MaybeInt64(&b.Session.SessionId),
		Dnd5eSettingID:       dnd5eSettingId,
		Name:                 body.Name,
		Sources:              DefaultStringArr(body.Sources),
		UserTags:             body.UserTags,
		Languages:            DefaultStringArr(body.Languages),
		Environments:         DefaultStringArr(body.Environments),
		IsLegendary:          DefaultBool(body.Legendary, false),
		IsUnique:             DefaultBool(body.Unique, false),
		MonsterType:          MaybeString(body.MonsterType),
		Alignment:            MaybeString(body.Alignment),
		SizeCategory:         MaybeString(body.SizeCategory),
		MilliChallengeRating: MaybeInt64(body.MilliChallengeRating),
		ArmorClass:           MaybeInt(body.ArmorClass),
		HitPoints:            MaybeInt(body.HitPoints),
		Description:          MaybeString(body.Description),
	})
	if err != nil {
		b.Logger.Err(err).Msg(ObjDnd5eMonster)
		return 0, ErrFailedCreate(ObjDnd5eMonster, newDnd5eMonsterId, true)
	}
	if changedRows == 0 {
		return 0, ErrUnknownStatusCreate(ObjDnd5eMonster, newDnd5eMonsterId, true)
	}

	return newDnd5eMonsterId, nil
}
