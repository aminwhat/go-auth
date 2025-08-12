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

		// Support both "Bearer <token>" and "<token>"
		var token string
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			token = parts[1]
		} else {
			token = authHeader
		}

		userId, err := jwtService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"succeed": false, "message": err.Error()})
			return
		}

		c.Set("userId", userId)
		c.Next()

	}
}
