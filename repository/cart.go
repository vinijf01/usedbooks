package repository

import (
	"database/sql"
	"fmt"
)

type CartRepo interface {
	InsertCart(inputCart CartForm) (Carts, error)
	FetchCartByIdUser(id_user int) ([]Cartlist, error)
	DeleteCart(id_cart int, id_user int) (bool, error)
}

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) InsertCart(inputCart CartForm) (Carts, error) {
	sqlStatement := `INSERT INTO carts (id_product, id_user) VALUES (?,?)`

	res, err := c.db.Prepare(sqlStatement)

	if err != nil {
		return Carts{}, err
	}
	defer res.Close()
	newRes, err := res.Exec(
		inputCart.Product.ID_Product,
		inputCart.User.ID_User,
	)
	fmt.Println("success", newRes)

	newCart := Carts{
		ID_Product: inputCart.Product.ID_Product,
		ID_User:    inputCart.User.ID_User,
	}
	if err != nil {
		return Carts{}, err
	}
	return newCart, nil
}

func (c *CartRepository) FetchCartByIdUser(id_user int) ([]Cartlist, error) {
	var carts []Cartlist
	sqlStatement := `SELECT c.id_cart, c.id_product, p.image, p.title, p.writer, p.price FROM carts c INNER JOIN products p ON c.id_product = p.id_product WHERE c.id_user = ?`

	row, err := c.db.Query(sqlStatement, id_user)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var cart Cartlist
		err := row.Scan(
			&cart.ID_Cart,
			&cart.ID_Product,
			&cart.Image,
			&cart.Title,
			&cart.Writer,
			&cart.Price,
		)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (c *CartRepository) DeleteCart(id_cart int, id_user int) (bool, error) {
	sqlStatement := `DELETE from carts WHERE id_cart = ? AND id_user = ?`

	_, err := c.db.Exec(sqlStatement, id_cart, id_user)
	if err != nil {
		return false, err
	}
	return true, nil
}
