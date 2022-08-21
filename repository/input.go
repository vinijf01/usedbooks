package repository

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ProductFrom struct {
	Image       string `json:"image"`
	Title       string `form:"title"`
	Writer      string `form:"writer"`
	Price       int    `form:"price"`
	Description string `form:"description"`
	User        Users
}

type ProductFromUpdate struct {
	ID_Product  int    `json:"id_product"`
	Image       string `json:"image"`
	Title       string `form:"title"`
	Writer      string `form:"writer"`
	Price       int    `form:"price"`
	Description string `form:"description"`
	User        Users
}

type CartForm struct {
	Product Products
	User    Users
}

type WishlistForm struct {
	Product Products
	User    Users
}
