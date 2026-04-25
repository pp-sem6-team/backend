package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pp-sem6-team/backend/internal/dto"
	"github.com/pp-sem6-team/backend/internal/middleware"
	"github.com/pp-sem6-team/backend/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary      Register user
// @Description  Creates new user and returns JWT tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body     dto.RegisterRequest  true  "Register request"
// @Success      201  {object}  dto.AuthResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Register(
		c.Request.Context(),
		req.Email,
		req.Password,
		req.Name,
		req.BirthDate,
		req.Gender,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	})
}

// Login godoc
// @Summary      Login user
// @Description  Authenticates user and returns access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body      dto.LoginRequest true "Login request"
// @Success      200  {object}  dto.AuthResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  Invalidates refresh token and logs user out from current session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body body      dto.LogoutRequest true "Logout request"
// @Success      200  {object}  dto.MessageResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req dto.LogoutRequest

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

	err := h.authService.Logout(c.Request.Context(), userID, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "successfully logged out",
	})
}

// Refresh godoc
// @Summary      Refresh tokens
// @Description  Issues new access and refresh tokens using refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body      dto.RefreshRequest true "Refresh request"
// @Success      200  {object}  dto.AuthResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid request body: " + err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Refresh(
		c.Request.Context(),
		req.RefreshToken,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	})
}
