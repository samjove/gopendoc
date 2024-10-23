package docgen

import (
	"fmt"
	"os"

	"github.com/samjove/gopendoc/parser"
)

// GenerateHTML generates the API documentation as HTML with embedded JavaScript and search functionality.
func GenerateHTML(apis []parser.APIMetadata, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(`
    <!DOCTYPE html>
    <html>
    <head>
        <title>API Documentation</title>
		<style>
            .api-item {
				margin-bottom: 15px;
				padding: 10px;
				border: 1px solid #686868;
				border-radius: 5px;
				background-color: #fefbea;
				font-family: sans-serif;
			}
			.api-header {
				cursor: pointer;
				background-color: #f2ecd5;
				padding: 10px;
				font-weight: bold;
				display: flex;
				justify-content: space-between;
				align-items: center;
				border-bottom: 1px solid #dcdcdc;
			}

			.api-details {
				display: none;
				padding: 10px;
				background-color: #e8e2cc;
				border-top: 1px solid #ccc;
			}
            .method-tag {
                font-size: 14px;
                padding: 4px 8px;
                border-radius: 3px;
                color: black;
            }
            .method-get { background-color: #28a745; }
            .method-post { background-color: #007bff; }
            .method-put { background-color: #ffc107; }
            .method-delete { background-color: #dc3545; }
            .collapsible {
                display: none;
            }
        </style>
</head>
<body>
<a href="/">Home</a>
<h1>API Documentation</h1>
<input type="text" id="search-bar" placeholder="Search endpoints..." />
`)

	if err != nil {
		return err
	}

	for _, api := range apis {
		id := fmt.Sprintf("%s-%s", api.Method, api.Path) // create unique id for each section

		// Write the compact header (method, path, summary)
		_, err = file.WriteString(fmt.Sprintf(`
		<div class="api-item">
			<div class="api-header" onclick="toggleDetails('%s')">
				<span class="method-tag method-%s">%s</span> %s
				<span>%s</span>
			</div>
			<div class="api-details" id="%s">
		`, id, api.Method, api.Method, api.Path, api.Summary, id))

		if err != nil {
			return err
		}

		// Write parameters section
		if len(api.Params) > 0 {
			_, err = file.WriteString("<h4>Parameters:</h4><ul>")
			if err != nil {
				return err
			}
			for _, param := range api.Params {
				_, err = file.WriteString(fmt.Sprintf(`
				<li><b>%s</b> (%s, %s) - %s</li>`, param.Name, param.In, param.Type, param.Description))
				if err != nil {
					return err
				}
			}
			_, err = file.WriteString("</ul>")
			if err != nil {
				return err
			}
		}

		// Write responses section
		if len(api.Responses) > 0 {
			_, err = file.WriteString("<h4>Responses:</h4><ul>")
			if err != nil {
				return err
			}
			for _, resp := range api.Responses {
				_, err = file.WriteString(fmt.Sprintf(`
				<li><b>%d</b> - %s: %s</li>`, resp.Status, resp.Type, resp.Description))
				if err != nil {
					return err
				}
			}
			_, err = file.WriteString("</ul>")
			if err != nil {
				return err
			}
		}

		// Close API details
		_, err = file.WriteString("</div></div>")
		if err != nil {
			return err
		}
	}

	// Close HTML tags
	_, err = file.WriteString(`
<script>
		function toggleDetails(id) {
			var details = document.getElementById(id);
			if (details.style.display === 'block') {
				details.style.display = 'none';
			} else {
				details.style.display = 'block';
			}
		}
		</script>
		<script>
		function filterAPIs() {
        let input = document.getElementById('search-bar');
        let filter = input.value.toLowerCase();
        let items = document.getElementsByClassName('api-item');

        for (let i = 0; i < items.length; i++) {
            let header = items[i].getElementsByClassName('api-header')[0];
            let textValue = header.textContent || header.innerText;
            
            if (textValue.toLowerCase().indexOf(filter) > -1) {
                items[i].style.display = "";
            } else {
                items[i].style.display = "none";
            }
        }
    }

    // Attach the filter function to the search bar
    document.getElementById('search-bar').addEventListener('input', filterAPIs);
	</script>
</body>
</html>
`)
	if err != nil {
		return err
	}

	return nil
}
