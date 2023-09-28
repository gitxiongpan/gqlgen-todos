package graph

import (
	"embed"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutablePublicSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutablePublicSchema(cfg Config) graphql.ExecutableSchema {
	es := &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
		schema:     getPublicSchema(),
	}
	return es
}

//go:embed "public.graphqls"
var publicSourcesFS embed.FS

var publicSources = []*ast.Source{
	{Name: "public.graphqls", Input: publicSourceData("public.graphqls"), BuiltIn: false},
}

func publicSourceData(filename string) string {
	data, err := publicSourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

func getPublicSchema() *ast.Schema {
	return gqlparser.MustLoadSchema(publicSources...)
}
