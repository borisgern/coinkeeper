package main

import (
	"context"
	"github.com/lygo/runner"
	"net/http"
	"test/coinkeeper/services"
	"time"
)

type expensesApp struct{
	Runner        *runner.App
	Config *services.Config
	Logger *services.CtxLogger
	HTTPServer *http.Server
}

func newExpensesApp (config *services.Config, logger *services.LogService) (*expensesApp, error) {
	app := &expensesApp{}
	app.Logger = logger.NewPrefix("app")
	app.Config = config

	app.HTTPServer = &http.Server{
		Addr: ":" + config.HTTPServer.ServerPort,
		Handler: app.newRouter(config.HTTPServer, app.Logger),
	}

	app.registerRunners()

	return app, nil
}

func (app *expensesApp) registerRunners() {
	var err error
	app.Runner = runner.New()
	defer func() {
		if err != nil {
			app.Runner.Shutdown()
		}
	}()

	app.Runner.Runners = append(app.Runner.Runners, func() error {
		app.Logger.Printf("http server started on port %v", app.HTTPServer.Addr)
		err := app.HTTPServer.ListenAndServe()
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	})

	app.Runner.Slams = append(app.Runner.Slams, func() error {
		ctxDown, cancelGraceful := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelGraceful()
		return app.HTTPServer.Shutdown(ctxDown)
	})
}