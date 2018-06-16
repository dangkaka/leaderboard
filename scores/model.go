package scores

import (
	"time"
)

type Score struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Score     uint      `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
