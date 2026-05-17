package app

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/pp-sem6-team/backend/internal/config"
	"github.com/pp-sem6-team/backend/internal/db"
	"github.com/pp-sem6-team/backend/internal/integration/ml"
	"github.com/pp-sem6-team/backend/internal/integration/storage/minio"
)

func loadConfig() *config.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	return config.Load()
}

func initInfrastructure(cfg *config.Config) (*gorm.DB, *minio.Client, ml.Client, error) {
	database, err := db.NewPostgres(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	storage, err := minio.New(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	mlClient := ml.NewHTTPClient(
		cfg.ML.BaseURL,
		cfg.ML.APIKey,
		cfg.ML.Timeout,
	)

	return database, storage, mlClient, nil
}
