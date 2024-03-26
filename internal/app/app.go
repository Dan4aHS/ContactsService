package app

import (
	grpcapp "ContactsService/internal/app/grpc"
	"ContactsService/internal/repository"
)

type App struct {
	GRPCsrv *grpcapp.App
}

func New(grpcPort int, repo repository.IContactRepository) *App {
	grpcApp := grpcapp.NewApp(grpcPort, repo)
	return &App{
		GRPCsrv: grpcApp,
	}
}
