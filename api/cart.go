package api

import (
	"fmt"
	"net/http"
	"strconv"
	"usedbooks/helper"
	"usedbooks/repository"

	"github.com/gin-gonic/gin"
)

func (api *API) AddCart(c *gin.Context) {
	ID := c.Param("id")
	productID, _ := strconv.Atoi(ID)

	var cart repository.CartForm
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	currentProduct, err := api.productRepository.FindProductByID(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	cart.Product = currentProduct

	currentUser := c.MustGet("currentUser").(repository.Users)
	cart.User = currentUser

	addCart, err := api.cartRepository.InsertCart(cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	response := helper.APIResponse("Success to Add Cart", http.StatusOK, "success", repository.FormatCart(addCart))
	c.JSON(http.StatusOK, response)

}

func (api *API) GetCarts(c *gin.Context) {
	var cart repository.Carts
	currentUser := c.MustGet("currentUser").(repository.Users)
	cart.User = currentUser

	res, err := api.cartRepository.FetchCartByIdUser(cart.User.ID_User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	if res == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "No activity",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (api *API) DeleteCarts(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("id"))
	var cart repository.Carts
	currentUser := c.MustGet("currentUser").(repository.Users)
	cart.User = currentUser

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := api.cartRepository.DeleteCart(cartId, cart.User.ID_User)
	fmt.Println(success)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Cart Delete Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
