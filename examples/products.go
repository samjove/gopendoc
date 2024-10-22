package examples

// @route GET /products/{id}
// @summary Get a product by ID
// @param id path int true "The ID of the product"
// @response 200 {object} Product "OK"
// @response 404 {object} ErrorResponse "Product not found"
func GetProduct() {
	// Logic to get a product
}

// @route POST /products
// @summary Create a new product
// @param none
// @response 201 {object} Product "Product successfully created"
// @response 400 {object} ErrorResponse "Invalid request"
func PostProduct() {
	// Logic to create a new product
}

// @route PUT /products/{id}
// @summary Update a product by ID
// @param id path int true "The ID of the product"
// @param body body Product true "Updated product data"
// @response 200 {object} Product "Product successfully updated"
// @response 404 {object} ErrorResponse "Product not found"
func UpdateProduct() {
	// Logic to update a product
}

// @route DELETE /products/{id}
// @summary Delete a product by ID
// @param id path int true "The ID of the product"
// @response 204 {string} string "Product successfully deleted"
// @response 404 {object} ErrorResponse "Product not found"
func DeleteProduct() {
	// Logic to delete a product
}