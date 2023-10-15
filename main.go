package main

import (
	"net/http"

	cmd "github.com/ampersandlab/bebolt/cmd/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(cmd.CreateServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
