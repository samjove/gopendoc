package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseGoFile tests the ParseGoFile function
func TestParseGoFile(t *testing.T) {
	// Test data simulating a Go file with comments
	content := `package main

// @route GET /users
// @summary Get user by ID
// @param id path int true "The user ID"
// @response 200 {object} User "Successful operation"
// @response 404 {string} string "User not found"
func GetUser() {
	// function body
}`

	// Create a temporary file for testing
	tempFile := t.TempDir() + "/testfile.go"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	assert.NoError(t, err)

	// Call ParseGoFile
	apis, err := ParseGoFile(tempFile)
	assert.NoError(t, err)
	assert.Len(t, apis, 1, "Expected 1 API to be parsed")

	api := apis[0]
	assert.Equal(t, "/users", api.Path)
	assert.Equal(t, "GET", api.Method)
	assert.Equal(t, "GetUser", api.Func)
	assert.Equal(t, "Get user by ID", api.Summary)

	// Check parameters
	assert.Len(t, api.Params, 1)
	assert.Equal(t, "id", api.Params[0].Name)
	assert.Equal(t, "path", api.Params[0].In)
	assert.Equal(t, "int", api.Params[0].Type)
	assert.True(t, api.Params[0].Required)
	assert.Equal(t, "The user ID", api.Params[0].Description)

	// Check responses
	assert.Len(t, api.Responses, 2)
	assert.Equal(t, 200, api.Responses[0].Status)
	assert.Equal(t, "object", api.Responses[0].Type)
	assert.Equal(t, "User", api.Responses[0].ClassName)
	assert.Equal(t, "Successful operation", api.Responses[0].Description)

	assert.Equal(t, 404, api.Responses[1].Status)
	assert.Equal(t, "string", api.Responses[1].Type)
	assert.Equal(t, "string", api.Responses[1].ClassName)
	assert.Equal(t, "User not found", api.Responses[1].Description)
}

// TestParseRoute tests the parseRoute function
func TestParseRoute(t *testing.T) {
	api := &APIMetadata{}
	comment := "// @route GET /users"
	parseRoute(comment, api)
	assert.Equal(t, "GET", api.Method)
	assert.Equal(t, "/users", api.Path)
}

// TestParseSummary tests the parseSummary function
func TestParseSummary(t *testing.T) {
	api := &APIMetadata{}
	comment := "// @summary Get user by ID"
	parseSummary(comment, api)
	assert.Equal(t, "Get user by ID", api.Summary)
}

// TestParseParam tests the parseParam function
func TestParseParam(t *testing.T) {
	api := &APIMetadata{}
	comment := `// @param id path int true "The user ID"`
	parseParam(comment, api)
	assert.Len(t, api.Params, 1)
	param := api.Params[0]
	assert.Equal(t, "id", param.Name)
	assert.Equal(t, "path", param.In)
	assert.Equal(t, "int", param.Type)
	assert.True(t, param.Required)
	assert.Equal(t, "The user ID", param.Description)
}

// TestParseResponse tests the parseResponse function
func TestParseResponse(t *testing.T) {
	api := &APIMetadata{}
	comment := `// @response 200 {object} User "Successful operation"`
	parseResponse(comment, api)
	assert.Len(t, api.Responses, 1)
	response := api.Responses[0]
	assert.Equal(t, 200, response.Status)
	assert.Equal(t, "object", response.Type)
	assert.Equal(t, "User", response.ClassName)
	assert.Equal(t, "Successful operation", response.Description)
}

