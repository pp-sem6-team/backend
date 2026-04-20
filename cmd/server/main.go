package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pp-sem6-team/backend/internal/config"
	"github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/handlers"

	_ "github.com/pp-sem6-team/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	cfg := config.Load()

	database, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	healthHandler := handlers.NewHealthHandler(database)

	r.GET("/health", healthHandler.Health)
	r.GET("/health/db", healthHandler.DBHealth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
