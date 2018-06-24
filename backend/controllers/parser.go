package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type graphQLParser struct{}

func (gqlp graphQLParser) Parse(req *http.Request) (request, error) {
	var (
		parsedReq request
		err       error
	)

	defer req.Body.Close()

	switch req.Method {
	case "POST":
		parsedReq, err = parsePost(req)
	case "GET":
		parsedReq, err = parseGet(req)
	default:
		err = errors.New("only POST and GET requests are supported")
	}

	if len(parsedReq.queries) == 0 {
		err = errors.New("no queries to execute")
	}

	return parsedReq, err
}

func parseGet(req *http.Request) (request, error) {
	var (
		value        = req.URL.Query()
		queries      = value["query"]
		names        = value["operationName"]
		variables    = value["variables"]
		queryLen     = len(queries)
		namesLen     = len(names)
		variablesLen = len(variables)

		requests = make([]query, 0, queryLen)
		isBatch  bool
	)

	if queryLen == 0 {
		return request{}, nil
	}

	for i, parsedQuery := range queries {
		var (
			parsedVars   = map[string]interface{}{}
			parsedOpName string
		)

		if i < namesLen {
			parsedOpName = names[i]
		}

		if i < variablesLen {
			str := variables[i]
			if err := json.Unmarshal([]byte(str), &parsedVars); err != nil {
				parsedVars = nil
			}
		}

		requests = append(requests, query{Query: parsedQuery, OpName: parsedOpName, Variables: parsedVars})
	}

	if queryLen > 1 {
		isBatch = true
	}

	return request{queries: requests, isBatch: isBatch}, nil
}

func parsePost(req *http.Request) (request, error) {
	var (
		queries []query
		isBatch bool
	)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return request{}, errors.New("unable to read request body")
	}

	if len(body) == 0 {
		return request{}, nil
	}

	switch body[0] {
	case '{':
		parsedQuery := query{}
		err := json.Unmarshal(body, &parsedQuery)
		if err == nil {
			queries = append(queries, parsedQuery)
		}
	case '[':
		isBatch = true
		_ = json.Unmarshal(body, &queries)
	}

	return request{queries: queries, isBatch: isBatch}, nil
}
