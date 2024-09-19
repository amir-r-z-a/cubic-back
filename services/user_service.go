package services

import (
	"fmt"
	"log/slog"

	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/amir-r-z-a/cubic-back/repos"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	Repo   *repos.UserRepo
	Logger *slog.Logger
}

func NewUserService(userRepo *repos.UserRepo, appConf *config.AppConfig) *UserService {
	return &UserService{Repo: userRepo, Logger: appConf.Logger}
}

func (us UserService) SignUp(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		us.Logger.Error("Failed to bind request body")
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	fmt.Println("Request Body:", requestBody)
	
	c.JSON(200, gin.H{"message": "Request received", "data": requestBody})
	
}
