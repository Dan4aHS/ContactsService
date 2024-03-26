package main

import (
	"ContactsService/internal/app"
	"ContactsService/internal/config"
	"ContactsService/internal/repository"
	in_memory "ContactsService/internal/repository/in-memory"
	"ContactsService/internal/repository/postgres"
	"ContactsService/internal/service"
	http_v1 "ContactsService/internal/transport/http-v1"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//router := gin.Default()
	cfg := config.GetConfig()
	dbx, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewContactRepository(dbx)
	application := app.New(44044, repo)
	application.GRPCsrv.Run()
	//start(router, cfg)
}

func start(router *gin.Engine, cfg *config.Config) {
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
	c := http_v1.NewController(cs)
	router.POST("/contacts/new", c.CreateContactHandler)
	router.GET("/contacts", c.ListContactsHandler)
	router.POST("/contacts/:id", c.UpdateContactHandler)
	router.DELETE("/contacts/:id", c.DeleteContactHandler)
	router.GET("/contacts/:id", c.GetContactByIDHandler)
	router.Run("localhost:8080")
}
