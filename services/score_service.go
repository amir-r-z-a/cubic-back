package services

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/amir-r-z-a/cubic-back/models"
	"github.com/amir-r-z-a/cubic-back/repos"
)

type ScoreService struct {
	scoreRepo *repos.ScoreRepository
	Logger    *slog.Logger
}

func NewScoreService(scoreRepo *repos.ScoreRepository, appConf *config.AppConfig) *ScoreService {
	return &ScoreService{
		scoreRepo: scoreRepo,
		Logger:    appConf.Logger,
	}
}

func (s *ScoreService) SubmitScore(c *gin.Context) {
	var input models.SubmitScoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		s.Logger.Error("Failed to bind JSON", "error", err)
		return
	}

	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(int))

	err := s.scoreRepo.SaveScore(userID, input.GameType, input.Score)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save score"})
		s.Logger.Error("Failed to save score", "error", err, "userID", userID)
		return
	}

	c.JSON(200, gin.H{"message": "Score submitted successfully"})
	s.Logger.Info("Score submitted successfully", "userID", userID)
}

func (s *ScoreService) GetUserScores(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(int))

	scores, err := s.scoreRepo.GetUserScores(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get scores"})
		s.Logger.Error("Failed to get user scores", "error", err, "userID", userID)
		return
	}

	c.JSON(200, gin.H{"scores": scores})
	s.Logger.Info("Retrieved user scores successfully", "userID", userID, "count", len(scores))
}
