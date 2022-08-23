package repository

import (
	"database/sql"
	"fmt"
)

type WishlistRepo interface {
	InsertWishlist(inputWishlist WishlistForm) (Wishlists, error)
	FetchWishlistByIdUser(id_user int) ([]Wishlist, error)
	DeleteWishlist(id_wishlist int, id_user int) (bool, error)
}

type WishlistRepository struct {
	db *sql.DB
}

func NewWishlistRepository(db *sql.DB) *WishlistRepository {
	return &WishlistRepository{db: db}
}

func (w *WishlistRepository) DeleteWishlist(id_wishlist int, id_user int) (bool, error) {
	sqlStatement := `DELETE from wishlists WHERE id_wishlist = ? AND id_user = ?`

	_, err := w.db.Exec(sqlStatement, id_wishlist, id_user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (w *WishlistRepository) InsertWishlist(inputWishlist WishlistForm) (Wishlists, error) {
	sqlStatement := `INSERT INTO wishlists (id_product, id_user) VALUES (?,?)`

	res, err := w.db.Prepare(sqlStatement)

	if err != nil {
		return Wishlists{}, err
	}
	defer res.Close()
	newRes, err := res.Exec(
		inputWishlist.Product.ID_Product,
		inputWishlist.User.ID_User,
	)
	fmt.Println("success", newRes)

	newWishlist := Wishlists{
		ID_Product: inputWishlist.Product.ID_Product,
		ID_User:    inputWishlist.User.ID_User,
	}
	if err != nil {
		return Wishlists{}, err
	}
	return newWishlist, nil
}

func (w *WishlistRepository) FetchWishlistByIdUser(id_user int) ([]Wishlist, error) {
	var wishlists []Wishlist
	sqlStatement := `SELECT w.id_wishlist, w.id_product, p.image, p.title, p.writer, p.price FROM wishlists w INNER JOIN products p ON w.id_product = p.id_product WHERE w.id_user = ?`

	row, err := w.db.Query(sqlStatement, id_user)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var wishlist Wishlist
		err := row.Scan(
			&wishlist.ID_Wishlist,
			&wishlist.ID_Product,
			&wishlist.Image,
			&wishlist.Title,
			&wishlist.Writer,
			&wishlist.Price,
		)
		if err != nil {
			return nil, err
		}
		wishlists = append(wishlists, wishlist)
	}
	return wishlists, nil
}
