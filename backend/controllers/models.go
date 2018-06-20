package controllers

import "net/http"

type parser interface {
	parse(req *http.Request) (request, error)
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
