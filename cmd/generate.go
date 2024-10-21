package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/samjove/gopendoc/docgen"
	"github.com/samjove/gopendoc/parser"
	"github.com/spf13/cobra"
)

var (
	rootDir  string
	outputDir string
)

var generateCmd = &cobra.Command{
	Use: "gen",
	Short: "Generate API documentation for all Go files in the directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating API documentation...")

		err := walkDirectory(rootDir)
		if err != nil {
			fmt.Println("Error walking directory:", err)
			return
		}
		fmt.Println("Documentation generated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&rootDir, "directory", "d", ".", "Root directory to scan for Go files")
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "./apidocs", "Directory to store the generated documentation")
}

// Function to walk through all files in the directory and subdirectories
func walkDirectory(rootDir string) error {
   return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Only process .go files
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
            fmt.Println("Parsing file:", path)
            apis, err := parser.ParseGoFile(path)
			if err != nil {
				fmt.Printf("Error parsing file %s: %v\n", path, err)
				return err
			}

			if _, err := os.Stat(outputDir); os.IsNotExist(err) {
				err := os.MkdirAll(outputDir, os.ModePerm)
				if err != nil {
					fmt.Printf("Error creating output directory: %v\n", err)
					return err
				}
			}
			
			outputFile := filepath.Join(outputDir, strings.TrimSuffix(info.Name(), ".go")+".html")

			// Generate HTML documentation.
			if len(apis) != 0 {
				err = docgen.GenerateHTML(apis, outputFile)
				if err != nil {
					fmt.Printf("Error generating documentation for file %s: %v\n", path, err)
					return err
				}
				fmt.Printf("Documentation generated for %s -> %s\n", path, outputFile)
			}
        }
        return nil
    })
}