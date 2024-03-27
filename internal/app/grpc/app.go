package grpcapp

import (
	"ContactsService/internal/service"
	grpchadlers "ContactsService/internal/transport/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewApp(port string, cs service.IContactService) *App {
	grpcServer := grpc.NewServer()
	grpchadlers.RegisterGRPCServer(grpcServer, cs)
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}
	return &App{
		gRPCServer: grpcServer,
		port:       p,
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
