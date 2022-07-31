package controller

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpi-relay-be-golang/model"
)

// RegisterInput validation
type RegisterInput struct {
	FirstName	string `json:"first_name" binding:"required"`
	LastName	string `json:"last_name" binding:"required"`
	Username 	string `json:"username" binding:"required"`
	Password 	string `json:"password" binding:"required"`
	Email 	string `json:"email" binding:"required,email"`
}

// Register an account
func Register(c *gin.Context){
	// Our inputs
	var input RegisterInput

	// Validate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Init User model
	user := model.User{}

	// Set model properties
	user.Username = input.Username
	user.Password = input.Password
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email

	// Save Data
	_,err := user.SaveUser()

	// Check for errors
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Show success message
	c.JSON(http.StatusOK, gin.H{"message":"Your registration is successful!"})
}
