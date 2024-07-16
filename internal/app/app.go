package app

import (
	"fmt"
	"net/http"

	router "github.com/velvetriddles/bank-app/internal/delivery/http"
	"github.com/velvetriddles/bank-app/internal/di"
)

type App struct {
	container  *di.Container
	httpServer *http.Server
}

func NewApp(container *di.Container) *App {
	return &App{
		container: container,
	}
}

func (a *App) Run() error {
	r := router.SetupRouter(a.container.AccountHandler, a.container.Logger)

	a.httpServer = &http.Server{
		Addr:    a.container.Config.Server.Port,
		Handler: r,
	}

	a.container.Logger.Info(fmt.Sprintf("Starting server on port %s", a.container.Config.Server.Port))
	return a.httpServer.ListenAndServe()
}
