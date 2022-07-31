package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rpi-relay-be-golang/util/token"
)

// JwtAuthMiddleware - check our JWT token if present in the request
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized!")
			c.Abort()
			return
		}
		c.Next()
	}
}
