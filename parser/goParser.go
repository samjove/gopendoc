package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// APIMetadata holds information about a single API endpoint.
type APIMetadata struct {
	Path string
	Method string
	Func string
}

// ParseGoFile parses a single Go file to extract API routes.
func ParseGoFile(filename string) ([]APIMetadata, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var apis []APIMetadata

	// Walk through AST.
	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				if strings.Contains(comment.Text, "@route") {
					api := extractAPIRoute(fn, comment.Text)
					apis = append(apis, api)
				}
			}
		}
		return true
	})

	return apis, nil
}

func extractAPIRoute(fn *ast.FuncDecl, comment string) APIMetadata {
	var api APIMetadata
	parts := strings.Split(comment, " ")

	// Assumes comment format is @route METHOD PATH
	if len(parts) >= 3 {
		api.Method = parts[1]
		api.Path = parts[2]
		api.Func = fn.Name.Name
	}

	return api
}