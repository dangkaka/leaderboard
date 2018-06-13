package main

import (
	"github.com/dangkaka/leaderboard/scores"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Group("/api")
	{
		r.GET("/scores", scores.Get)
		r.POST("/scores", scores.Create)
		r.DELETE("/scores/:id", scores.Delete)
	}

	return r
}
