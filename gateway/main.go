package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/siddheshRajendraNimbalkar/bookstore/gateway/graphql"
)

func main() {
	fmt.Println("gateway service is running")

	resolver := graphql.NewResolver()

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("GraphQL Gateway running on :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
