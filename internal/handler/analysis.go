package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/middleware"
	"github.com/pp-sem6-team/backend/internal/service"
)

type AnalysisHandler struct {
	analysisService *service.AnalysisService
}

func NewAnalysisHandler(analysisService *service.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{
		analysisService: analysisService,
	}
}

// CreateAnalysis godoc
// @Summary      Create analysis from photo
// @Description  Upload photo and start skin analysis
// @Tags         analyses
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        photo formData file true "Photo to analyze"
// @Success      202  {object}  dto.CreateAnalysisResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /analyses [post]
func (h *AnalysisHandler) CreateAnalysis(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
			Code:    "INVALID_REQUEST",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to open file",
			Code:    "FILE_OPEN_FAILED",
		})
		return
	}
	defer src.Close()

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	analysis, err := h.analysisService.Create(c.Request.Context(), userID, src, file.Size, file.Filename)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, analysis)
}

// ListAnalysis godoc
// @Summary      Get user analysis history
// @Description  Returns paginated list of user analyses
// @Tags         analyses
// @Produce      json
// @Security     BearerAuth
// @Param        limit query int false "Limit" default(10)
// @Param        offset query int false "Offset" default(0)
// @Success      200  {array}   dto.AnalysisListItemResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /analyses [get]
func (h *AnalysisHandler) ListAnalyses(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid offset",
			Code:    "INVALID_OFFSET",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid limit",
			Code:    "INVALID_LIMIT",
		})
		return
	}

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	analyses, err := h.analysisService.ListByUserID(
		c.Request.Context(),
		userID,
		offset,
		limit,
	)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, analyses)
}

// GetAnalysisByID godoc
// @Summary      Get analysis by ID
// @Description  Returns detailed analysis result
// @Tags         analyses
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Analysis ID"
// @Param        recs_limit query int false "Recommendations limit" default(10)
// @Param        ings_limit query int false "Ingredients limit" default(10)
// @Success      200  {object}   dto.AnalysisDetailResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      403  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router /analyses/{id} [get]
func (h *AnalysisHandler) GetAnalysisByID(c *gin.Context) {
	idStr := c.Param("id")

	recsLimitStr := c.DefaultQuery("recs_limit", "10")
	ingsLimitStr := c.DefaultQuery("ings_limit", "10")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid id",
			Code:    "INVALID_ID",
		})
		return
	}

	recsLimit, err := strconv.Atoi(recsLimitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid recs_limit",
			Code:    "INVALID_RECS_LIMIT",
		})
		return
	}

	ingsLimit, err := strconv.Atoi(ingsLimitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid ings_limit",
			Code:    "INVALID_INGS_LIMIT",
		})
		return
	}

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	analysis, err := h.analysisService.GetByID(c.Request.Context(), userID, id, recsLimit, ingsLimit)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, analysis)
}

// DeleteAnalysis godoc
// @Summary      Delete analysis by ID
// @Description  Deletes user analysis by id
// @Tags         analyses
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Analysis ID"
// @Success      204
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      403  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /analyses/{id} [delete]
func (h *AnalysisHandler) DeleteAnalysis(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid id",
			Code:    "INVALID_ID",
		})
		return
	}

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "unauthorized",
			Code:    "UNAUTHORIZED",
		})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	if err := h.analysisService.Delete(c.Request.Context(), userID, id); err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
