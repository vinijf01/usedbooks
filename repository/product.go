package repository

import (
	"database/sql"
	"fmt"
)

type ProductRepo interface {
	FetchProducts() ([]Products, error)
	FindProductByID(ID int) (Products, error)
	InsertProduct(inputProduct ProductFrom) (Products, error)
	UpdateProduct(product ProductFromUpdate) (Products, error)
	DeleteProduct(id int) (bool, error)
	FetchNameImgProductId(id int) (*string, error)
	// UpdateProduct(id_product int, id_user int, image string, title string, writer string, price int, description string) (bool, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProducts() ([]Products, error) {
	var products []Products

	row, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}
	for row.Next() {
		var product Products

		err := row.Scan(
			&product.ID_Product,
			&product.ID_User,
			&product.Image,
			&product.Title,
			&product.Writer,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) FindProductByID(ID int) (Products, error) {
	product := Products{}
	sqlStatement := `SELECT * From products WHERE id_product = ?`

	row := p.db.QueryRow(sqlStatement, ID)
	err := row.Scan(
		&product.ID_Product,
		&product.ID_User,
		&product.Image,
		&product.Title,
		&product.Writer,
		&product.Price,
		&product.Description,
	)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) InsertProduct(inputProduct ProductFrom) (Products, error) {
	sqlStatement := `INSERT INTO products (id_user, image, title, writer, price, description) VALUES (?,?,?,?,?,?)`

	res, err := p.db.Prepare(sqlStatement)

	if err != nil {
		return Products{}, err
	}
	defer res.Close()
	newRes, err := res.Exec(
		inputProduct.User.ID_User,
		inputProduct.Image,
		inputProduct.Title,
		inputProduct.Writer,
		inputProduct.Price,
		inputProduct.Description,
	)
	fmt.Println("success", newRes)

	newProduct := Products{
		ID_User:     inputProduct.User.ID_User,
		Image:       inputProduct.Image,
		Title:       inputProduct.Title,
		Writer:      inputProduct.Writer,
		Price:       inputProduct.Price,
		Description: inputProduct.Description,
	}
	if err != nil {
		return Products{}, err
	}
	return newProduct, nil
}

// func (c *ProductRepository) UpdateProduct(id_product int, id_user int, image string, title string, writer string, price int, description string) (bool, error) {
// 	sqlStatement := `UPDATE products SET id_user = ?, image = ?, title = ?, writer = ?, price = ?, description = ? WHERE id_product = ?`

// 	_, err := c.db.Exec(sqlStatement, id_user, image, title, writer, price, description, id_product)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

func (p *ProductRepository) UpdateProduct(product ProductFromUpdate) (Products, error) {
	sqlStatement := `UPDATE products SET id_user = ?, image = ?, title = ?, writer = ?, price = ?, description = ? WHERE id_product = ?`

	res, err := p.db.Prepare(sqlStatement)

	if err != nil {
		return Products{}, err
	}
	defer res.Close()
	newRes, err := res.Exec(
		product.User.ID_User,
		product.Image,
		product.Title,
		product.Writer,
		product.Price,
		product.Description,
		product.ID_Product,
	)
	fmt.Println("success", newRes)

	newProduct := Products{
		ID_User:     product.User.ID_User,
		Image:       product.Image,
		Title:       product.Title,
		Writer:      product.Writer,
		Price:       product.Price,
		Description: product.Description,
		ID_Product:  product.ID_Product,
	}
	if err != nil {
		return Products{}, err
	}
	return newProduct, nil
}

func (p *ProductRepository) DeleteProduct(id int) (bool, error) {
	sqlStatement := `DELETE from products WHERE id_product = ?`

	_, err := p.db.Exec(sqlStatement, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *ProductRepository) FetchNameImgProductId(id int) (*string, error) {
	product := Products{}
	sqlStatement := `SELECT image FROM products WHERE id_product = ?`

	row := p.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&product.Image,
	)
	if err != nil {
		return nil, err
	}
	return &product.Image, nil
}
