package api

import (
	"fmt"
	"net/http"
	"strconv"
	"usedbooks/helper"
	"usedbooks/repository"

	"github.com/gin-gonic/gin"
)

func (api *API) AddWishlist(c *gin.Context) {
	ID := c.Param("id")
	productID, _ := strconv.Atoi(ID)

	var wishlist repository.WishlistForm
	if err := c.ShouldBind(&wishlist); err != nil {
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
	wishlist.Product = currentProduct

	currentUser := c.MustGet("currentUser").(repository.Users)
	wishlist.User = currentUser

	addWishlist, err := api.wishlistRepository.InsertWishlist(wishlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	response := helper.APIResponse("Success to Add Wishlist", http.StatusOK, "success", repository.FormatWishlist(addWishlist))
	c.JSON(http.StatusOK, response)

}

func (api *API) GetWishlist(c *gin.Context) {
	var wishlist repository.Wishlists
	currentUser := c.MustGet("currentUser").(repository.Users)
	wishlist.User = currentUser

	res, err := api.wishlistRepository.FetchWishlistByIdUser(wishlist.User.ID_User)
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

func (api *API) DeleteWishlist(c *gin.Context) {
	wishlistId, err := strconv.Atoi(c.Param("id"))
	var wishlist repository.Wishlists
	currentUser := c.MustGet("currentUser").(repository.Users)
	wishlist.User = currentUser

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := api.wishlistRepository.DeleteWishlist(wishlistId, wishlist.User.ID_User)
	fmt.Println(success)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Wishlist Delete Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
