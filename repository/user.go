package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	FetcUsers() ([]Users, error)
	InsertUser(name string, email string, phone string, password string, role string) (*string, error)
	LoginUser(input UserLogin) (Users, error)
	FetchUserByEmail(email string) (Users, error)
	PushToken(id_user int, token string, expired_at time.Time) (*string, error)
	DeleteToken(token string) (bool, error)
	GetUserByID(Id_user int) (Users, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUsers() ([]Users, error) {
	var users []Users

	rows, err := u.db.Query("SELECT * FROM users")

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Users

		err := rows.Scan(
			&user.ID_User,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
		)

		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) InsertUser(name string, email string, phone string, password string, role string) (*string, error) {
	users, err := u.FetchUsers()

	for _, value := range users {
		if value.Email == email {
			return nil, err
		}
	}

	_, err = u.db.Exec("INSERT INTO users (name, email, phone, password, role) VALUES (?,?,?,?,?)", name, email, phone, password, role)

	if err != nil {
		return nil, err
	}
	return &email, err
}

func (u *UserRepository) LoginUser(input UserLogin) (Users, error) {
	email := input.Email
	password := input.Password

	//mencari user dg email yg diinputkan
	user, err := u.FetchUserByEmail(email)
	if err != nil {
		return user, errors.New("login failed111")
	}

	if user.ID_User == 0 {
		fmt.Println(user)
		return user, errors.New("No User found on that email")
	}

	//mencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) FetchUserByEmail(email string) (Users, error) {
	user := Users{}
	sqlStatement := `SELECT * FROM users WHERE email = ?`

	row := u.db.QueryRow(sqlStatement, email)
	err := row.Scan(
		&user.ID_User,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) PushToken(id_user int, token string, expired_at time.Time) (*string, error) {
	_, err := u.db.Exec("INSERT INTO auth (id_user, token, expired_at) VALUES (?, ?,?)",
		id_user, token, expired_at)

	if err != nil {
		return nil, err
	}
	return &token, err
}

func (u *UserRepository) DeleteToken(token string) (bool, error) {
	_, err := u.db.Exec("DELETE FROM auth WHERE token = ?", token)

	if err != nil {
		return false, err
	}
	return true, err
}

func (u *UserRepository) GetUserByID(Id_user int) (Users, error) {
	user := Users{}
	sqlStatement := `SELECT * FROM users WHERE id_user = ?`

	row := u.db.QueryRow(sqlStatement, Id_user)
	err := row.Scan(
		&user.ID_User,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
