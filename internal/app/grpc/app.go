package grpcapp

import (
	"ContactsService/internal/service"
	grpchadlers "ContactsService/internal/transport/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewApp(port int, cs service.IContactService) *App {
	grpcServer := grpc.NewServer()
	grpchadlers.RegisterGRPCServer(grpcServer, cs)

	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err = a.gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
