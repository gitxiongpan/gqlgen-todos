package main

import (
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gitxiongpan/gqlgen-todos/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	schema := execSchema.Schema()
	newFields := ast.FieldList{}
	for _, field := range schema.Query.Fields {
		if field.Name != "todos" {
			newFields = append(newFields, field)
		}
	}
	schema.Query.Fields = newFields

	publicSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Schema:    schema,
		Resolvers: &graph.Resolver{},
	}))
	http.Handle("/public", playground.Handler("GraphQL playground", "/public/query"))
	http.Handle("/public/query", publicSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
