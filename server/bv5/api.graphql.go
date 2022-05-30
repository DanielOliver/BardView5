package bv5

import (
	"context"
	_ "embed"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
	"server/bardlog"
)

//go:embed "bv5.graphql"
var graphqlSchema string

type graphqlQuery struct {
	bv5 *BardView5
}

func (g *graphqlQuery) session(c context.Context) *BardView5Graphql {
	return &BardView5Graphql{
		BardView5: g.bv5,
		Logger:    bardlog.GetLoggerFromContext(c),
		Session:   *SessionCriteria(c),
		Context:   c,
	}
}

func (g *graphqlQuery) Setting(ctx context.Context, args struct{ ID graphql.ID }) (*settingResolver, error) {
	return &settingResolver{bv5: g.session(ctx)}, nil
}

func GraphqlHandler(bv5 *BardView5) http.Handler {
	opts := []graphql.SchemaOpt{graphql.MaxDepth(4)}
	schema := graphql.MustParseSchema(graphqlSchema, &graphqlQuery{bv5: bv5}, opts...)
	return &relay.Handler{Schema: schema}
}

type settingResolver struct {
	bv5 *BardView5Graphql
}

func (s *settingResolver) ID() graphql.ID {
	return "1"
}

func (s *settingResolver) Name() string {
	return "example"
}
