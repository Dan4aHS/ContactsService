package http_v1

import (
	"ContactsService/internal/models/mapper"
	"ContactsService/internal/models/rest"
	"ContactsService/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Controller struct {
	cs service.IContactService
}

func NewController(cs service.IContactService) *Controller {
	return &Controller{cs: cs}
}

func (c *Controller) CreateContactHandler(ctx *gin.Context) {
	var contact rest.CreateContactRequest
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, err := c.cs.CreateContact(ctx, mapper.ContactCreateRestToDB(contact))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (c *Controller) UpdateContactHandler(ctx *gin.Context) {
	var contact rest.UpdateContactRequest
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := c.cs.UpdateContact(ctx, mapper.ContactUpdateRestToDB(contact))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": contact.ID})
}

func (c *Controller) DeleteContactHandler(ctx *gin.Context) {

	id, err := uuid.Parse(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = c.cs.DeleteContact(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (c *Controller) GetContactByIDHandler(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	contact, err := c.cs.GetContactByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"contact": contact})
}

func (c *Controller) ListContactsHandler(ctx *gin.Context) {
	f := getFilters(ctx)
	contacts, err := c.cs.ListContacts(ctx, f)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"contacts": contacts})
}
