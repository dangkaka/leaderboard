package scores

import (
	"github.com/dangkaka/leaderboard/helper"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	query := h.db.Order("score desc")
	gameId := c.Query("game_id")
	if gameId != "" {
		query = query.Where("game_id = ?", gameId)
	}
	var scores []Score
	if err := query.Find(&scores).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get scores"},
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
			gin.H{"error": "Failed to decode request body"},
		)
		return
	}

	if score.Username == "" || score.Score == 0 {
		helper.ReportError("Wrong input", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Wrong input"},
		)
		return
	}

	if err := h.db.Create(&score).Error; err != nil {
		helper.ReportError("Failed to create row", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to create row"},
		)
		return
	}

	c.JSON(200, score)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty user id"})
		return
	}
	conn := h.db.Where("id = ?", id).Delete(Score{})
	if err := conn.Error; err != nil {
		helper.ReportError("Failed to delete row", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to delete row"},
		)
		return
	} else if conn.RowsAffected == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "Row not found"},
		)
		return
	}
	c.Status(http.StatusOK)
}
