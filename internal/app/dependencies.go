package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/config"
	"github.com/pp-sem6-team/backend/internal/handlers"
	"github.com/pp-sem6-team/backend/internal/integration/storage/minio"
	"github.com/pp-sem6-team/backend/internal/middleware"
	"github.com/pp-sem6-team/backend/internal/repository/postgres"
	"github.com/pp-sem6-team/backend/internal/security/jwt"
	"github.com/pp-sem6-team/backend/internal/service"
	"gorm.io/gorm"
)

type Dependencies struct {
	HealthHandler *handlers.HealthHandler
	AuthHandler   *handlers.AuthHandler

	AuthMiddleware gin.HandlerFunc
}

func initDependencies(
	cfg *config.Config,
	database *gorm.DB,
	storage *minio.Client,
) *Dependencies {
	userRepo := postgres.NewUserRepository(database)
	refreshTokenRepo := postgres.NewRefreshTokenRepository(database)

	jwtManager := jwt.NewManager(
		cfg.JWT.Secret,
		cfg.JWT.AccessTokenTTL,
		cfg.JWT.RefreshTokenTTL,
	)

	authService := service.NewAuthService(
		userRepo,
		refreshTokenRepo,
		jwtManager,
	)

	return &Dependencies{
		HealthHandler: handlers.NewHealthHandler(database, storage),
		AuthHandler:   handlers.NewAuthHandler(authService),

		AuthMiddleware: middleware.AuthMiddleware(jwtManager),
	}
}
