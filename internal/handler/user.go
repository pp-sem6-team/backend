package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/middleware"
	"github.com/pp-sem6-team/backend/internal/service"
	"github.com/pp-sem6-team/backend/internal/service/mapper"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetMe godoc
// @Summary      Get current user
// @Description  Returns information about the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.UserResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/me [get]
func (h *UserHandler) GetMe(c *gin.Context) {
	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "unauthorized"})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	user, err := h.userService.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mapper.ToResponse(user))
}

// UpdateMe godoc
// @Summary      Update current user
// @Description  Updates current user and returns updated user information
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body body      dto.UserRequest true "User Request"
// @Success      200  {object}  dto.UserResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/me [patch]
func (h *UserHandler) UpdateMe(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "unauthorized"})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	user, err := h.userService.Update(c.Request.Context(), userID, req.Email, req.Name, req.BirthDate, req.Gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mapper.ToResponse(user))
}

// UpdateMyPassword godoc
// @Summary      Update password
// @Description  Changes current user password after verifying current password
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body body      dto.UserPasswordRequest true "User Password Request"
// @Success      204
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/me/password [patch]
func (h *UserHandler) UpdateMyPassword(c *gin.Context) {
	var req dto.UserPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}

	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "unauthorized"})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	err := h.userService.UpdatePassword(c.Request.Context(), userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteMe godoc
// @Summary      Delete current user
// @Description  Deletes authenticated user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      204
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/me [delete]
func (h *UserHandler) DeleteMe(c *gin.Context) {
	userIDRaw, exists := c.Get(string(middleware.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "unauthorized"})
		return
	}

	userID := userIDRaw.(uuid.UUID)

	err := h.userService.Delete(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
