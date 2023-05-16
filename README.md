# Final Project Studi Independent
 # Project UsedBooks

This is a backend project for a used books e-commerce website. The project is built using the Go programming language and SQLite database.

## API Endpoints

### User
- `POST /user/register`: Register a new user.
- `POST /user/login`: Login a user.
- `POST /user/logout`: Logout a user.
- `GET /user/:id`: Get user information by user ID.

### Product
- `GET /product/list`: Get a list of all available products.
- `GET /product/detail/:id`: Get product details by product ID.
- `POST /product/insert`: Add a new product.
- `PATCH /product/update/:id`: Update product information by product ID.
- `DELETE /product/delete/:id`: Delete a product by product ID.

### Cart
- `POST /cart/add/:id`: Add a product to the cart.
- `GET /cart/list`: Get a list of products in the cart.
- `DELETE /cart/delete/:id`: Remove a product from the cart by product ID.

### Wishlist
- `POST /wishlist/add/:id`: Add a product to the wishlist.
- `GET /wishlist/list`: Get a list of products in the wishlist.
- `DELETE /wishlist/delete/:id`: Remove a product from the wishlist by product ID.

## How to Run the Project

1. Clone the repository to your local machine.
2. Install Go and SQLite if you haven't already.
3. Navigate to the project directory in your terminal.
4. Run `go run main.go` to start the server.
5. Access the API endpoints using a client such as Postman.

