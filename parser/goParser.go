package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

// APIMetadata holds information about a single API endpoint.
type APIMetadata struct {
	Path      string
	Method    string
	Func      string
	Summary   string
	Params    []APIParam
	Responses []APIResponse
}

// APIParam represents a parameter in an API route.
type APIParam struct {
	Name        string
	In          string
	Type        string
	Required    bool
	Description string
}

// APIResponse represents a response from an API.
type APIResponse struct {
	Status      int
	Type        string
	ClassName   string
	Description string
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
			api := APIMetadata{
				Func: fn.Name.Name,
			}

			for _, comment := range fn.Doc.List {
				commentText := strings.TrimSpace(comment.Text)
				if strings.HasPrefix(commentText, "// @route") {
					parseRoute(commentText, &api)
				}
				if strings.HasPrefix(commentText, "// @summary") {
					parseSummary(commentText, &api)
				}
				if strings.HasPrefix(commentText, "// @param") {
					parseParam(commentText, &api)
				}
				if strings.HasPrefix(commentText, "// @response") {
					parseResponse(commentText, &api)
				}
			}

			if api.Path != "" && api.Method != "" {
				apis = append(apis, api)
			}
		}
		return true
	})

	return apis, nil
}

// parseRoute parses the @route tag for method and path.
func parseRoute(comment string, api *APIMetadata) {
	fmt.Printf("comment before route parse: %s\n", comment)
	parts := strings.Split(comment, " ")
	if len(parts) >= 3 {
		api.Method = parts[2]
		api.Path = parts[3]
	}
}

// parseSummary parses the @summary tag for API description.
func parseSummary(comment string, api *APIMetadata) {
	api.Summary = strings.TrimSpace(strings.TrimPrefix(comment, "// @summary"))
}

// parseParam parses the @param tag for API parameters
func parseParam(comment string, api *APIMetadata) {
	paramRegex := regexp.MustCompile(`@param (\w+) (path|query) (\w+) (true|false) "(.*)"`)
	matches := paramRegex.FindStringSubmatch(comment)
	if len(matches) == 6 {
		param := APIParam{
			Name:        matches[1],
			In:          matches[2],
			Type:        matches[3],
			Required:    matches[4] == "true",
			Description: matches[5],
		}
		api.Params = append(api.Params, param)
	}
}

// parseResponse parses the @response tag for API responses.
func parseResponse(comment string, api *APIMetadata) {
	responseRegex := regexp.MustCompile(`@response (\d{3}) \{(\w+)\} (\w+) "(.*)"`)
	matches := responseRegex.FindStringSubmatch(comment)
	if len(matches) == 5 {
		response := APIResponse{
			Status:      strToInt(matches[1]),
			Type:        matches[2],
			ClassName:   matches[3],
			Description: matches[4],
		}
		api.Responses = append(api.Responses, response)
	}
}

// Converts string to int.
func strToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
