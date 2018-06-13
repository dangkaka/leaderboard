package scores

import (
	"github.com/gin-gonic/gin"
)

func (m *Model) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

func Get(m *Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"get": "get",
		})
	}
}


func Create(m *Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "test",
		})
	}
}

func Delete(m *Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "test",
		})
	}
}

