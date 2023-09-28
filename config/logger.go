package config

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	suger := logger.Sugar()
	suger.Infoln("Setting up logger.")
	defer logger.Sync()
	return suger
}
