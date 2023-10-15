package routes

import (
	http "github.com/ampersandlab/bebolt/internal/adapter/http"
)

func (r *Router) SetWebRoutes() {
	//web routes go here
	action := http.NewAction(r.App)

	r.Bunrouter.GET("/", action.WelcomeHandler)
}
