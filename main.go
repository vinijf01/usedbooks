package main

import (
	"database/sql"
	"usedbooks/api"
	"usedbooks/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/UsedBooks.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	wishlistRepo := repository.NewWishlistRepository(db)

	mainAPI := api.NewAPI(*usersRepo, *productRepo, *cartRepo, *wishlistRepo)

	mainAPI.Start()
}
