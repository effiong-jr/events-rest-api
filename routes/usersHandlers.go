package routes

import (
	"net/http"

	"example.com/events-rest-api/models"
	"example.com/events-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signupUserHandler(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Somthing went wrong", "error": err.Error()})
		return
	}

	u, err := user.RegisterUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registration successful", "user": u})

}

func loginHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "An error occured.", "error": err.Error()})
		return
	}

	u, err := user.LoginUser()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email or Password incorrect"})
		return
	}

	token, err := utils.GenerateJWT(u.Email, u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
