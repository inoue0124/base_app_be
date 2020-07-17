package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inoue0124/salon-be/forms"
	"github.com/inoue0124/salon-be/services"
)

type UserController struct{}

func (u UserController) RetrieveUser(c *gin.Context) {
	if c.Param("id") != "" {
		userService := new(services.UserService)
		user, err := userService.GetByID(c.Param("id"))
		if err != nil {
			println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

func (u UserController) SignupUser(c *gin.Context) {
	userPayload := forms.UserSignup{}
	err := c.ShouldBindJSON(&userPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to signup user", "error": err.Error()})
		c.Abort()
		return
	}
	userService := new(services.UserService)
	user, err := userService.Signup(userPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to signup user", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered!", "user": user})
	return
}

func (u UserController) SigninUser(c *gin.Context) {
	return
}

func (u UserController) SignoutUser(c *gin.Context) {
	return
}
