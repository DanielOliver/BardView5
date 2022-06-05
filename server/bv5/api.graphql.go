package bv5

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
	"server/bardlog"
	"server/db"
	"strconv"
	"time"
)

//go:embed "bv5.graphql"
var graphqlSchema string

type graphqlQuery struct {
	b *BardView5
}

func (g *graphqlQuery) session(c context.Context) *BardView5Graphql {
	return &BardView5Graphql{
		BardView5: g.b,
		Logger:    bardlog.GetLoggerFromContext(c),
		Session:   *SessionCriteria(c),
		Context:   c,
	}
}

func gqIdTo64(id graphql.ID) (int64, error) {
	return strconv.ParseInt(string(id), 10, 64)
}

func (g *graphqlQuery) Setting(ctx context.Context, args struct{ ID graphql.ID }) (*settingResolver, error) {
	id, err := gqIdTo64(args.ID)
	if err != nil {
		return nil, ErrNotFoundGq(ObjDnd5eSetting, args.ID)
	}

	b := g.session(ctx)

	setting, err := dnd5eSettingById2(b, id)
	if err != nil {
		return nil, err
	}
	if err = dnd5eSettingHasAccess2(b, &setting); err != nil {
		return nil, err
	}

	return &settingResolver{b: g.session(ctx), setting: &setting}, nil
}

func GraphqlHandler(bv5 *BardView5) http.Handler {
	opts := []graphql.SchemaOpt{graphql.MaxDepth(4)}
	schema := graphql.MustParseSchema(graphqlSchema, &graphqlQuery{b: bv5}, opts...)
	return &relay.Handler{Schema: schema}
}

type settingResolver struct {
	b       *BardView5Graphql
	setting *db.Dnd5eSetting
}

func (s *settingResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(s.setting.Dnd5eSettingID, 10))
}
func (s *settingResolver) Name() string {
	return s.setting.Name
}
func (s *settingResolver) Version() string {
	return strconv.FormatInt(s.setting.Version, 10)
}
func (s *settingResolver) CreatedDate() string {
	return s.setting.CreatedAt.Format(time.RFC3339)
}
func (s *settingResolver) CommonAccess() string {
	return s.setting.CommonAccess
}
func (s *settingResolver) Active() bool {
	return s.setting.IsActive
}
func (s *settingResolver) Module() *string {
	return SMaybeString(s.setting.Module)
}
func (s *settingResolver) Description() string {
	return s.setting.Description
}
func (s *settingResolver) Monsters() ([]*monsterResolver, error) {
	monsters, err := s.b.Q().Dnd5eMonstersFindBySetting(s.b.C(), db.Dnd5eMonstersFindBySettingParams{
		Dnd5eSettingID: s.setting.Dnd5eSettingID,
		RowOffset:      0,
		RowLimit:       1000000,
	})
	if err != nil {
		return nil, err
	}
	results := make([]*monsterResolver, len(monsters))
	for i, monster := range monsters {
		m := monster
		results[i] = &monsterResolver{
			b:       s.b,
			monster: &m,
		}
	}
	return results, nil
}

type monsterResolver struct {
	b       *BardView5Graphql
	monster *db.Dnd5eMonster
}

func (m *monsterResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(m.monster.Dnd5eSettingID, 10))
}
func (m *monsterResolver) Name() string {
	return m.monster.Name
}
func (m *monsterResolver) MonsterType() *string {
	return SMaybeString(m.monster.MonsterType)
}

func dnd5eSettingById2(b WebRequest, dnd5eSettingId int64) (db.Dnd5eSetting, error) {
	dnd5eSettings, err := b.Q().Dnd5eSettingFindById(b.C(), dnd5eSettingId)

	empty := db.Dnd5eSetting{}
	if err != nil {
		b.L().Err(err).Msg(ObjDnd5eSetting)
		return empty, ErrFailedRead(ObjDnd5eSetting, dnd5eSettingId, true)
	}
	if len(dnd5eSettings) == 0 {
		return empty, ErrNotFound(ObjDnd5eSetting, dnd5eSettingId)
	}
	return dnd5eSettings[0], nil
}

func dnd5eSettingHasAccess2(b WebRequest, dnd5eSetting *db.Dnd5eSetting) error {
	switch dnd5eSetting.CommonAccess {
	case CommonAccessPublic:
		return nil
	case CommonAccessAnyUser:
		if b.S().Anonymous {
			return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
		}
		return nil
	case CommonAccessPrivate:
		settingAssignments, err := b.Q().RoleAssignmentFindByScopeId(b.C(), db.RoleAssignmentFindByScopeIdParams{
			UserID:      b.S().SessionId,
			ScopeID:     dnd5eSetting.Dnd5eSettingID,
			RoleSubject: ObjDnd5eSetting,
		})
		if err != nil {
			fmt.Println(err.Error())
			b.L().Err(err).Msg(ObjRoleAssignment)
			return ErrFailedRead(ObjRoleAssignment, dnd5eSetting.Dnd5eSettingID, true)
		}
		if len(settingAssignments) == 0 {
			return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
		}
		return nil
	default:
		return ErrNotAuthorized(ObjDnd5eSetting, dnd5eSetting.Dnd5eSettingID)
	}
}
