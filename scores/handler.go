package scores

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

