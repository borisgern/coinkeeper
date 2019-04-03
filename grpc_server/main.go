package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/coinkeeper/proto"
	"test/coinkeeper/services"
)

func main() {
	config, err := services.NewConfig()
	if err != nil {
		log.Fatalf("unable to create config: %v", err)
	}

	logger := services.NewLogger(config.Logger.LoggerLevel)
	em := ExpensesManager{
		Logger:logger.NewPrefix("grpc_server"),
	}

	lis, err := net.Listen("tcp", config.GRPCServer.ServerHost + ":" + config.GRPCServer.ServerPort)
	if err != nil {
		em.Logger.Fatalf("failed to create grpc listener: %v", err)
	}

	server := grpc.NewServer()
	expensespb.RegisterExpensesServiceServer(server, &em)

	em.Logger.Printf("grpc server started on port :%v", config.GRPCServer.ServerPort)
	err = server.Serve(lis)
	if err != nil {
		em.Logger.Fatalf("failed to serve: %v", err)
	}
}
