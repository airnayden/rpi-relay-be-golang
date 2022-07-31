package controller

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	// OK Response
	c.JSON(http.StatusOK, gin.H{"message": "Input data is valid!"})
}