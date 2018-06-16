package main

import (
	"github.com/dangkaka/leaderboard/scores"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/health", HealthCheckHandler)
	h := scores.New(db)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/scores", h.Get)
		v1.POST("/scores", h.Create)
		v1.DELETE("/scores/:id", h.Delete)
	}

	return r
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
