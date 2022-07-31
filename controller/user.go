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

// TODO: CRUD for users
