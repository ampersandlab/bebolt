package internal

import (
	"sync"

	"github.com/ampersandlab/bebolt/config"
	"go.uber.org/zap"
)

var (
	app     *App
	appOnce sync.Once
)

type App struct {
	Orm    *Database
	Logger *zap.SugaredLogger
	Env    *config.Env
}

func NewApp() *App {
	appOnce.Do(func() {
		env := config.NewEnv()
		logger := config.NewLogger()
		connect := NewDatabaseConnection(env)
		app = &App{
			Orm:    connect,
			Logger: logger,
			Env:    env,
		}
	})
	return app
}
