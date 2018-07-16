package scores

import (
	"time"
)

type Score struct {
	Id        int       `json:"id" gorm:"primary_key"`
	GameId    int       `json:"game_id" gorm:"index:game_id"`
	Username  string    `json:"username"`
	Score     int       `json:"score" gorm:"index:score"`
	CreatedAt time.Time `json:"created_at" gorm:"index:created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Game struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
