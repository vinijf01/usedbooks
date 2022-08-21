package api

import (
	"net/http"
	"strconv"
	"time"
	"usedbooks/helper"
	"usedbooks/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (api *API) RegisterUser(c *gin.Context) {
	var user repository.RegisterUserInput

	err := c.ShouldBindJSON(&user)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors} //gin.H adalah map key string value int

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	res, err := api.userRepository.InsertUser(user.Name, user.Email, user.Phone, string(hashedPassword), user.Role)
	if res == nil && err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Email Has Registered!",
		})
		return
	}

	//response json
	newUser, err := api.userRepository.FetchUserByEmail(user.Email)

	formatter := repository.FormatUserRegister(newUser) //untuk menampilkan json response sesuai format yang diinginkan
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (api *API) Login(c *gin.Context) {
	var input repository.UserLogin

	//process binding
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := api.userRepository.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)

	//generete token
	token, err := NewService().GenerateToken(loggedinUser.ID_User, loggedinUser.Email, loggedinUser.Role, expirationTime.Unix())
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//pushToken
	var pushtoken Token

	pushtoken = Token{
		UserId:    loggedinUser.ID_User,
		Token:     token,
		ExpiresAt: expirationTime,
	}

	tknToDb, err := api.userRepository.PushToken(pushtoken.UserId, pushtoken.Token, pushtoken.ExpiresAt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
		Path:    ("/"),
	})

	//untuk menampilkan json response sesuai format yang diinginkan
	response := helper.APIResponse("Successfuly Loggedin", http.StatusOK, "success", repository.LoginResponse{
		Email:     *&loggedinUser.Email,
		Token:     *tknToDb,
		ExpiresAt: expirationTime,
	})
	c.JSON(http.StatusOK, response)

}

func (api *API) LogoutUser(c *gin.Context) {
	token, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no cookie"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if token.Value == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no cookie"})
		return
	}

	api.userRepository.DeleteToken(token.Value)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
		Path:    ("/"),
	})

	response := helper.APIResponse("Logout successful", http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)

}

func (api *API) GetUserByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	result, err := api.userRepository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
