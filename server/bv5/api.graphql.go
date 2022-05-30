package bv5

import (
	_ "embed"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

//go:embed "bv5.graphql"
var graphqlSchema string

type graphqlQuery struct{}

func (_ *graphqlQuery) Hello() string { return "Hello, world!" }

func GraphqlHandler() http.Handler {
	opts := []graphql.SchemaOpt{graphql.MaxDepth(4)}
	schema := graphql.MustParseSchema(graphqlSchema, &graphqlQuery{}, opts...)
	return &relay.Handler{Schema: schema}
}
