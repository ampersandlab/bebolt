package routes

import (
	"github.com/ampersandlab/bebolt/internal"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type Router struct {
	Bunrouter *bunrouter.Router
	App       *internal.App
}

func NewRouter(app *internal.App) *Router {
	return &Router{
		App: app,
		Bunrouter: bunrouter.New(
			bunrouter.Use(reqlog.NewMiddleware(
				reqlog.WithVerbose(false),
				reqlog.FromEnv("BUNDEBUG"),
			)),
		),
	}
}
