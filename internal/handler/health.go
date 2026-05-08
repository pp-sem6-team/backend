package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/integration/storage"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db      *gorm.DB
	storage storage.Storage
}

func NewHealthHandler(db *gorm.DB, storage storage.Storage) *HealthHandler {
	return &HealthHandler{
		db:      db,
		storage: storage,
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
// @Failure      503  {object}  dto.HealthResponse
// @Router       /health/db [get]
func (h *HealthHandler) DBHealth(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.HealthResponse{
			Status: "db error",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.HealthResponse{
			Status: "db down",
		})
		return
	}

	c.JSON(http.StatusOK, dto.HealthResponse{
		Status: "db ok",
	})
}

// StorageHealth godoc
// @Summary      Check storage health
// @Description  Returns OK if storage is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  dto.HealthResponse
// @Failure      503  {object}  dto.HealthResponse
// @Router       /health/storage [get]
func (h *HealthHandler) StorageHealth(c *gin.Context) {
	err := h.storage.Health(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.HealthResponse{
			Status: "storage error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.HealthResponse{
		Status: "storage ok",
	})
}
