package repository

type Users struct {
	ID_User  int
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string
}

type Products struct {
	ID_Product  int
	ID_User     int
	Image       string
	Title       string
	Writer      string
	Price       int
	Description string
	User        Users //relasi
}

type Carts struct {
	ID_Cart    int
	ID_Product int
	ID_User    int
	User       Users    //relasi
	Product    Products //relasi
}

type Wishlists struct {
	ID_Wishlist int
	ID_Product  int
	ID_User     int
	User        Users    //relasi
	Product     Products //relasi
}
