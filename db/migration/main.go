package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../UsedBooks.db")

	if err != nil {
		panic(err)
	}

	//role users : buyer or seller
	_, err = db.Exec(`
	
	CREATE TABLE IF NOT EXISTS users(
		id_user integer not null primary key AUTOINCREMENT,
		name varchar(255) not null,
		email varchar(255) not null,
		phone varchar(255) not null,
		password varchar(255) not null,
		role varchar(255) not null 
	);

	CREATE TABLE IF NOT EXISTS products(
		id_product integer not null primary key AUTOINCREMENT,
		id_user int not null,
		image varchar(255) not null,
		title varchar(255) not null,
		writer varchar(255) not null,
		price int not null,
		description text not null,
		FOREIGN KEY (id_user) REFERENCES users(id_user)
		
	);

	CREATE TABLE IF NOT EXISTS carts(
		id_cart integer not null primary key AUTOINCREMENT,
		id_product int not null,
		id_user int not null,
		FOREIGN KEY (id_product) REFERENCES products(id_product),
		FOREIGN KEY (id_user) REFERENCES users(id_user)
	);

	CREATE TABLE IF NOT EXISTS wishlists(
		id_wishlist integer not null primary key AUTOINCREMENT,
		id_product int not null,
		id_user int not null,
		FOREIGN KEY (id_user) REFERENCES users(id_user),
		FOREIGN KEY (id_product) REFERENCES products(id_product)
	);

	CREATE TABLE IF NOT EXISTS auth(
		id_auth integer not null primary key AUTOINCREMENT,
		id_user integer,
		token varchar(255) not null,
		expired_at datetime not null,
		FOREIGN KEY (id_user) REFERENCES users(id_user)
	);
	
	INSERT INTO users(name, email, phone, password, role)
	VALUES
	('vini', 'vini@gmail.com', '08123456', '$2a$08$gA/bRMrbE7kDuYTfMKdU4OMsarGWn1/qCXBjUhy4sV.8If0nL63I6', 'seller'),
	('jasmine', 'jasmine@gmail.com', '084567', '$2a$08$gA/bRMrbE7kDuYTfMKdU4OMsarGWn1/qCXBjUhy4sV.8If0nL63I6', 'seller'),
	('afif', 'afif@gmail.com', '08789', '$2a$08$gA/bRMrbE7kDuYTfMKdU4OMsarGWn1/qCXBjUhy4sV.8If0nL63I6', 'buyer'),
	('riski', 'riski@gmail.com', '08789', '$2a$08$gA/bRMrbE7kDuYTfMKdU4OMsarGWn1/qCXBjUhy4sV.8If0nL63I6', 'buyer'),
	('atika', 'atika@gmail.com', '08789', '$2a$08$gA/bRMrbE7kDuYTfMKdU4OMsarGWn1/qCXBjUhy4sV.8If0nL63I6', 'buyer');

	INSERT INTO products(id_product, id_user, image, title, writer, price, description)
	VALUES
	(1, 1, 'image/BI', 'Bahasa Indonesia', 'Aerlangga', 10000, 'Buku Bahasa Indonesia untuk Sekolah Dasar'),
	(2, 1, 'image/BING', 'Bahasa Inggire', 'Aerlangga', 10000, 'Buku Bahasa Inggris untuk Sekolah Dasar'),
	(3, 2, 'image/Biologi', 'Biologi', 'Aerlangga', 0, 'Buku Biologi untuk Sekolah Lanjut Tingkat Pertama');
	
	INSERT INTO carts(id_product, id_user)
	VALUES
	(1,4),
	(2,5),
	(3,3);

	INSERT INTO wishlists(id_product, id_user)
	VALUES
	(1,3),
	(2,3),
	(3,3);

	`)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
