package main

// Import Dependencies
import (
	"github.com/gin-gonic/gin"
	"rpi-relay-be-golang/controller"
	"rpi-relay-be-golang/middleware"
	"rpi-relay-be-golang/model"
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

	// Users CRUD
	usersGroup := protected.Group("/users")
	usersGroup.GET("/self",controller.Self)
	usersGroup.GET("/index",controller.IndexUsers)
	usersGroup.GET("/:userId",controller.ShowUser)
	usersGroup.POST("",controller.StoreUser)
	usersGroup.PUT("/:userId",controller.UpdateUser)
	usersGroup.DELETE("/:userId",controller.DestroyUser)

	// Relays CRUD
	relaysGroup := protected.Group("/relays")
	relaysGroup.GET("/index",controller.IndexRelays)
	relaysGroup.GET("/:relayId",controller.ShowRelay)
	relaysGroup.POST("",controller.StoreRelay)
	relaysGroup.PUT("/:relayId",controller.UpdateRelay)
	relaysGroup.DELETE("/:relayId",controller.DestroyRelay)

	// Relay Schedules CRUD
	relaySchedulesGroup := protected.Group("/relay-schedules")
	relaySchedulesGroup.GET("/index",controller.IndexSchedules)
	relaySchedulesGroup.GET("/:scheduleId",controller.ShowSchedule)
	relaySchedulesGroup.POST("",controller.StoreSchedule)
	relaySchedulesGroup.PUT("/:scheduleId",controller.UpdateSchedule)
	relaySchedulesGroup.DELETE("/:scheduleId",controller.DestroySchedule)

	// Relay Logs Index + Show
	relayLogsGroup := protected.Group("/relay-logs")
	relayLogsGroup.GET("/index",controller.IndexRelayLogs)
	relayLogsGroup.GET("/:scheduleId",controller.ShowRelayLog)

	// Create application server
	router.Run(":8080")
}