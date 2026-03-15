package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pp-sem6-team/backend/internal/handlers"

	_ "github.com/pp-sem6-team/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
