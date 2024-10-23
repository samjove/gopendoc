# Gopendoc

### API Documentation Generator
Gopendoc is a lightweight tool that generates HTML API documentation for Go projects based on code annotations. It scans your Go files for specific comments (e.g., // @route, // @summary, // @param, // @response), generates user-friendly HTML pages with search functionality, and serves them in your local environment.

#### Features
- Parses Go files to extract API metadata from annotations.
- Supports routes, parameters, and responses using tags like @route, @param, and @response.
- Generates a complete HTML document with collapsible sections for each API endpoint.
- Enables viewing with search functionality to filter API endpoints.
- Minimal dependencies and easy integration into Go projects.

#### Installation
Ensure that Go is installed and set up in your environment.

To install the tool, use the following command:

`go get -u github.com/samjove/gopendoc`



#### Usage
Once installed, you can generate API documentation using the following command:

`gopendoc gen`

This will scan your project and generate HTML files into "./apidocs" by default.

Alternatively, you can specify a root directory and an output directory with the -directory (-d) and -output (-o) flags, respectively.

You can serve the generated files for reference in your browser with the following command:

`gopendoc serve`

#### Supported Annotations
The following annotations are supported:
```
@route <method> <path>
@summary <description>
@param <name> <in> <type> <required> <description>
@response <status> <type> <className> <description>
```

##### Example
Here’s an example of how to annotate your Go code for the tool:

```
// @route GET /users
// @summary Get user by ID
// @param id path int true "User ID"
// @response 200 {object} User "Successful response"
// @response 404 {string} string "User not found"
func GetUser(w http.ResponseWriter, r *http.Request) {
    // implementation here
}
```
This will generate an HTML section for the GET /users endpoint, listing its parameters and responses.

The "examples" directory contains additional examples and the "apidocs" directory holds the generated HTML.

#### Output
Gopendoc generates an HTML file in the specified output directory. The generated HTML includes:

- A list of API routes, parameters, and responses.
- Collapsible sections for each API route.
- A search bar for quickly finding API routes.

#### Contributing
Contributions are welcome! If you find any issues or have ideas for improvement, feel free to open an issue or submit a pull request.

Typical contribution process:
- Fork the repository
- Create a new branch
- Commit your changes
- Push to the branch
- Open a pull request

#### License
This project is licensed under the MIT License. See the LICENSE file for details.