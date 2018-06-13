package scores

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Get(c *gin.Context) {
	var scores []Scores
	if err := h.db.Find(&scores).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, scores)
	}

}

func (h *Handler) Create(c *gin.Context) {
	var scores Scores
	c.BindJSON(&scores)
	h.db.Create(&scores)
	c.JSON(200, scores)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var scores Scores
	d := h.db.Where("id = ?", id).Delete(&scores)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
