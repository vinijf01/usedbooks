package repository

import "time"

type UserFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

//untuk mengubah response json sesuai denganformat
func FormatUserRegister(user Users) UserFormatter {
	formatter := UserFormatter{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return formatter
}

type LoginResponse struct {
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ProductResponse struct {
	ID_Product  int    `json:"id_product"`
	ID_User     int    `json:"id_user"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Writer      string `json:"writer"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func FormatProduct(h Products) ProductResponse {
	return ProductResponse{
		ID_Product:  h.ID_Product,
		ID_User:     h.ID_User,
		Image:       h.Image,
		Title:       h.Title,
		Writer:      h.Writer,
		Price:       h.Price,
		Description: h.Description,
	}
}

type CartResponse struct {
	ID_Product int `json:"id_product"`
	ID_User    int `json:"id_user"`
}

func FormatCart(h Carts) CartResponse {
	return CartResponse{
		ID_Product: h.ID_Product,
		ID_User:    h.ID_User,
	}
}

type Cartlist struct {
	ID_Cart    int    `json:"id_cart"`
	ID_Product int    `json:"id_product"`
	Image      string `json:"image"`
	Title      string `json:"title"`
	Writer     string `json:"writer"`
	Price      int    `json:"price"`
}

func FormatCartList(h Cartlist) Cartlist {
	return Cartlist{
		ID_Cart:    h.ID_Cart,
		ID_Product: h.ID_Product,
		Image:      h.Image,
		Title:      h.Title,
		Writer:     h.Writer,
		Price:      h.Price,
	}
}

type Wishlist struct {
	ID_Wishlist int    `json:"id_wishlist"`
	ID_Product  int    `json:"id_product"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Writer      string `json:"writer"`
	Price       int    `json:"price"`
}

type WishlistResponse struct {
	ID_Product int `json:"id_product"`
	ID_User    int `json:"id_user"`
}

func FormatWishlist(h Wishlists) WishlistResponse {
	return WishlistResponse{
		ID_Product: h.ID_Product,
		ID_User:    h.ID_User,
	}
}
