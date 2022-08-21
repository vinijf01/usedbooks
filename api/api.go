package api

import (
	"fmt"
	"os"
	"usedbooks/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type API struct {
	userRepository    repository.UserRepository
	productRepository repository.ProductRepository
	cartRepository    repository.CartRepository
	gin               *gin.Engine
}

func NewAPI(userRepository repository.UserRepository, productRepository repository.ProductRepository, cartRepository repository.CartRepository) API {
	gin := gin.Default()

	gin.Use(CORSMiddleware())

	api := API{
		userRepository, productRepository, cartRepository, gin,
	}

	v1 := gin.Group("/api/v1")
	gin.Use(cors.Default())

	//USER
	v1.POST("/user/register", api.POST(api.RegisterUser))
	v1.POST("/user/login", api.POST(api.Login))
	v1.POST("/user/logout", api.POST(api.LogoutUser))
	v1.GET("/user/:id", api.GET(api.GetUserByID))

	//Product
	//list product
	v1.GET("/product/list", api.GET(api.GetProducts))
	v1.GET("/product/detail/:id", api.AuthMiddleware(api.GetProductByID))
	v1.POST("/product/insert", api.POST(api.AuthMiddleware(api.SellerMiddlerware(api.AddNewProduct))))
	v1.PATCH("/product/update/:id", api.PATCH(api.AuthMiddleware(api.SellerMiddlerware(api.UpdateProduct))))
	v1.DELETE("/product/delete/:id", api.DELETE(api.AuthMiddleware(api.SellerMiddlerware(api.DeleteProduct))))

	//Cart
	//list cart
	v1.POST("/cart/add/:id", api.POST(api.AuthMiddleware(api.BuyerMiddlerware(api.AddCart))))
	v1.GET("/cart/list", api.GET(api.AuthMiddleware(api.BuyerMiddlerware(api.GetCarts))))
	v1.DELETE("/cart/delete/:id", api.DELETE(api.AuthMiddleware(api.BuyerMiddlerware(api.DeleteCarts))))

	//wishlist

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	api.Handler().Run(":" + port)
}
