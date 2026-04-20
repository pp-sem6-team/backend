package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/dto"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

// Health godoc
// @Summary      Check service health
// @Description  Returns OK if service is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  dto.HealthResponse
// @Router       /health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, dto.HealthResponse{
		Status: "ok",
	})
}

// DBHealth godoc
// @Summary      Check db health
// @Description  Returns OK if db is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  dto.HealthResponse
// @Failure      500  {object}  dto.HealthResponse
// @Router       /health/db [get]
func (h *HealthHandler) DBHealth(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HealthResponse{
			Status: "db error",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, dto.HealthResponse{
			Status: "db down",
		})
		return
	}

	c.JSON(http.StatusOK, dto.HealthResponse{
		Status: "db ok",
	})
}
