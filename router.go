package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Group("/api")
	{
		r.GET("/scores", func(c *gin.Context) {
			c.String(http.StatusOK, "Get scores")
		})
		r.POST("/scores", func(c *gin.Context) {
			c.String(http.StatusOK, "Post scores")
		})
		r.DELETE("/scores/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.String(http.StatusOK, "Delete %s", id)
		})
	}

	return r
}
