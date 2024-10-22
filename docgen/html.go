package docgen

import (
	"fmt"
	"os"
	"strings"

	"github.com/samjove/gopendoc/parser"
)

// GenerateHTML generates the API documentation as HTML with embedded JavaScript and search functionality.
func GenerateHTML(apis []parser.APIMetadata, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start writing HTML content
	htmlContent := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>API Documentation</title>"

	// Add CSS for styling
	htmlContent += `<style>
		body { font-family: Arial, sans-serif; margin: 20px; }
		h1, h2, h3 { color: #333; }
		input#searchBar { padding: 10px; width: 100%; margin-bottom: 20px; }
		#api-list h2 { margin-top: 20px; }
		#searchResults p { margin: 5px 0; }
	</style>`

	// Close the head and start the body
	htmlContent += "</head><body><h1>API Documentation</h1>"

	// Add search bar and search script
	htmlContent += `
<input type="text" id="searchBar" placeholder="Search API...">
<div id="searchResults"></div>

<div id="api-list">
`

	// Generate HTML for the API list
	for _, api := range apis {
		fmt.Printf("api.Path before replace: %s\n", api.Path)
		anchor := strings.ReplaceAll(api.Path, "/", "-")
		fmt.Printf("api.Path: %s\n", api.Path)
		fmt.Printf("api Summary: %s\n", api.Summary)
		htmlContent += fmt.Sprintf(`<div class="api-item" data-method="%s" data-path="%s" data-summary="%s">
			<h2 id="%s">%s %s</h2>`, api.Method, api.Path, api.Summary, anchor, api.Method, api.Path)

		if api.Summary != "" {
			htmlContent += fmt.Sprintf(`<p><i>%s</i></p>`, api.Summary)
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

		htmlContent += "</div>" // Close .api-item
	}

	htmlContent += `</div>` // Close #api-list

	// Add the search filtering JavaScript
	htmlContent += `
<script>
	const searchBar = document.getElementById('searchBar');
	const apiItems = document.querySelectorAll('.api-item');
	console.log(apiItems)

	searchBar.addEventListener('input', function() {
		const query = searchBar.value.toLowerCase();

		apiItems.forEach(item => {
			const method = item.getAttribute('data-method').toLowerCase();
			const path = item.getAttribute('data-path').toLowerCase();
			const summary = item.getAttribute('data-summary').toLowerCase();

			if (method.includes(query) || path.includes(query) || summary.includes(query)) {
				item.style.display = '';  // Show matching items
			} else {
				item.style.display = 'none';  // Hide non-matching items
			}
		});
	});
</script>
`

	// Close HTML
	htmlContent += "</body></html>"

	// Write the final HTML content to the output file
	_, err = file.WriteString(htmlContent)
	if err != nil {
		return err
	}

	return nil
}
