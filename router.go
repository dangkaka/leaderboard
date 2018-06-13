package main

import (
	"github.com/dangkaka/leaderboard/scores"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := scores.New(db)
	r.Group("/api")
	{
		r.GET("/scores", h.Get)
		r.POST("/scores", h.Create)
		r.DELETE("/scores/:id", h.Delete)
	}

	return r
}
