package routes

import (
	"github.com/ampersandlab/bebolt/internal/actions/examples"
)

func (r *Router) SetWebRoutes() {
	//web routes go here
	ea := examples.NewExampleAction(r.App)

	r.Bunrouter.GET("/", ea.WelcomeHandler)
}
