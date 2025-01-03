package middleware

import (
	"net/http"
	"strings"

	"github.com/amir-r-z-a/cubic-back/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

		// Extract username from claims
		mapClaims := claims.(jwt.MapClaims)
		username := mapClaims["username"].(string)

		// Get user from database
		user, err := userService.Repo.GetUser(username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			userService.Logger.Error("User not found", "error", err)
			c.Abort()
			return
		}

		// Add user ID to claims
		mapClaims["user_id"] = user.ID
		c.Set("claims", mapClaims)

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
