package app

import (
	"github.com/pp-sem6-team/backend/internal/logger"
	"go.uber.org/zap"
)

func Run() error {
	cfg := loadConfig()

	if err := logger.Init(cfg); err != nil {
		return err
	}

	defer logger.Log.Sync()

	database, storage, mlClient, err := initInfrastructure(cfg)
	if err != nil {
		logger.Log.Error("failed to init infrastructure", zap.Error(err))
		return err
	}

	deps := initDependencies(cfg, database, storage, mlClient)

	router := setupRouter(deps)

	return router.Run(":8080")
}
