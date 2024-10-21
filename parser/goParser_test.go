package parser

import (
	"reflect"
	"testing"
)

// Test function for parseResponse.
func TestParseResponse(t *testing.T) {
    tests := []struct {
        input  string
        expect APIResponse
    }{
        // Test case 1: Valid 200 OK response.
        {
            input: `@response 200 {object} User "OK"`,
            expect: APIResponse{
                Status:      200,
                Type:        "object",
                ClassName:   "User",
                Description: "OK",
            },
        },
        // Test case 2: Valid 404 Error response.
        {
            input: `@response 404 {object} ErrorResponse "User not found"`,
            expect: APIResponse{
                Status:      404,
                Type:        "object",
                ClassName:   "ErrorResponse",
                Description: "User not found",
            },
        },
        // Test case 3: Valid 500 Internal Server Error response.
        {
            input: `@response 500 {error} ServerError "Internal server error"`,
            expect: APIResponse{
                Status:      500,
                Type:        "error",
                ClassName:   "ServerError",
                Description: "Internal server error",
            },
        },
        // Test case 4: Valid 204 No content.
        {
            input: `@response 204 {none} NoContent "No content"`,
            expect: APIResponse{
                Status:      204,
                Type:        "none",
                ClassName:   "NoContent",
                Description: "No content",
            },
        },
        // Test case 5: Invalid case, should not match anything.
        {
            input:  `@response 201 {object} "Incomplete response"`,
            expect: APIResponse{}, // No response expected for malformed input.
        },
    }

    for _, test := range tests {
        // Initialize the API Metadata object.
        api := &APIMetadata{}

        // Call parseResponse to parse the input string.
        parseResponse(test.input, api)

        if len(api.Responses) == 0 && test.expect != (APIResponse{}) {
            t.Errorf("Expected a response but got none for input: %s", test.input)
        } else if len(api.Responses) > 0 {
            result := api.Responses[0]
            if !reflect.DeepEqual(result, test.expect) {
                t.Errorf("Expected %v, but got %v for input: %s", test.expect, result, test.input)
            }
        }
    }
}

// Test if string to int conversion works correctly.
func TestStrToInt(t *testing.T) {
    result := strToInt("200")
    if result != 200 {
        t.Errorf("Expected 200 but got %d", result)
    }
}

// Edge case: Missing or malformed tags
func TestMalformedResponse(t *testing.T) {
    malformedInputs := []string{
        `@response 300 {object} "No class name"`,           // Missing class name
        `@response abc {object} User "Invalid status code"`, // Invalid status code
        `@response 200 {object} User`,                      // No description
        `response 200 {object} User "OK"`,                  // Missing @ symbol
    }

    for _, input := range malformedInputs {
        api := &APIMetadata{}
        parseResponse(input, api)

        // Expect no responses to be parsed for malformed inputs
        if len(api.Responses) > 0 {
            t.Errorf("Did not expect any parsed response for malformed input: %s", input)
        }
    }
}
