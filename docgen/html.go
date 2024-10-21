package docgen

import (
	"fmt"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
	"github.com/samjove/gopendoc/parser"
)

// GenerateHTML generates the API documentation as HTML with embedded search.
func GenerateHTML(apis []parser.APIMetadata, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start writing HTML content.
	htmlContent := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>API Documentation</title>"

	// Add any CSS for styling.
	htmlContent += `<style>
		body { font-family: Arial, sans-serif; margin: 20px; }
		h1, h2, h3 { color: #333; }
		input#searchBar { padding: 10px; width: 100%; margin-bottom: 20px; }
		#searchResults p { margin: 5px 0; }
	</style>`

	// Close the head and start the body.
	htmlContent += "</head><body><h1>API Documentation</h1>"

	// Add search bar and search script using Fuse.js.
	htmlContent += `
<input type="text" id="searchBar" placeholder="Search API...">
<div id="searchResults"></div>

<script src="https://cdn.jsdelivr.net/npm/fuse.js@6.5.3"></script>
<script>
const searchBar = document.getElementById('searchBar');
const resultsContainer = document.getElementById('searchResults');
const apis = [
`

	// Prepare the API data for the search script.
	for _, api := range apis {
		anchor := strings.ReplaceAll(api.Path, "/", "-")
		htmlContent += fmt.Sprintf(`{ "method": "%s", "path": "%s", "anchor": "%s" },`, api.Method, api.Path, anchor)
	}

	// Finish the search script.
	htmlContent += `
];

searchBar.addEventListener('input', function() {
	const query = searchBar.value.toLowerCase();
	const results = apis.filter(api => 
		api.path.toLowerCase().includes(query) || 
		api.method.toLowerCase().includes(query)
	);
	resultsContainer.innerHTML = '';
	results.forEach(result => {
		resultsContainer.innerHTML += '<p><a href="#' + result.anchor + '">' + result.method + ' ' + result.path + '</a></p>';
	});
});
</script>
`

	// Add API documentation content.
	htmlContent += "<h2>Table of Contents</h2><ul>"
	for _, api := range apis {
		anchor := strings.ReplaceAll(api.Path, "/", "-")
		htmlContent += fmt.Sprintf(`<li><a href="#%s">%s %s</a></li>`, anchor, api.Method, api.Path)
	}
	htmlContent += "</ul>"

	for _, api := range apis {
		anchor := strings.ReplaceAll(api.Path, "/", "-")
		htmlContent += fmt.Sprintf(`<h2 id="%s">%s %s</h2>`, anchor, api.Method, api.Path)
		if api.Summary != "" {
			htmlContent += fmt.Sprintf("<p><i>%s</i></p>", api.Summary)
		}

		if len(api.Params) > 0 {
			htmlContent += "<h3>Parameters</h3><ul>"
			for _, param := range api.Params {
				htmlContent += fmt.Sprintf("<li><b>%s</b> (%s, %s) - %s</li>", param.Name, param.In, param.Type, param.Description)
			}
			htmlContent += "</ul>"
		}

		if len(api.Responses) > 0 {
			htmlContent += "<h3>Responses</h3><ul>"
			for _, resp := range api.Responses {
				htmlContent += fmt.Sprintf("<li><b>%d</b>: %s - %s</li>", resp.Status, resp.Type, resp.Description)
			}
			htmlContent += "</ul>"
		}
	}

	htmlContent += "</body></html>"

	// Convert markdown to HTML using blackfriday (in case there are any markdown elements).
	htmlOutput := blackfriday.Run([]byte(htmlContent))

	// Write the final HTML content to the output file.
	_, err = file.Write(htmlOutput)
	if err != nil {
		return err
	}

	return nil
}
