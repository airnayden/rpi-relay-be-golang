package controller

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpi-relay-be-golang/model"
)

// LoginInput validation
type LoginInput struct {
	//Username	string `json:"username" binding:"required"`
	Email	string `json:"email" binding:"required,email"`
	Password	string `json:"password" binding:"required"`
}

// Login functionality
func Login(c *gin.Context) {
	// Input object
	var input LoginInput

	// Validate and show errors
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Get our User object
	user := model.User{}

	user.Email = input.Email
	user.Password = input.Password

	token, err := model.LoginCheck(user.Email, user.Password)

	// Check if record exists
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Wrong credentials!"})
		return
	}

	// OK response
	c.JSON(http.StatusOK, gin.H{"token":token})
}