package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/integration/ml"
	"github.com/pp-sem6-team/backend/internal/integration/storage"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db       *gorm.DB
	storage  storage.Storage
	mlClient ml.Client
}

func NewHealthHandler(
	db *gorm.DB,
	storage storage.Storage,
	mlClient ml.Client,
) *HealthHandler {
	return &HealthHandler{
		db:       db,
		storage:  storage,
		mlClient: mlClient,
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

// MLHealth godoc
// @Summary      Check ML service health
// @Description  Returns status of ML service availability
// @Tags         health
// @Produce      json
// @Success      200  {object}  dto.HealthResponse
// @Failure      503  {object}  dto.HealthResponse
// @Router       /health/ml [get]
func (h *HealthHandler) MLHealth(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	if err := h.mlClient.Health(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.HealthResponse{
			Status: "ml error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.HealthResponse{
		Status: "ml ok",
	})
}
