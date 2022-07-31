package main

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"rpi-relay-be-golang/controller"
	"rpi-relay-be-golang/model"
)

// Entry Method
func main() {
	// Init DB Connection
	model.ConnectToDb()

	// Init Router
	r := gin.Default()

	// Create `api` route group
	public := r.Group("/api")

	// Register Endpoint
	public.POST("/register", controller.Register)

	// Create application server
	r.Run(":8080")
}