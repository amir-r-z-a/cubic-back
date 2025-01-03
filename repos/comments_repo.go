package repos

import (
	"time"

	"github.com/amir-r-z-a/cubic-back/models"
)

type CommentRepository struct {
	Repo *AppRepo
}

func NewCommentRepository(repo *AppRepo) *CommentRepository {
	return &CommentRepository{Repo: repo}
}

func (r *CommentRepository) SaveComment(userID int, gameType models.GameType, rating int, feedback string) error {
	query := `INSERT INTO comments (user_id, game_type, rating, feedback, created_at) VALUES (?, ?, ?, ?, ?)`
	result := r.Repo.DB.Exec(query, userID, gameType, rating, feedback, time.Now())
	return result.Error
}

func (r *CommentRepository) GetGameComments(gameType models.GameType) ([]models.Comment, error) {
	var comments []models.Comment
	result := r.Repo.DB.Raw(`SELECT id, user_id, game_type, rating, feedback, created_at FROM comments 
	                         WHERE game_type = ? ORDER BY created_at DESC`, gameType).Scan(&comments)
	return comments, result.Error
}