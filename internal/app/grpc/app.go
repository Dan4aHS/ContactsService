package grpcapp

import (
	"ContactsService/internal/repository"
	grpchadlers "ContactsService/internal/transport/grpc"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	grpcServer *grpc.Server
	port       int
}

func NewApp(port int, repo repository.IContactRepository) *App {
	grpcServer := grpc.NewServer()

	grpchadlers.RegisterGRPCServer(grpcServer, repo)

	return &App{
		grpcServer: grpcServer,
		port:       port,
	}
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	a.grpcServer.GracefulStop()
}
