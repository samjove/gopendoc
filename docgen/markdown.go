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

		if api.Summary != "" {
			_, err = file.WriteString(fmt.Sprintf("_%s_\n\n", api.Summary))
			if err != nil {
				return err
			}
		}

		if len(api.Params) > 0 {
			_, err = file.WriteString("### Parameters\n\n")
			if err != nil {
				return err
			}
			for _, param := range api.Params {
				_, err = file.WriteString(fmt.Sprintf("- **%s** (%s, %s) - %s\n", param.Name, param.In, param.Type, param.Description))
				if err != nil {
					return err
				}
			}
			_, err = file.WriteString("\n")
		}

		if len(api.Responses) > 0 {
			_, err = file.WriteString("### Responses\n\n")
			if err != nil {
				return err
			}
			for _, resp := range api.Responses {
				_, err = file.WriteString(fmt.Sprintf("- **%d**: %s - %s\n", resp.Status, resp.Type, resp.Description))
				if err != nil {
					return err
				}
			}
			_, err = file.WriteString("\n")
		}
	}

	return nil
}
