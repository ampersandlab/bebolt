package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ampersandlab/bebolt/internal"
	"github.com/ampersandlab/bebolt/internal/routes"
	"github.com/klauspost/compress/gzhttp"
	"github.com/rs/cors"
	"go.uber.org/fx"
)

func CreateServer(lc fx.Lifecycle) *http.Server {
	app := internal.NewApp()
	router := routes.NewRouter(app)
	router.SetWebRoutes()
	router.SetAPIRoutes()

	handler := http.Handler(router.Bunrouter)
	handler = cors.Default().Handler(handler)
	handler = gzhttp.GzipHandler(handler)

	srvr := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	addr := fmt.Sprintf("%s:%s", app.Env.Host, app.Env.Port)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			app.Logger.Infof("listening on %s\n", addr)
			go srvr.Serve(ln)
			app.Logger.Infoln("Press CTRL+C to exit...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.Logger.Warnln("Shutting down server...")
			if err := srvr.Shutdown(ctx); err != nil {
				log.Println(err)
			}
			app.Logger.Warnln("Server is DOWN")
			return nil
		},
	})

	return &srvr
}
