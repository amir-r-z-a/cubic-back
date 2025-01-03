package router

import (
	"github.com/amir-r-z-a/cubic-back/middleware"
	"github.com/amir-r-z-a/cubic-back/services"
	"github.com/gin-gonic/gin"
)

func SetupScoreRoutes(r *gin.Engine, scoreService *services.ScoreService, userService *services.UserService) {
	scores := r.Group("/api/scores")
	scores.Use(middleware.AuthMiddleware(userService))
	{
		scores.POST("/submit", scoreService.SubmitScore)
		scores.GET("/my-scores", scoreService.GetUserScores)
	}
}
