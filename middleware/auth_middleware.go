
package middleware

import (
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/amir-r-z-a/cubic-back/services"

)

func verifyToken(tokenString string, secret []byte) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	return token.Claims, nil
}

func AuthMiddleware(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
			token = token[7:]
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			userService.Logger.Error("No token provided")
			c.Abort()
			return
		}

		claims, err := verifyToken(token, userService.SecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			userService.Logger.Error("Invalid token")
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
