package router

import (
	"github.com/amir-r-z-a/cubic-back/middleware"
	"github.com/amir-r-z-a/cubic-back/services"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.Engine, commentService *services.CommentService, userService *services.UserService) {
	comments := r.Group("/api/comments")
	comments.Use(middleware.AuthMiddleware(userService))
	{
		comments.POST("/submit", commentService.SubmitComment)
		comments.GET("/game-comments", commentService.GetGameComments)
	}
}