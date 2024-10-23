package docgen

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/samjove/gopendoc/parser"
)

// TestGenerateHTMLSuccess tests the successful generation of an HTML file.
func TestGenerateHTMLSuccess(t *testing.T) {
	apis := []parser.APIMetadata{
		{
			Path:    "/api/v1/example",
			Method:  "GET",
			Summary: "Example API",
			Params: []parser.APIParam{
				{Name: "id", In: "query", Type: "string", Required: true, Description: "ID of the example"},
			},
			Responses: []parser.APIResponse{
				{Status: 200, Type: "application/json", ClassName: "ExampleResponse", Description: "Successful response"},
			},
		},
	}

	// Create a temporary file to store the generated HTML.
	tmpFile, err := os.CreateTemp("", "test_generate_html_*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Run the function
	err = GenerateHTML(apis, tmpFile.Name())
	if err != nil {
		t.Fatalf("GenerateHTML returned an error: %v", err)
	}

	// Read the content of the file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Check if the generated file contains expected HTML elements.
	if !strings.Contains(string(content), "<title>API Documentation</title>") {
		t.Errorf("Generated HTML does not contain the expected title")
	}

	if !strings.Contains(string(content), "<h1>API Documentation</h1>") {
		t.Errorf("Generated HTML does not contain the expected header")
	}

	if !strings.Contains(string(content), "GET-/api/v1/example") {
		t.Errorf("Generated HTML does not contain the expected API path and method")
	}

	if !strings.Contains(string(content), "Example API") {
		t.Errorf("Generated HTML does not contain the expected API summary")
	}

	if !strings.Contains(string(content), "(query, string)") {
		t.Errorf("Generated HTML does not contain the expected API parameter")
	}

	if !strings.Contains(string(content), "Successful response") {
		t.Errorf("Generated HTML does not contain the expected API response")
	}
}

// TestGenerateHTMLEmptyAPI tests the generation of HTML when no APIs are provided.
func TestGenerateHTMLEmptyAPI(t *testing.T) {
	var apis []parser.APIMetadata

	// Create a temporary file to store the generated HTML.
	tmpFile, err := ioutil.TempFile("", "test_generate_html_empty_*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after the test.

	// Run the function
	err = GenerateHTML(apis, tmpFile.Name())
	if err != nil {
		t.Fatalf("GenerateHTML returned an error: %v", err)
	}

	// Read the content of the file.
	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Check if the generated file contains expected HTML elements.
	if !strings.Contains(string(content), "<h1>API Documentation</h1>") {
		t.Errorf("Generated HTML does not contain the expected header")
	}

	if strings.Contains(string(content), "<div class=\"api-item\">") {
		t.Errorf("Generated HTML should not contain API items when no APIs are provided")
	}
}

// TestGenerateHTMLError tests the error handling when an invalid file path is provided.
func TestGenerateHTMLError(t *testing.T) {
	apis := []parser.APIMetadata{
		{
			Path:    "/api/v1/example",
			Method:  "GET",
			Summary: "Example API",
		},
	}

	// Use an invalid file path
	err := GenerateHTML(apis, "/invalid_path/output.html")
	if err == nil {
		t.Error("Expected an error but got none")
	}
}
