package examples

// @route GET /users/{id}
// @summary Get a user by ID
// @param id path int true "The ID of the user"
// @response 200 {object} User "OK"
// @response 404 {object} ErrorResponse "User not found"
func GetUser() {
	// Logic to get a user
}

// @route POST /users
// @summary Post a new user
// @param none
// @response 201 {object} Success "User successfully created"
// @response 403 {object} ErrorResponse "Bad Request"
func PostUser() {
	// Logic to post a user
}
