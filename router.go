package main

import (
	"github.com/dangkaka/leaderboard/scores"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	scoresModel := &scores.Model{DB: db}
	r.Group("/api")
	{
		r.GET("/scores", scores.Get(scoresModel))
		r.POST("/scores", scores.Create(scoresModel))
		r.DELETE("/scores/:id", scores.Delete(scoresModel))
	}

	return r
}
