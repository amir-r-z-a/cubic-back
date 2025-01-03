package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameType  GameType  `json:"game_type"`
	Rating    int       `json:"rating"`
	Feedback  string    `json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
}

type SubmitCommentInput struct {
	GameType GameType `json:"game_type" binding:"required,min=1,max=3"`
	Rating   int      `json:"rating" binding:"required,min=1,max=5"`
	Feedback string   `json:"feedback" binding:"max=1000"`
}
