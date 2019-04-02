package main

import (
	"context"
	"github.com/lygo/runner"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"test/coinkeeper/grpc_server"
	"test/coinkeeper/proto"
	"test/coinkeeper/services"
	"time"
)

type expensesApp struct{
	Runner        *runner.App
	Config *services.Config
	Logger *services.CtxLogger
	HTTPServer *http.Server
	GRPCServer *grpc.Server
}

func newExpensesApp (config *services.Config, logger *services.LogService) (*expensesApp, error) {
	app := &expensesApp{}
	app.Logger = logger.NewPrefix("app")
	app.Config = config

	app.HTTPServer = &http.Server{
		Addr: ":" + config.HTTPServer.ServerPort,
		Handler: app.newRouter(config.HTTPServer, app.Logger),
	}

	app.GRPCServer = grpc.NewServer()
	expensespb.RegisterExpensesServiceServer(app.GRPCServer, &grpc_server.ExpensesManager{})

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

	lis, err := net.Listen("tcp", app.Config.GRPCServer.ServerHost + ":" + app.Config.GRPCServer.ServerPort)
	if err != nil {
		app.Logger.Fatalf("failed to create grpc listener: %v", err)
	}

	app.Runner.Runners = append(app.Runner.Runners, func() error {
		app.Logger.Printf("grpc server started on port :%v", app.Config.GRPCServer.ServerPort)
		err := app.GRPCServer.Serve(lis)
		if err == grpc.ErrServerStopped {
			return nil
		}
		return err
	})

	app.Runner.Slams = append(app.Runner.Slams, func() error {
		app.GRPCServer.GracefulStop()
		return nil
	})

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