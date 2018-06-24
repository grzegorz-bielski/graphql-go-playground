package controllers

import "net/http"

// Parser parsers HTTP GraphQL requests
type Parser interface {
	Parse(req *http.Request) (request, error)
}

type responder interface {
	respond(res http.ResponseWriter, body []byte, code int)
}

type query struct {
	OpName    string                 `json:"operationName"`
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type request struct {
	queries []query
	isBatch bool
}
