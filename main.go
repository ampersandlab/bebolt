package main

import (
	"net/http"

	"github.com/ampersandlab/bebolt/cmd"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(cmd.CreateServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
