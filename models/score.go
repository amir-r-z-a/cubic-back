package models

import "time"

type Score struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameType  GameType  `json:"game_type"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

type SubmitScoreInput struct {
	GameType GameType `json:"game_type" binding:"required,min=1,max=3"`
	Score    int      `json:"score" binding:"min=0"`
}
