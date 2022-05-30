package bv5

import (
	"context"
	_ "embed"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

//go:embed "bv5.graphql"
var graphqlSchema string

type graphqlQuery struct{}

func (_ *graphqlQuery) Setting(ctx context.Context, args struct{ ID graphql.ID }) (*settingResolver, error) {
	return &settingResolver{}, nil
}

func GraphqlHandler() http.Handler {
	opts := []graphql.SchemaOpt{graphql.MaxDepth(4)}
	schema := graphql.MustParseSchema(graphqlSchema, &graphqlQuery{}, opts...)
	return &relay.Handler{Schema: schema}
}

type settingResolver struct {
}

func (s *settingResolver) ID() graphql.ID {
	return "1"
}

func (s *settingResolver) Name() string {
	return "example"
}
