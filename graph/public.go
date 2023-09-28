package graph

import (
	"embed"
	"fmt"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

//go:embed "public.graphqls"
var publicSourcesFS embed.FS

func publicSourceData(filename string) string {
	data, err := publicSourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

func GetPublicSchema() *ast.Schema {
	var publicSources = []*ast.Source{
		{Name: "public.graphqls", Input: publicSourceData("public.graphqls"), BuiltIn: false},
	}
	return gqlparser.MustLoadSchema(publicSources...)
}
