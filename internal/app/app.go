package app

import (
	grpcapp "ContactsService/internal/app/grpc"
	restapp "ContactsService/internal/app/rest"
	"ContactsService/internal/service"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	GRPCApp *grpcapp.App
	RESTApp *restapp.App
}

func New(grpcPort string, restPort string, cs service.IContactService) *App {
	grpcApp := grpcapp.NewApp(grpcPort, cs)
	restApp := restapp.NewApp(restPort, cs)
	return &App{
		GRPCApp: grpcApp,
		RESTApp: restApp,
	}
}

func (a *App) Run() {
	go a.GRPCApp.Run()
	go a.RESTApp.Run()
}

func Lock(ch chan os.Signal) {
	defer func() {
		ch <- os.Interrupt
	}()
	if ch == nil {
		ch = make(chan os.Signal, 1)
	}
	// The correctness of the application is closed by a signal
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-ch
}
