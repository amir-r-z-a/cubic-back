package services

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/amir-r-z-a/cubic-back/models"
	"github.com/amir-r-z-a/cubic-back/repos"
)

type CommentService struct {
	commentRepo *repos.CommentRepository
	Logger      *slog.Logger
}

func NewCommentService(commentRepo *repos.CommentRepository, appConf *config.AppConfig) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		Logger:      appConf.Logger,
	}
}

func (s *CommentService) SubmitComment(c *gin.Context) {
	var input models.SubmitCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		s.Logger.Error("Failed to bind JSON", "error", err)
		return
	}

	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(int))

	err := s.commentRepo.SaveComment(userID, input.GameType, input.Rating, input.Feedback)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save comment"})
		s.Logger.Error("Failed to save comment", "error", err, "userID", userID)
		return
	}

	c.JSON(200, gin.H{"message": "Comment submitted successfully"})
	s.Logger.Info("Comment submitted successfully", "userID", userID)
}

func (s *CommentService) GetGameComments(c *gin.Context) {
	var input struct {
		GameType models.GameType `form:"game_type" binding:"required,min=1,max=3"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		s.Logger.Error("Failed to bind query parameters", "error", err)
		return
	}

	comments, err := s.commentRepo.GetGameComments(input.GameType)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get comments"})
		s.Logger.Error("Failed to get game comments", "error", err, "gameType", input.GameType)
		return
	}

	c.JSON(200, gin.H{"comments": comments})
	s.Logger.Info("Retrieved game comments successfully", "gameType", input.GameType, "count", len(comments))
}
