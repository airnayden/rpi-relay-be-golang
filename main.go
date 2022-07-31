package main

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"rpi-relay-be-golang/controller"
	"rpi-relay-be-golang/model"
	"rpi-relay-be-golang/middleware"
)

// Entry Method
func main() {
	// Init DB Connection
	model.ConnectToDb()

	// Init Router
	router := gin.Default()

	// Create `api` route group
	public := router.Group("/api")

	// Register Endpoint
	public.POST("/register", controller.Register)

	// Login Endpoint
	public.POST("/login",controller.Login)

	// Create e new route group for authenticated User
	protected := router.Group("/api/admin")

	// Attach Middleware before executing the request
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/self",controller.Self)

	// Create application server
	router.Run(":8080")
}