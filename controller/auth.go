package controller

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpi-relay-be-golang/model"
	"rpi-relay-be-golang/util/token"
)

// Self - display the logged-in User data
func Self(c *gin.Context){
	userId, err := token.ExtractTokenId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Get the user by his ID
	user,err := model.GetUserById(userId)

	// Show Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Show OK message
	c.JSON(http.StatusOK, gin.H{"message":"Success","data":user})
}

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