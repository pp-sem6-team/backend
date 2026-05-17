package app

import (
	"github.com/gin-gonic/gin"

	_ "github.com/pp-sem6-team/backend/docs"
	"github.com/pp-sem6-team/backend/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(deps *Dependencies) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	r.GET("/health", deps.HealthHandler.Health)
	r.GET("/health/db", deps.HealthHandler.DBHealth)
	r.GET("/health/storage", deps.HealthHandler.StorageHealth)
	r.GET("/health/ml", deps.HealthHandler.MLHealth)

	r.POST("/auth/register", deps.AuthHandler.Register)
	r.POST("/auth/login", deps.AuthHandler.Login)
	r.POST("/auth/refresh", deps.AuthHandler.Refresh)

	auth := r.Group("/auth")
	auth.Use(deps.AuthMiddleware)

	auth.POST("/logout", deps.AuthHandler.Logout)

	user := r.Group("/users")
	user.Use(deps.AuthMiddleware)
	{
		user.GET("/me", deps.UserHandler.GetMe)
		user.PATCH("/me", deps.UserHandler.UpdateMe)
		user.PATCH("/me/password", deps.UserHandler.UpdateMyPassword)
		user.DELETE("/me", deps.UserHandler.DeleteMe)
	}

	analysis := r.Group("/analyses")
	analysis.Use(deps.AuthMiddleware)
	{
		analysis.POST("", deps.AnalysisHandler.CreateAnalysis)
		analysis.GET("", deps.AnalysisHandler.ListAnalyses)
		analysis.GET("/:id", deps.AnalysisHandler.GetAnalysisByID)
		analysis.DELETE("/:id", deps.AnalysisHandler.DeleteAnalysis)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
