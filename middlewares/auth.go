package middlewares

import (
	"net/http"
	"strings"

	"go-auth/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"succeed": false, "message": "Authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"succeed": false, "message": "Invalid authorization format"})
			return
		}

		userId, err := jwtService.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"succeed": false, "message": err.Error()})
			return
		}

		// Store userId in context for handlers to use
		c.Set("userId", userId)

		c.Next()
	}
}
