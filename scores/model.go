package scores

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Scores struct {
	ID        uint   `gorm:"primary_key"`
	Username  string
	Scores    uint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	DB *gorm.DB
}
