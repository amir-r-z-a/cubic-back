package repos

import (
	"time"

	"github.com/amir-r-z-a/cubic-back/models"
)

type ScoreRepository struct {
	Repo *AppRepo
}

func NewScoreRepository(repo *AppRepo) *ScoreRepository {
	return &ScoreRepository{Repo: repo}
}

func (r *ScoreRepository) SaveScore(userID int, gameType models.GameType, score int) error {
	query := `INSERT INTO scores (user_id, game_type, score, created_at) VALUES (?, ?, ?, ?)`
	result := r.Repo.DB.Exec(query, userID, gameType, score, time.Now())
	return result.Error
}

func (r *ScoreRepository) GetUserScores(userID int) ([]models.Score, error) {
	var scores []models.Score
	result := r.Repo.DB.Raw(`SELECT id, user_id, game_type, score, created_at FROM scores
                            WHERE user_id = ? ORDER BY created_at DESC`, userID).Scan(&scores)
	return scores, result.Error
}
