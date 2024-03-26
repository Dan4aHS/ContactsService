package service

import (
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"context"
	"github.com/google/uuid"
)

type IContactService interface {
	CreateContact(ctx context.Context, contact dbmodels.Contact) (uuid.UUID, error)
	UpdateContact(ctx context.Context, contact dbmodels.Contact) error
	DeleteContact(ctx context.Context, id uuid.UUID) error
	GetContactByID(ctx context.Context, id uuid.UUID) (entity.Contact, error)
	ListContacts(ctx context.Context, f map[string]any) ([]entity.Contact, error)
}
