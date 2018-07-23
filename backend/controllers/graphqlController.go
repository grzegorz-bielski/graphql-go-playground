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

func (c GraphQLController) execQueries(parsedReq request, ctx context.Context) []*graphql.Response{
	var (
		queriesNum = len(parsedReq.queries)
		responses = make([]*graphql.Response, queriesNum)
		wg        sync.WaitGroup 
	)

	wg.Add(queriesNum)
	for i, parsedQuery := range parsedReq.queries {
		go func(i int, parsedQuery query) {
			defer wg.Done()
			response := c.Schema.Exec(ctx, parsedQuery.Query, parsedQuery.OpName, parsedQuery.Variables)
			response.Errors = queryErr(response.Errors)
			responses[i] = response
		}(i, parsedQuery)
	}
	wg.Wait()

	return responses
}

func (c GraphQLController) stringifyResponse(response []*graphql.Response, isBatch bool) ([]byte, error){
	var resp []byte
	if req.isBatch {
		return json.Marshal(responses)
	} else if len(responses) > 0 {
		return json.Marshal(responses[0])
	}
}

func (c GraphQLController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	parsedReq, err := c.parser.Parse(req)
	if err != nil {
		respond(res, jsonErr(err.Error()), http.StatusBadRequest)
		return
	}

	// auth here

	ctx = c.rootLoader.Attach(req.Context())
	responses := c.execQueries(parsedReq, ctx)
	stringifiedResponse := c.stringifyResponse(responses, parsedReq.isBatch)

	if err != nil {
		respond(res, jsonErr("server error"), http.StatusInternalServerError)
		return
	}

	respond(res, stringifiedResponse , http.StatusOK)
}
