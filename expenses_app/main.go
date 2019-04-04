package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/coinkeeper/services"
)

func main() {
	config, err := services.NewConfig()
	if err != nil {
		log.Fatalf("unable to create config: %v", err)
	}

	logger := services.NewLogger(config.Logger.LoggerLevel)

	app, err := newExpensesApp(config, logger)
	if err != nil {
		logger.Fatalf("unable to start app")
	}

	app.Runner.Run()

	// buffer should be more than one!
	ch := make(chan os.Signal, 10)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	defer close(ch)

	<-app.Runner.Started

	go func() {
		sig, ok := <-ch
		if ok {
			log.Printf("shutdown process on %s system signal\n", sig)
		}
		app.Runner.Shutdown()
	}()

	os.Exit(<-app.Runner.Done)
}
