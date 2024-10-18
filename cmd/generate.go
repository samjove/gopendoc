package cmd

import (
	"fmt"

	"github.com/samjove/gopendoc/docgen"
	"github.com/samjove/gopendoc/parser"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
)

var generateCmd = &cobra.Command{
	Use: "gen",
	Short: "Generate API documentation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating API documentation...")

		apis, err := parser.ParseGoFile(inputFile)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return
		}
		err = docgen.GenerateMarkdown(apis, outputFile)
		if err != nil {
			fmt.Println("Error generating documentation:", err)
			return
		}
		fmt.Println("Documentation generated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input Go file")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "api.md", "Output file for generated documentation")
}