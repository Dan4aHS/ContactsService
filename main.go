package main

import (
	"ContactsService/internal/app"
	"ContactsService/internal/config"
	"ContactsService/internal/repository"
	in_memory "ContactsService/internal/repository/in-memory"
	"ContactsService/internal/repository/postgres"
	"ContactsService/internal/service"
	"log"
	"os"
)

func main() {
	cfg := config.GetConfig()
	var repo repository.IContactRepository
	if cfg.App.Repo == "postgres" {
		dbx, err := postgres.ConnectPostgres(cfg)
		if err != nil {
			log.Fatal(err)
		}
		repo = postgres.NewContactRepository(dbx)
	} else {
		repo = in_memory.NewContactRepository()
	}
	cs := service.NewContactService(repo)
	application := app.New(cfg.App.GRPC.Port, cfg.App.Rest.Port, cs)
	application.Run()
	app.Lock(make(chan os.Signal, 1))
}
