package scores

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/dangkaka/leaderboard/helper"
	"net/http"
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
	var scores []Score
	if err := h.db.Find(&scores).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&helper.Error{Message: "Failed to get scores"},
		)
		return
	}
	c.JSON(200, scores)
}

func (h *Handler) Create(c *gin.Context) {
	var score Score
	err := c.BindJSON(&score)

	if err != nil {
		helper.ReportError("Failed to decode request body", err)
		c.JSON(
			http.StatusBadRequest,
			helper.Error{Message: "Failed to decode request body"},
			)
		return
	}

	if err := h.db.Create(&score).Error; err != nil {
		helper.ReportError("Failed to create row", err)
		c.JSON(
			http.StatusInternalServerError,
			helper.Error{Message: "Failed to create row"},
			)
		return
	}

	c.JSON(200, score)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, &helper.Error{Message: "Empty user id"})
		return
	}
	conn := h.db.Where("id = ?", id).Delete(Score{})
	if err := conn.Error; err != nil {
		helper.ReportError("Failed to delete row", err)
		c.JSON(
			http.StatusInternalServerError,
			&helper.Error{Message: "Failed to delete row"},
		)
		return
	} else if conn.RowsAffected == 0 {
		c.JSON(
			http.StatusNotFound,
			&helper.Error{Message: "Row not found"},
		)
		return
	}
	c.Status(http.StatusOK)
}
