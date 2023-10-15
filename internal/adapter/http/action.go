package http

import "github.com/ampersandlab/bebolt/internal"

type Action struct {
	app *internal.App
}

func NewAction(app *internal.App) *Action {
	return &Action{
		app: app,
	}
}
