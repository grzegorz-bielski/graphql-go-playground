package main

import (
	"time"
	"fmt"
	"net/http"

	"github.com/graph-gophers/graphql-go"

	"grzegorz-bielski/microstream/backend/loaders"
	"grzegorz-bielski/microstream/backend/resolvers"
	"grzegorz-bielski/microstream/backend/schema"
	"grzegorz-bielski/microstream/backend/services"
)

func main() {
	rootService := services.NewRootService()
	rootResolver := resolvers.NewRootResolver(rootService)
	rootLoader := loaders.NewRootLoader(rootService)

	schema := graphql.MustParseSchema(schema.String(), rootResolver)

	router := http.NewServeMux()
	router.Handle("/graphiq", controllers.NewGraphiQLController())
	router.Handle("/graphql", controllers.NewGraphQLController(schema, rootLoader))

	if err = &http.Server{
		Addr: ":8000",
		Handler: router,
		ReadHeaderTimeout: 1 * time.Second,
		writeTimeout: 10 * time.Second,
		idleTimeout: 90 * time.Second,
		maxHeaderBytes: http.DefaultMaxHeaderBytes
	}.ListenAndServe(), err != nil {
		log.Println("Serve error", err)
	}
}
