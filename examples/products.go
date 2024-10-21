package examples

// @route GET /products/{id}
// @summary Get a Product by ID
// @param id path int true "The ID of the Product"
// @response 200 {object} Product "OK"
// @response 404 {object} ErrorResponse "Product not found"
func GetProduct() {
	// Logic to get a Product
}

// @route POST /products
// @summary Post a new Product
// @param none
// @response 201 {object} Success "Product successfully created"
// @response 403 {object} ErrorResponse "Bad Request"
func PostProduct() {
	// Logic to post a Product
}
