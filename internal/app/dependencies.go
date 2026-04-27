package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pp-sem6-team/backend/internal/config"
	"github.com/pp-sem6-team/backend/internal/handler"
	"github.com/pp-sem6-team/backend/internal/integration/ml"
	"github.com/pp-sem6-team/backend/internal/integration/storage/minio"
	"github.com/pp-sem6-team/backend/internal/middleware"
	"github.com/pp-sem6-team/backend/internal/repository/postgres"
	"github.com/pp-sem6-team/backend/internal/security/jwt"
	"github.com/pp-sem6-team/backend/internal/service"
	"gorm.io/gorm"
)

type Dependencies struct {
	HealthHandler   *handler.HealthHandler
	AuthHandler     *handler.AuthHandler
	UserHandler     *handler.UserHandler
	AnalysisHandler *handler.AnalysisHandler

	AuthMiddleware gin.HandlerFunc
}

func initDependencies(
	cfg *config.Config,
	database *gorm.DB,
	storage *minio.Client,
	mlClient ml.Client,
) *Dependencies {
	userRepo := postgres.NewUserRepository(database)
	refreshTokenRepo := postgres.NewRefreshTokenRepository(database)
	photoRepo := postgres.NewPhotoRepository(database)
	analysisRepo := postgres.NewAnalysisRepository(database)
	recommendationRepo := postgres.NewRecommendationRepository(database)
	skinTypeIngredientRepo := postgres.NewSkinTypeIngredientRepository(database)
	ingredientRepo := postgres.NewIngredientRepository(database)

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

	userService := service.NewUserService(
		userRepo,
	)

	analysisService := service.NewAnalysisService(
		photoRepo,
		analysisRepo,
		recommendationRepo,
		skinTypeIngredientRepo,
		ingredientRepo,
		mlClient,
		storage,
		cfg.Minio.PresignTTL,
	)

	return &Dependencies{
		HealthHandler:   handler.NewHealthHandler(database, storage),
		AuthHandler:     handler.NewAuthHandler(authService),
		UserHandler:     handler.NewUserHandler(userService),
		AnalysisHandler: handler.NewAnalysisHandler(analysisService),

		AuthMiddleware: middleware.AuthMiddleware(jwtManager),
	}
}
