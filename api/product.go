package api

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"usedbooks/helper"
	"usedbooks/repository"

	"github.com/gin-gonic/gin"
)

func (api *API) GetProducts(c *gin.Context) {
	products, err := api.productRepository.FetchProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var productsResponse []repository.ProductResponse

	for _, product := range products {
		productResponse := repository.FormatProduct(product)

		productsResponse = append(productsResponse, productResponse)
	}

	response := helper.APIResponse("Success Get List products", http.StatusOK, "success", productsResponse)
	c.JSON(http.StatusOK, response)
}

func (api *API) GetProductByID(c *gin.Context) {
	ID := c.Param("id")
	product_id, _ := strconv.Atoi(ID)

	res, err := api.productRepository.FindProductByID(product_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	formatter := repository.FormatProduct(res) //untuk menampilkan json response sesuai format yang diinginkan

	response := helper.APIResponse("Success Get Detail product", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (api *API) AddNewProduct(c *gin.Context) {
	var product repository.ProductFrom
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	currentUser := c.MustGet("currentUser").(repository.Users)
	product.User = currentUser

	product.Image = UploadFileProduct(c)

	if product.Image == "" {
		c.JSON(http.StatusBadRequest, "error: Input is Empty")
		return
	}

	NewProduct, err := api.productRepository.InsertProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	response := helper.APIResponse("Success to Add product", http.StatusOK, "success", repository.FormatProduct(NewProduct))
	c.JSON(http.StatusOK, response)
}

func (api *API) UpdateProduct(c *gin.Context) {
	var product repository.ProductFromUpdate

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	oldImg, err := api.productRepository.FetchNameImgProductId(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Image = UploadFileProduct(c)

	if product.Image != "" {
		if err := api.DeleteFileProduct(c, productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		product.Image = *oldImg
	}
	currentUser := c.MustGet("currentUser").(repository.Users)
	product.User = currentUser
	product.ID_Product = productId

	NewProduct, err := api.productRepository.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	response := helper.APIResponse("Success to Update product", http.StatusOK, "success", repository.FormatProduct(NewProduct))
	c.JSON(http.StatusOK, response)
}

func (api *API) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	if err := api.DeleteFileProduct(c, productId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := api.productRepository.DeleteProduct(productId)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Product Delete Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

}

///FILE/////

func UploadFileProduct(c *gin.Context) (fileName string) {
	c.Request.ParseMultipartForm(10 << 20)

	file, err := c.FormFile("image")

	if file != nil && err == nil {
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer src.Close()

		currentUser := c.MustGet("currentUser").(repository.Users)
		userID := strconv.Itoa(currentUser.ID_User)
		namefile := userID + "-" + file.Filename
		filepath := "image/" + namefile

		dst, err := os.Create(filepath)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}

		return namefile
	}
	return ""
}

func (api *API) DeleteFileProduct(c *gin.Context, id int) error {
	oldFile, err := api.productRepository.FetchNameImgProductId(id)

	filepath := "image/" + *oldFile

	if *oldFile != "" {
		if err := os.Remove(filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		return err
	}
	return nil
}
