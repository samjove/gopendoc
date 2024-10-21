package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	port      int
	docDir    string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the generated API documentation via an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting server at http://localhost:%d\n", port)
		fmt.Printf("Serving documentation from: %s\n", docDir)

		// Check if the documentation directory exists.
		if _, err := os.Stat(docDir); os.IsNotExist(err) {
			log.Fatalf("Documentation directory does not exist: %s\n", docDir)
		}

		// Serve static files from the documentation directory.
		http.Handle("/", http.FileServer(http.Dir(docDir)))

		// Start the HTTP server.
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to serve documentation")
	serveCmd.Flags().StringVarP(&docDir, "dir", "d", "./apidocs", "Directory containing the generated documentation")
}
