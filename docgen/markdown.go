package docgen

import (
	"fmt"
	"os"

	"github.com/samjove/gopendoc/parser"
)

// GenerateMarkdown generates API documentation in Markdown format.
func GenerateMarkdown(apis []parser.APIMetadata, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("# API Documentation\n\n")
	if err != nil {
		return err
	}
	for _, api := range apis {
		_, err := file.WriteString(fmt.Sprintf("## %s %s\n", api.Method, api.Path))
		if err != nil {
			return err
		}
		_, err = file.WriteString(fmt.Sprintf("Function: `%s`\n\n", api.Func))
		if err != nil {
			return err
		}
	}
	return nil
}