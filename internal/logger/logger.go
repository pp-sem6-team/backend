package logger

import (
	"github.com/pp-sem6-team/backend/internal/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(cfg *config.Config) error {
	var (
		err error
	)

	if cfg.App.Env == "dev" {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}

	return nil
}
