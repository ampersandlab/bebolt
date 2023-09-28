package examples

import "github.com/ampersandlab/bebolt/internal"

type ExampleAction struct {
	app *internal.App
}

func NewExampleAction(app *internal.App) *ExampleAction {
	return &ExampleAction{
		app: app,
	}
}
