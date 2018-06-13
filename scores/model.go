package scores

import (
	"time"
)

type Scores struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Scores    uint      `json:"scores"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
