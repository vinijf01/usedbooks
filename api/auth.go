package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//reponse token
type Token struct {
	UserId    int       `json:"id_user" binding:"required"`
	Token     string    `json:"token" binding:"required"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

var jwtKey = []byte("secret") //key rahasi yang digunakan untuk signature key

type Claims struct {
	Email   string
	Role    string
	Id_User int
	jwt.StandardClaims
}

////////////////

type AuthService interface {
	GenerateToken(ID_User int, Email string, Role string) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ID_User int, Email string, Role string, Exp int64) (string, error) {
	//data dalam payload / claim
	claims := jwt.MapClaims{}
	claims["id_user"] = ID_User
	claims["email"] = Email
	claims["role"] = Role
	claims["exp"] = Exp

	//generate token / membuat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//signature key
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
