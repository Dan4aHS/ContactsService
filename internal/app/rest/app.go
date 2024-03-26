package restapp

import (
	"ContactsService/internal/service"
	http_v1 "ContactsService/internal/transport/http-v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router *gin.Engine
	Port   string
}

func NewApp(port string, cs service.IContactService) *App {
	c := http_v1.NewController(cs)
	router := gin.Default()
	router.POST("/contacts/new", c.CreateContactHandler)
	router.GET("/contacts", c.ListContactsHandler)
	router.POST("/contacts/:id", c.UpdateContactHandler)
	router.DELETE("/contacts/:id", c.DeleteContactHandler)
	router.GET("/contacts/:id", c.GetContactByIDHandler)
	return &App{
		Router: router,
		Port:   port,
	}
}

func (a *App) Run() {
	connStr := fmt.Sprintf("localhost:%s", a.Port)
	err := a.Router.Run(connStr)
	if err != nil {
		log.Fatal(err)
	}
}
