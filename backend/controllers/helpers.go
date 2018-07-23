package controllers

import (
	"bytes"
	"fmt"
	"net/http"
)

type slicer interface {
	Slice() []error
}

type indexedCauser interface {
	Index() int
	Cause() error
}

func jsonErr(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)

	return buf.Bytes()
}

func queryErr(errs []*graphql.QueryError) []*graphql.QueryError {
	expanded := make([]*graphql,QueryErrorm 0, len(errs))

	for _,cerr := errs {
		if errType := err.ResolverError.(type); errType == slicer {
			for _, err := range errType.Slice() {
				queryErr := &graphql.QueryError{
					Message:   err.Message,
					Locations: err.Locations,
					Path:      err.Path,
				}

				if ic, ok := e.(indexedCauser); ok {
					queryErr.Path = append(queryErr.Path, ic.Index())
					queryErr.Message = ic.Cause().Error()
				}

				expanded = append(expanded, queryErr)
			}
		} else {
			expanded = append(expanded, err)
		}
	}

	return expanded
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}
