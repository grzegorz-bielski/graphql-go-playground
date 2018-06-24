package controllers

import (
	"grzegorz-bielski/microstream/backend/loaders"
	"grzegorz-bielski/microstream/backend/schema"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

// GraphQLController handles graphql requests
type GraphQLController struct {
	schema     *graphql.Schema
	rootLoader loaders.RootLoader
	parser     Parser
}

func NewGraphQLController(schema: *graphql.Schema, rootLoader loaders.RootLoader, parser Parser) *GraphQLController {
	return &GraphQLController{
		schema: *graphql.Schema,
		rootLoader: rootLoader
		parser: parser
	}
}

func (c GraphQLController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	parsedReq, err := c.parser.Parse(req)
	if err != nil {
		respond(res, jsonErr(err.Error()), http.StatusBadRequest)
		return
	}

	// auth here
}
