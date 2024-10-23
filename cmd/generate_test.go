package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/samjove/gopendoc/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for the Parser.
type MockParser struct {
	mock.Mock
}

func (m *MockParser) ParseGoFile(file string) ([]parser.APIMetadata, error) {
	args := m.Called(file)
	return args.Get(0).([]parser.APIMetadata), args.Error(1)
}

// Mock for the Doc Generator.
type MockDocGenerator struct {
	mock.Mock
}

func (m *MockDocGenerator) GenerateHTML(apis []parser.APIMetadata, outputFile string) error {
	args := m.Called(apis, outputFile)
	err := os.WriteFile(outputFile, []byte("<html><body>Sample API Documentation</body></html>"), 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	return args.Error(0)
}

func TestGenerateAPIDocs(t *testing.T) {
	// Create temporary root and output directories for testing
	rootDir := t.TempDir()
	outputDir := filepath.Join(rootDir, "apidocs")
	err := os.Mkdir(outputDir, os.ModePerm)
	assert.NoError(t, err, "Expected output directory to be created")

	// Create sample .go file in the rootDir
	sampleFile := filepath.Join(rootDir, "sample.go")
	err = os.WriteFile(sampleFile, []byte("// @route GET /users"), 0644)
	assert.NoError(t, err, "Expected sample.go to be created")

	// Mock the parser and doc generator
	mockParser := new(MockParser)
	mockDocGen := new(MockDocGenerator)

	// Define expected API metadata
	expectedAPIs := []parser.APIMetadata{
		{
			Method:  "GET",
			Path:    "/users",
			Summary: "Get user by ID",
		},
	}

	// Set expectations for ParseGoFile and GenerateHTML
	mockParser.On("ParseGoFile", sampleFile).Return(expectedAPIs, nil)

	// Use mock.MatchedBy to accept any file path for the output HTML
	mockDocGen.On("GenerateHTML", expectedAPIs, mock.MatchedBy(func(outputFile string) bool {
		return filepath.Base(outputFile) == "sample.html"
	})).Return(nil)

	// Override the global parser and doc generator instances
	parserInstance = mockParser.ParseGoFile
	docgenInstance = mockDocGen.GenerateHTML

	// Run the generateAPIDocs logic
	err = walkDirectory(rootDir)
	assert.NoError(t, err, "Expected walkDirectory to complete without error")

	// Verify that ParseGoFile and GenerateHTML were called with correct arguments
	mockParser.AssertCalled(t, "ParseGoFile", sampleFile)
	mockDocGen.AssertCalled(t, "GenerateHTML", expectedAPIs, mock.MatchedBy(func(outputFile string) bool {
		return filepath.Base(outputFile) == "sample.html"
	}))

	t.Cleanup(func() {
		err = os.RemoveAll("./apidocs")
		assert.NoError(t, err, "Failed to clean up sample.html file")
	})
}
