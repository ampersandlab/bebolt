package http

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func (ea *Action) WelcomeHandler(w http.ResponseWriter, req bunrouter.Request) error {
	ea.app.Logger.Infoln("Logging from welcome handler")
	return bunrouter.JSON(w, bunrouter.H{
		"message": "Hello from Go API!",
	})
}
