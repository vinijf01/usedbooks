package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"usedbooks/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyCustomClaims struct {
	Id_User int `json:"id_user"`
	jwt.StandardClaims
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//menentukan domain apa saja yg boleh mengakses data/aplikiasi ini
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//method apa saja yang boleh diakses
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "application/json, multipart/form-data")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}

func (api *API) AuthMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		tokenString := token.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "email", claims.Email)
		ctx = context.WithValue(c.Request.Context(), "id_user", claims.Id_User)
		ctx = context.WithValue(ctx, "role", claims.Role)
		ctx = context.WithValue(ctx, "props", claims)
		c.Request = c.Request.WithContext(ctx)

		next(c)
	}
}

func (api *API) SellerMiddlerware(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		role := c.Request.Context().Value("role")
		if role != "seller" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden access"})
			c.Abort()
			return
		}

		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization Header Not Found"})
			return
		}

		splitToken := strings.Split(auth, "Bearer ")
		auth = splitToken[1]

		token, _ := jwt.ParseWithClaims(auth, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AccessToken"), nil
		})

		claim, _ := token.Claims.(*MyCustomClaims)
		user, err := api.userRepository.GetUserByID(claim.Id_User)
		if err != nil {
			response := helper.APIResponse("Unauthorized Get use by id", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//set contex
		c.Set("currentUser", user)

		next(c)
	})
}

func (api *API) BuyerMiddlerware(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		role := c.Request.Context().Value("role")
		if role != "buyer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden access"})
			c.Abort()
			return
		}
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization Header Not Found"})
			return
		}

		splitToken := strings.Split(auth, "Bearer ")
		auth = splitToken[1]

		token, _ := jwt.ParseWithClaims(auth, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AccessToken"), nil
		})

		claim, _ := token.Claims.(*MyCustomClaims)
		user, err := api.userRepository.GetUserByID(claim.Id_User)
		if err != nil {
			response := helper.APIResponse("Unauthorized Get use by id", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)

		next(c)
	})
}

func (api *API) GET(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need GET Method!"})
			return
		}
		next(c)
	})
}

func (api *API) POST(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need POST Method!"})
			return
		}
		next(c)
	})
}

func (api *API) DELETE(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Method != http.MethodDelete {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need DELETE Method!"})
			return
		}
		next(c)
	})
}

func (api *API) PATCH(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Method != http.MethodPatch {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"Error": "Need PATCH Method!"})
			return
		}
		next(c)
	})
}
